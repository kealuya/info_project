package handler

import (
	"fmt"
	"github.com/jinzhu/copier"
	"my_pilot/common"
	"my_pilot/pkg/api_szjl"
)

// QueryRatePlanRequest  报价查询请求参数
type QueryRatePlanRequest struct {
	HotelId      int    `json:"hotelId"`      // 酒店编号
	CheckInDate  string `json:"checkInDate"`  // 入住日期
	CheckOutDate string `json:"checkOutDate"` // 离店日期
	RoomGroups   []struct {
		Adults    int    `json:"adults,omitempty"`    // 成人数
		Children  int    `json:"children,omitempty"`  // 儿童数
		ChildAges string `json:"childAges,omitempty"` // 儿童年龄
	} // 房间信息
	IsSkipCheckCondition bool `json:"isSkipCheckCondition"` //是否跳过预订规则校验
}

func QueryHotelRatePlan(qr QueryRatePlanRequest) (queryRatePlanResult *api_szjl.QueryRatePlanResult, bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})

	qrRequest := api_szjl.QueryRatePlanRequestData{}
	err := copier.Copy(&qrRequest, &qr)
	common.ErrorHandler(err)

	qrRequest.QueryType = 0
	hotelRatePlan, err := api_szjl.QueryRatePlan(qrRequest)
	common.ErrorHandler(err, fmt.Sprintf("%v", qr))

	return hotelRatePlan, nil
}
