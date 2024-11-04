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
	workerCount = 20
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
	locationInfo := repository.GetDistinctStatesByCountry(70007)
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

func SaveHotelDetailInfo() (bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})

	const (
		saveHotelDetailInfoPageSize    = 1000 // 每次从数据库获取的酒店数量
		saveHotelDetailInfoWorkerCount = 20   // 并发工作的协程数量
	)

	// 创建任务通道和错误通道
	taskChan := make(chan []repository.HotelInfo, saveHotelDetailInfoWorkerCount)
	errChan := make(chan error, saveHotelDetailInfoWorkerCount)

	// 创建等待组
	var wg sync.WaitGroup

	// 启动工作协程池
	for i := 0; i < saveHotelDetailInfoWorkerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processHotelBatch(taskChan, errChan)
		}()
	}

	// 错误处理协程
	go func() {
		for err := range errChan {
			if err != nil {
				logs.Error(fmt.Printf("Error: %v\n", err))
			}
		}
	}()

	start := 0
	totalProcessed := 0

	for {
		// 分页获取酒店信息
		hotels := repository.GetHotelsAll(start, saveHotelDetailInfoPageSize)

		// 如果没有更多数据，退出循环
		if len(hotels) == 0 {
			break
		}

		// 将酒店数据发送到任务通道
		taskChan <- hotels

		totalProcessed += len(hotels)
		logs.Info(fmt.Sprintf("已发送 %d 条酒店记录到处理队列\n", totalProcessed))

		// 更新起始位置
		start += len(hotels)
	}

	// 关闭任务通道
	close(taskChan)

	// 等待所有工作协程完成
	wg.Wait()

	// 关闭错误通道
	close(errChan)

	logs.Info(fmt.Sprintf("酒店详情数据获取完成，共处理 %d 条记录\n", totalProcessed))
	return nil
}

// 处理一批酒店数据的工作函数
func processHotelBatch(taskChan <-chan []repository.HotelInfo, errChan chan<- error) {

	defer common.RecoverHandler(func(err error) {
		fmt.Println(err)
		errChan <- err
	})

	now := time.Now()
	for hotels := range taskChan {
		for _, hotel := range hotels {
			// 查询酒店详情
			requestData := api_szjl.QueryHotelDetailRequestData{
				HotelId: hotel.HotelId,
				Params:  "1,2",
			}

			hotelDetail, err := api_szjl.QueryHotelDetail(requestData)
			common.ErrorHandler(err, hotel.HotelId)
			for _, item := range hotelDetail.HotelDetailList {

				// 转换并保存酒店静态信息
				hotelStaticInfo := convertToHotelStaticInfo(item.HotelInfo, now)

				repository.InsertHotelStaticInfo(hotelStaticInfo)

				// 转换并保存房型信息
				var roomTypes []repository.RoomType
				for _, roomType := range item.RoomTypeList {
					roomTypeInfo := convertToRoomType(roomType, hotel.HotelId, now)
					roomTypes = append(roomTypes, roomTypeInfo)
				}

				if len(roomTypes) > 0 {
					repository.InsertRoomTypes(roomTypes)
				}
			}
			logs.Info(fmt.Sprintf("酒店 %d 详情处理完成\n", hotel.HotelId))
		}
	}
}

// convertToHotelStaticInfo 将API返回的HotelInfo转换为数据库的HotelStaticInfo
func convertToHotelStaticInfo(hotel api_szjl.HotelStaticInfo, now time.Time) repository.HotelStaticInfo {
	return repository.HotelStaticInfo{
		HotelId:           hotel.HotelId,
		ThemeType:         hotel.ThemeType,
		HotelNameCn:       hotel.HotelNameCn,
		HotelNameEn:       hotel.HotelNameEn,
		Star:              hotel.Star,
		AddressCn:         hotel.AddressCn,
		AddressEn:         hotel.AddressEn,
		CountryId:         hotel.CountryId,
		StateId:           hotel.StateId,
		CityId:            hotel.CityId,
		Area:              hotel.Area,
		BusinessCircle:    hotel.BusinessCircle,
		Phone:             hotel.Phone,
		Longitude:         hotel.Longitude,
		Latitude:          hotel.Latitude,
		Fax:               hotel.Fax,
		Email:             hotel.Email,
		PostCode:          hotel.PostCode,
		CheckPolicy:       hotel.CheckPolicy,
		ChildrenPolicy:    hotel.ChildrenPolicy,
		PetPolicy:         hotel.PetPolicy,
		EstablishmentDate: hotel.EstablishmentDate,
		RenovationDate:    hotel.RenovationDate,
		HotelGroup:        hotel.HotelGroup,
		HotelBrand:        hotel.HotelBrand,
		Facilities:        hotel.Facilities,
		CardType:          hotel.CardType,
		IntroduceCn:       hotel.IntroduceCn,
		IntroduceEn:       hotel.IntroduceEn,
		SellType:          hotel.SellType,
		Created:           now,
		Updated:           now,
	}
}

// convertToRoomType 将API返回的RoomType转换为数据库的RoomType
func convertToRoomType(room api_szjl.RoomType, hotelId int, now time.Time) repository.RoomType {
	return repository.RoomType{
		HotelId:         hotelId,
		RoomTypeId:      room.RoomTypeId,
		RoomTypeCn:      room.RoomTypeCn,
		RoomTypeEn:      room.RoomTypeEn,
		BasisRoomId:     room.BasisRoomId,
		BasisRoomCn:     room.BasisRoomCn,
		BedType:         room.BedType,
		Maximize:        room.Maximize,
		Acreage:         room.Acreage,
		BedWidth:        room.BedWidth,
		FloorDistribute: room.FloorDistribute,
		Facilities:      room.Facilities,
		ExtraBedState:   room.ExtraBedState,
		BedCount:        room.BedCount,
		Created:         now,
		Updated:         now,
	}
}
