package api_szjl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

type QueryCityRequestData struct {
	CountryId int `json:"countryId,omitempty"`
	StateId   int `json:"stateId,omitempty"`
	CityId    int `json:"cityId,omitempty"`
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

// HotelGeo Response structures
type HotelGeo struct {
	CountryId int    `json:"countryId"`
	CountryCn string `json:"countryCn"`
	CountryEn string `json:"countryEn"`
	StateId   int    `json:"stateId"`
	StateCn   string `json:"stateCn"`
	StateEn   string `json:"stateEn"`
	CityId    int    `json:"cityId"`
	CityCn    string `json:"cityCn"`
	CityEn    string `json:"cityEn"`
}

type QueryCityResult struct {
	Count        int        `json:"count"`
	PageIndex    int        `json:"pageIndex"`
	PageSize     int        `json:"pageSize"`
	HotelGeoList []HotelGeo `json:"hotelGeoList"`
}

var queryCityUrl = baseURL + "/city/queryCity.json"

func QueryCity(requestData QueryCityRequestData) (*QueryCityResult, error) {
	// Create resty client
	client := resty.New()

	// Generate timestamp
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// Generate sign
	sign := generateSign(secretKey, appKey, timestamp)

	// Create request body
	req := Request[QueryCityRequestData]{
		Head: RequestHead{
			AppKey:    appKey,
			Timestamp: timestamp,
			Sign:      sign,
			Version:   "3.0.1",
		},
		Data: requestData,
		//	RequestData{
		//	PageIndex: 1,
		//	PageSize:  1000,
		//	// 可选参数
		//	CountryId: 70007, // 中国  或者不写，全球
		//	//StateId:   70020,
		//	// CityId: 70002,
		//},
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %v\n", err)
	}

	// Send request using resty
	var response Response[QueryCityResult]
	resp, err := client.R().
		SetQueryParam("reqData", string(jsonData)).
		SetResult(&response).
		Get(queryCityUrl)

	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v\n", err)

	}

	// Check response status
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("Request failed with status code: %d\n", resp.StatusCode())
	}

	// Print results
	if response.Code == 0 {
		//fmt.Println("Request successful!")
		//fmt.Printf("Response ID: %s\n", response.RespId)
		//fmt.Printf("Total cities: %d\n", response.Result.Count)
		//fmt.Printf("Page: %d/%d\n", response.Result.PageIndex, response.Result.PageSize)
		//
		//// Print city details
		//for i, geo := range response.Result.HotelGeoList {
		//	fmt.Printf("\nLocation %d:\n", i+1)
		//	fmt.Printf("Country: %s (%s) [ID: %d]\n",
		//		geo.CountryCn, geo.CountryEn, geo.CountryId)
		//	fmt.Printf("State: %s (%s) [ID: %d]\n",
		//		geo.StateCn, geo.StateEn, geo.StateId)
		//	fmt.Printf("City: %s (%s) [ID: %d]\n",
		//		geo.CityCn, geo.CityEn, geo.CityId)
		//}

		return &response.Result, nil
	} else {
		return nil, fmt.Errorf("Request failed with code %d: %s\n", response.Code, response.ErrorMsg)
	}
}
