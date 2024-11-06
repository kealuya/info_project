package api_szjl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

// QueryOrderPriceRequestData 订单报价请求参数
type QueryOrderPriceRequestData struct {
	HotelId       int         `json:"hotelId"`       // 酒店编号
	KeyId         string      `json:"keyId"`         // 产品编号
	CheckInDate   string      `json:"checkInDate"`   // 入住日期
	CheckOutDate  string      `json:"checkOutDate"`  // 离店日期
	NightlyPrices string      `json:"nightlyPrices"` // 每日价格
	RoomGroups    []RoomGroup `json:"roomGroups"`    // 房间信息
	QueryType     int         `json:"queryType"`     // 数据类型
}

// BookingMessage 预订信息
type BookingMessage struct {
	Code    int    `json:"code"`    // 可订代码
	Message string `json:"message"` // 预订错误信息
}

// OrderPrice 订单价格信息
type OrderPrice struct {
	HotelRatePlans []HotelRatePlan `json:"hotelRatePlans"` // 酒店产品数组
	BookingMessage BookingMessage  `json:"bookingMessage"` // 是否可预订信息
}

// QueryOrderPriceResult 订单报价响应结果
type QueryOrderPriceResult struct {
	OrderPrice OrderPrice `json:"orderPrice"` // 订单价格信息
}

// QueryOrderPrice 查询订单报价
func QueryOrderPrice(requestData QueryOrderPriceRequestData) (*QueryOrderPriceResult, error) {
	var queryOrderPriceUrl = baseURL + "/hotel/queryOrderPrice.json"

	// 创建 resty 客户端
	client := resty.New()

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	sign := generateSign(secretKey, appKey, timestamp)

	// 创建请求体
	req := Request[QueryOrderPriceRequestData]{
		Head: RequestHead{
			AppKey:    appKey,
			Timestamp: timestamp,
			Sign:      sign,
			Version:   "3.0.1",
		},
		Data: requestData,
	}

	// 转换请求为JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %v", err)
	}

	// 发送请求
	var response Response[QueryOrderPriceResult]
	resp, err := client.R().
		SetQueryParam("reqData", string(jsonData)).
		SetResult(&response).
		Post(queryOrderPriceUrl)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	// 检查响应状态
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	// 检查业务响应码
	if response.Code == 0 {
		return &response.Result, nil
	} else {
		return nil, fmt.Errorf("request failed with code %d: %s", response.Code, response.ErrorMsg)
	}

}
