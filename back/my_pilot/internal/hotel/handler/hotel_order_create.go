package handler

import (
	"fmt"
	"github.com/jinzhu/copier"
	"my_pilot/common"
	"my_pilot/pkg/api_szjl"
)

// CreateOrderRequest  报价查询请求参数
type CreateOrderRequest struct {
	CustomerOrderCode string  `json:"customerOrderCode"` // 客户订单号
	HotelId           int     `json:"hotelId"`           // 酒店编号
	KeyId             string  `json:"keyId"`             // 产品编号
	CheckInDate       string  `json:"checkInDate"`       // 入住日期
	CheckOutDate      string  `json:"checkOutDate"`      // 离店日期
	NightlyPrices     string  `json:"nightlyPrices"`     // 每日价格
	TotalPrice        float64 `json:"totalPrice"`        // 总金额
	RoomGroups        []struct {
		Adults          int    `json:"adults,omitempty"`    // 成人数
		Children        int    `json:"children,omitempty"`  // 儿童数
		ChildAges       string `json:"childAges,omitempty"` // 儿童年龄
		CheckInPersions []struct {
			LastName    string `json:"lastName"`              // 入住人的姓
			FirstName   string `json:"firstName"`             // 入住人的名
			Nationality string `json:"nationality,omitempty"` // 国籍
		} `json:"checkInPersions"` // 入住人信息
	} `json:"roomGroups"` // 房间信息
	HotelRemark string `json:"hotelRemark"` // 给酒店备注
}

func CreateOrder(qr CreateOrderRequest) (queryOrderPriceResult *api_szjl.CreateOrderResult, bizError error) {
	defer common.RecoverHandler(func(err error) {
		bizError = err
	})

	qrRequest := api_szjl.CreateOrderRequestData{}
	err := copier.Copy(&qrRequest, &qr)
	common.ErrorHandler(err)

	createOrder, err := api_szjl.CreateOrder(qrRequest)
	common.ErrorHandler(err, fmt.Sprintf("%v", qr))

	return createOrder, nil
}
