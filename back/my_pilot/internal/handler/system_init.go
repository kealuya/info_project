package handler

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"my_pilot/common"
	"my_pilot/internal/repository"
	"my_pilot/pkg/api_szjl"
	"sync"
	"time"
)

// 定义工作池大小和批处理大小
const (
	workerCount = 100
	batchSize   = 100
)

// SaveLocationInfo 批量城市列表落地
func SaveLocationInfo() (bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})

	// 创建任务通道
	locationChan := make(chan []repository.LocationInfo, workerCount)
	// 创建错误通道
	errChan := make(chan error, workerCount)
	// 创建等待组
	var wg sync.WaitGroup

	// 启动工作池
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go insertLocationsWorker(locationChan, errChan, &wg)
	}

	// 错误处理goroutine
	go func() {
		for err := range errChan {
			if err != nil {
				logs.Error(fmt.Printf("Error: %v\n", err))
			}
		}
	}()

	totalProcessed := 0
	pageIndex := 1

	var now = time.Now()
	for {
		// 1. 查询数据
		requestData := api_szjl.QueryCityRequestData{
			PageIndex: pageIndex,
			PageSize:  1000,
		}
		cityResult, err := api_szjl.QueryCity(requestData)
		common.ErrorHandler(err)

		if cityResult == nil || len(cityResult.HotelGeoList) == 0 {
			break
		}

		// 2. 转换数据
		locations := make([]repository.LocationInfo, len(cityResult.HotelGeoList))

		for i, hotel := range cityResult.HotelGeoList {
			locations[i] = convertToLocationInfo(hotel, now)
		}

		// 3. 分批发送到通道
		for i := 0; i < len(locations); i += batchSize {
			end := i + batchSize
			if end > len(locations) {
				end = len(locations)
			}
			batch := locations[i:end]
			locationChan <- batch
		}

		// 4. 更新进度
		totalProcessed += len(locations)
		logs.Info(fmt.Sprintf("已处理 %d/%d 条记录\n", totalProcessed, cityResult.Count))

		if totalProcessed >= cityResult.Count {
			break
		}

		pageIndex++
	}

	// 关闭通道
	close(locationChan)

	// 等待所有工作完成
	wg.Wait()
	close(errChan)

	logs.Info(fmt.Sprintf("数据获取完成，共处理 %d 条记录\n", totalProcessed))
	return nil
}

// insertLocationsWorker 函数处理数据插入
func insertLocationsWorker(locationChan chan []repository.LocationInfo, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	for locations := range locationChan {
		func() {
			defer common.RecoverHandler(func(err error) {
				errChan <- err
			})
			repository.InsertLocations(locations)
		}()
	}
}

// convertToLocationInfo 将HotelGeo转换为LocationInfo
func convertToLocationInfo(hotel api_szjl.HotelGeo, now time.Time) repository.LocationInfo {
	return repository.LocationInfo{
		CountryId: hotel.CountryId,
		CountryCn: hotel.CountryCn,
		CountryEn: hotel.CountryEn,
		StateId:   hotel.StateId,
		StateCn:   hotel.StateCn,
		StateEn:   hotel.StateEn,
		CityId:    hotel.CityId,
		CityCn:    hotel.CityCn,
		CityEn:    hotel.CityEn,
		Created:   now,
		Updated:   now,
	}
}

// SaveHotelInfo 批量保存酒店信息
func SaveHotelInfo() (bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})
	// 只获取中国本地酒店 70007
	locationInfo, err := repository.GetDistinctStatesByCountry(70007)
	common.ErrorHandler(err)
	for _, locationInfo := range locationInfo {
		err := saveHotelInfoWithState(locationInfo.StateId, locationInfo.StateCn)
		common.ErrorHandler(err)
	}
	return nil
}
func saveHotelInfoWithState(stateId int, stateCn string) (bizError error) {

	// 创建任务通道
	hotelChan := make(chan []repository.HotelInfo, workerCount)
	// 创建错误通道
	errChan := make(chan error, workerCount)
	// 创建等待组
	var wg sync.WaitGroup

	// 启动工作池
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go insertHotelsWorker(hotelChan, errChan, &wg)
	}

	// 错误处理goroutine
	go func() {
		for err := range errChan {
			if err != nil {
				logs.Error(fmt.Printf("Error: %v\n", err))
			}
		}
	}()

	totalProcessed := 0
	pageIndex := 1

	var now = time.Now()
	for {
		// 1. 查询数据
		requestData := api_szjl.QueryHotelRequestData{
			CountryId: 70007,
			StateId:   stateId,
			PageIndex: pageIndex,
			PageSize:  1000,
		}
		hotelResult, err := api_szjl.QueryHotel(requestData)
		common.ErrorHandler(err)

		if hotelResult == nil || len(hotelResult.Hotels) == 0 {
			break
		}

		// 2. 转换数据
		hotels := make([]repository.HotelInfo, len(hotelResult.Hotels))
		for i, hotel := range hotelResult.Hotels {
			hotels[i] = convertToHotelInfo(hotel, now)
		}

		// 3. 分批发送到通道
		for i := 0; i < len(hotels); i += batchSize {
			end := i + batchSize
			if end > len(hotels) {
				end = len(hotels)
			}
			batch := hotels[i:end]
			hotelChan <- batch
		}

		// 4. 更新进度
		totalProcessed += len(hotels)
		logs.Info(fmt.Sprintf("已处理[%s] %d/%d 条酒店记录\n", stateCn, totalProcessed, hotelResult.Count))

		if totalProcessed >= hotelResult.Count {
			break
		}

		pageIndex++
	}

	// 关闭通道
	close(hotelChan)

	// 等待所有工作完成
	wg.Wait()
	close(errChan)

	logs.Info(fmt.Sprintf("[%s]酒店数据获取完成，共处理 %d 条记录\n", stateCn, totalProcessed))
	return nil
}

// insertHotelsWorker 处理酒店数据插入
func insertHotelsWorker(hotelChan chan []repository.HotelInfo, errChan chan error, wg *sync.WaitGroup) {
	defer wg.Done()

	for hotels := range hotelChan {
		func() {
			defer common.RecoverHandler(func(err error) {
				errChan <- err
			})
			repository.InsertHotels(hotels)
		}()
	}
}

// convertToHotelInfo 将API返回的HotelInfo转换为数据库的HotelInfo
func convertToHotelInfo(hotel api_szjl.HotelInfo, now time.Time) repository.HotelInfo {
	updateTime, _ := time.Parse("2006-01-02 15:04:05", hotel.UpdateTime)

	return repository.HotelInfo{
		HotelId:             hotel.HotelId,
		CountryId:           hotel.CountryId,
		StateId:             hotel.StateId,
		CityId:              hotel.CityId,
		Star:                hotel.Star,
		HotelNameCn:         hotel.HotelNameCn,
		HotelNameEn:         hotel.HotelNameEn,
		AddressCn:           hotel.AddressCn,
		AddressEn:           hotel.AddressEn,
		Phone:               hotel.Phone,
		Longitude:           hotel.Longitude,
		Latitude:            hotel.Latitude,
		InstantConfirmation: hotel.InstantConfirmation,
		SellType:            hotel.SellType,
		UpdateTime:          updateTime,
		Created:             now,
		Updated:             now,
	}
}
