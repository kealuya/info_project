package handler

import (
	"fmt"
	"github.com/jinzhu/copier"
	"my_pilot/common"
	"my_pilot/pkg/api_szjl"
)

// QueryOrderPriceRequest  报价查询请求参数
type QueryOrderPriceRequest struct {
	HotelId       int    `json:"hotelId"`       // 酒店编号
	KeyId         string `json:"keyId"`         // 产品编号
	CheckInDate   string `json:"checkInDate"`   // 入住日期
	CheckOutDate  string `json:"checkOutDate"`  // 离店日期
	NightlyPrices string `json:"nightlyPrices"` // 每日价格
	RoomGroups    struct {
		Adults    int    `json:"adults,omitempty"`    // 成人数
		Children  int    `json:"children,omitempty"`  // 儿童数
		ChildAges string `json:"childAges,omitempty"` // 儿童年龄
	} // 房间信息
}

func QueryOrderPrice(qr QueryOrderPriceRequest) (queryOrderPriceResult *api_szjl.QueryOrderPriceResult, bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})

	qrRequest := api_szjl.QueryOrderPriceRequestData{}
	err := copier.Copy(&qrRequest, &qr)
	common.ErrorHandler(err)

	qrRequest.QueryType = 0
	hotelOrderPrice, err := api_szjl.QueryOrderPrice(qrRequest)
	common.ErrorHandler(err, fmt.Sprintf("%v", qr))

	return hotelOrderPrice, nil
}
