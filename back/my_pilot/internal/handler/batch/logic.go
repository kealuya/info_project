package batch

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/jinzhu/copier"
	"my_pilot/internal/repository"
	"my_pilot/pkg/api_szjl"
	"time"
)

func getHotelInfoFromDb() chan []repository.HotelInfo {
	// 一次读取1000条，发一次是1000条，chan容量2，所以一次存储2000条。
	taskChan := make(chan []repository.HotelInfo, 2)

	const saveHotelDetailInfoPageSize = 1000 // 每次从数据库获取的酒店数量
	start := 0
	totalProcessed := 0
	go func() {
		defer close(taskChan)
		for {
			// 分页获取酒店信息
			var hotels []repository.HotelInfo
			func() {
				defer func() {
					if err := recover(); err != nil {
						logs.Error(fmt.Sprintf("getHotelInfoFromDb error ::%v,start:%v,limit:%v", err, start, saveHotelDetailInfoPageSize))
					}
				}()
				hotels = repository.GetHotelsAll(start, saveHotelDetailInfoPageSize)
			}()

			// 如果没有更多数据，退出循环
			if len(hotels) == 0 {
				break
			}

			// 将酒店数据发送到任务通道
			taskChan <- hotels

			totalProcessed += len(hotels)
			logs.Info(fmt.Sprintf("已从db读取 %d 条酒店记录\n", totalProcessed))

			// 更新起始位置
			start += len(hotels)
		}
	}()

	return taskChan

}

func getHotelStaticInfoFromSZJLAndInsertDb(hotel repository.HotelInfo) (bizErr error) {

	defer func() {
		if err := recover(); err != nil {
			logs.Error(fmt.Sprintf("getHotelStaticInfoFromSZJLAndInsertDb error ::%v,hotel_id:%v", err, hotel.HotelId))
			bizErr = fmt.Errorf("getHotelStaticInfoFromSZJLAndInsertDb error ::%v,hotel_id:%v", err, hotel.HotelId)
			return
		}
	}()

	// 查询酒店详情
	requestData := api_szjl.QueryHotelDetailRequestData{
		HotelId: hotel.HotelId,
		Params:  "1,2",
	}

	hotelDetail, err := api_szjl.QueryHotelDetail(requestData)
	if err != nil {
		return fmt.Errorf("QueryHotelDetail error: %v , hotel_id: %v, id: %v", err, hotel.HotelId, hotel.Id)
	}
	now := time.Now()
	for _, item := range hotelDetail.HotelDetailList {

		// 转换并保存酒店静态信息
		hotelStaticInfo := convertToHotelStaticInfo(item.HotelInfo, now)

		err := repository.InsertHotelStaticInfo(hotelStaticInfo)
		if err != nil {
			return fmt.Errorf("insert hotel static info error: %v,%+v", err, hotelStaticInfo)
		}

		// 转换并保存房型信息
		var roomTypes []repository.RoomType
		for _, roomType := range item.RoomTypeList {

			resoStruct := repository.RoomType{}
			err := copier.Copy(&resoStruct, &roomType)
			if err != nil {
				return fmt.Errorf("copier error: %v,%+v", err, roomType)
			}
			resoStruct.HotelId = hotel.HotelId
			resoStruct.Created = now
			resoStruct.Updated = now
			roomTypes = append(roomTypes, resoStruct)
		}

		if len(roomTypes) > 0 {
			err := repository.InsertRoomTypes(roomTypes)
			if err != nil {
				return fmt.Errorf("repository.InsertRoomTypes(roomTypes) error: %v , hotel_id: %v", err, hotel.HotelId)
			}
		}

		// 转换并保存价格类型信息
		var rateTypes []repository.RateType
		for _, rateType := range item.RateTypeList {
			//roomTypeInfo := convertToRoomType(roomType, hotel.HotelId, now)

			resoStruct := repository.RateType{}
			err := copier.Copy(&resoStruct, &rateType)
			if err != nil {
				return fmt.Errorf("copier error: %v,%+v", err, rateType)
			}
			resoStruct.HotelId = hotel.HotelId
			resoStruct.Created = now
			resoStruct.Updated = now
			rateTypes = append(rateTypes, resoStruct)
		}

		if len(rateTypes) > 0 {
			err := repository.InsertRateTypes(rateTypes)
			if err != nil {
				return fmt.Errorf("repository.InsertRateTypes(rateTypes) error: %v , hotel_id: %v", err, hotel.HotelId)
			}
		}

		// 转换并保存图片类型信息
		var images []repository.Image
		for _, image := range item.ImageList {
			//roomTypeInfo := convertToRoomType(roomType, hotel.HotelId, now)

			resoStruct := repository.Image{}
			err := copier.Copy(&resoStruct, &image)
			if err != nil {
				return fmt.Errorf("copier error: %v,%+v", err, image)
			}
			resoStruct.HotelId = hotel.HotelId
			resoStruct.Created = now
			resoStruct.Updated = now
			images = append(images, resoStruct)
		}

		if len(images) > 0 {
			err := repository.InsertImages(images)
			if err != nil {
				return fmt.Errorf("repository.InsertImages(images) error: %v , hotel_id: %v", err, hotel.HotelId)
			}
		}

	}
	return nil
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
