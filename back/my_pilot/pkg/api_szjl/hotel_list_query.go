package api_szjl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

// QueryHotelRequestData 请求参数结构体
type QueryHotelRequestData struct {
	CountryId  int    `json:"countryId,omitempty"`
	StateId    int    `json:"stateId,omitempty"`
	CityId     int    `json:"cityId,omitempty"`
	PageIndex  int    `json:"pageIndex"`
	PageSize   int    `json:"pageSize"`
	UpdateTime string `json:"updateTime,omitempty"`
	IsConfirm  int    `json:"isConfirm,omitempty"`
}

// HotelInfo 酒店信息结构体
type HotelInfo struct {
	HotelId             int    `json:"hotelId"`
	CountryId           int    `json:"countryId"`
	StateId             int    `json:"stateId"`
	CityId              int    `json:"cityId"`
	Star                int    `json:"star"`
	HotelNameCn         string `json:"hotelNameCn"`
	HotelNameEn         string `json:"hotelNameEn"`
	AddressCn           string `json:"addressCn"`
	AddressEn           string `json:"addressEn"`
	Phone               string `json:"phone"`
	Longitude           string `json:"longitude"`
	Latitude            string `json:"latitude"`
	InstantConfirmation int    `json:"instantConfirmation"`
	SellType            int    `json:"sellType"`
	UpdateTime          string `json:"updateTime"`
}

// QueryHotelResult 响应结果结构体
type QueryHotelResult struct {
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
	Hotels    []HotelInfo `json:"hotels"`
}

var queryHotelUrl = baseURL + "/hotel/queryHotelList.json"

func QueryHotel(requestData QueryHotelRequestData) (*QueryHotelResult, error) {
	// 创建 resty 客户端
	client := resty.New()

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	sign := generateSign(secretKey, appKey, timestamp)

	// 创建请求体
	req := Request[QueryHotelRequestData]{
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
	var response Response[QueryHotelResult]
	resp, err := client.R().
		SetQueryParam("reqData", string(jsonData)).
		SetResult(&response).
		Get(queryHotelUrl)

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
