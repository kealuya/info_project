package api_szjl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

// CreateOrderRequestData 创建订单请求参数
type CreateOrderRequestData struct {
	CustomerOrderCode string      `json:"customerOrderCode"` // 客户订单号
	HotelId           int         `json:"hotelId"`           // 酒店编号
	KeyId             string      `json:"keyId"`             // 产品编号
	CheckInDate       string      `json:"checkInDate"`       // 入住日期
	CheckOutDate      string      `json:"checkOutDate"`      // 离店日期
	NightlyPrices     string      `json:"nightlyPrices"`     // 每日价格
	TotalPrice        float64     `json:"totalPrice"`        // 总金额
	RoomGroups        []RoomGroup `json:"roomGroups"`        // 房间信息
	HotelRemark       string      `json:"hotelRemark"`       // 给酒店备注
}

// CheckInPersion 创建订单请求参数
type CheckInPersion struct {
	LastName    string `json:"lastName"`              // 入住人的姓
	FirstName   string `json:"firstName"`             // 入住人的名
	Nationality string `json:"nationality,omitempty"` // 国籍
}

// CreateOrderResult 创建订单响应结果
type CreateOrderResult struct {
	CreateOrder struct {
		OrderCode   string `json:"orderCode"`   // 订单编号
		OrderStatus int    `json:"orderStatus"` // 订单状态
	} `json:"createOrder"`
}

// CreateOrder 创建订单
func CreateOrder(requestData CreateOrderRequestData) (*CreateOrderResult, error) {
	var createOrderUrl = baseURL + "/order/createOrder.json"

	// 创建 resty 客户端
	client := resty.New()
	client.SetTimeout(30 * time.Second) // 设置30秒超时

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	sign := generateSign(secretKey, appKey, timestamp)

	// 创建请求体
	req := Request[CreateOrderRequestData]{
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
	var response Response[CreateOrderResult]
	resp, err := client.R().
		SetQueryParam("reqData", string(jsonData)).
		SetResult(&response).
		Post(createOrderUrl)

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
