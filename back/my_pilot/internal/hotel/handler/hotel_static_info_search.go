package handler

import (
	"fmt"
	"my_pilot/common"
	"my_pilot/pkg/api_szjl"
)

func SearchHotelStaticInfo(hotelId int) (queryHotelDetailResponse *api_szjl.QueryHotelDetailResponse, bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})
	requestData := api_szjl.QueryHotelDetailRequestData{
		HotelId: hotelId,
		Params:  "1,2",
	}
	detail, err := api_szjl.QueryHotelDetail(requestData)
	common.ErrorHandler(err, fmt.Sprintf("hotelId:%d", hotelId))

	return detail, nil
}
