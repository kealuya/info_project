package api_szjl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

// 请求体结构
type QueryHotelDetailRequestData struct {
	HotelId int    `json:"hotelId"`
	Params  string `json:"params,omitempty"` // 1,2
}

// 酒店详情结构
type QueryHotelDetailResponse struct {
	HotelDetailList []struct {
		HotelId      int             `json:"hotelId"`
		HotelInfo    HotelStaticInfo `json:"hotelInfo"`
		RoomTypeList []RoomType      `json:"roomTypeList,omitempty"`
	} `json:"hotelDetailList"`
}

// 酒店信息结构
type HotelStaticInfo struct {
	HotelId           int    `json:"hotelId"`
	ThemeType         string `json:"themeType"`         // 酒店主题
	HotelNameCn       string `json:"hotelNameCn"`       // 酒店中文名称
	HotelNameEn       string `json:"hotelNameEn"`       // 酒店英文名称
	Star              string `json:"star"`              // 酒店星级
	AddressCn         string `json:"addressCn"`         // 中文地址
	AddressEn         string `json:"addressEn"`         // 英文地址
	CountryId         int    `json:"countryId"`         // 国家编号
	StateId           int    `json:"stateId"`           // 省份编号
	CityId            int    `json:"cityId"`            // 城市编号
	Area              int    `json:"area"`              // 行政区
	BusinessCircle    int    `json:"businessCircle"`    // 商业区
	Phone             string `json:"phone"`             // 酒店总机
	Longitude         string `json:"longitude"`         // 经度
	Latitude          string `json:"latitude"`          // 纬度
	Fax               string `json:"fax"`               // 传真
	Email             string `json:"email"`             // 邮箱
	PostCode          string `json:"postCode"`          // 邮编号
	CheckPolicy       string `json:"checkPolicy"`       // 酒店入离政策
	ChildrenPolicy    string `json:"childrenPolicy"`    // 儿童政策
	PetPolicy         string `json:"petPolicy"`         // 宠物政策
	EstablishmentDate string `json:"establishmentDate"` // 开业时间
	RenovationDate    string `json:"renovationDate"`    // 装修时间
	HotelGroup        string `json:"hotelGroup"`        // 集团
	HotelBrand        string `json:"hotelBrand"`        // 品牌
	Facilities        string `json:"facilities"`        // 酒店设施
	CardType          string `json:"cardType"`          // 信用卡
	IntroduceCn       string `json:"introduceCn"`       // 酒店中文介绍
	IntroduceEn       string `json:"introduceEn"`       // 酒店英文介绍
	SellType          int    `json:"sellType"`          // 是否热销酒店
	UpdateTime        string `json:"updateTime"`        // 修改时间
}

// 房型信息结构
type RoomType struct {
	RoomTypeId      string `json:"roomTypeId"`      // 房型编号
	RoomTypeCn      string `json:"roomTypeCn"`      // 客房中文名称
	RoomTypeEn      string `json:"roomTypeEn"`      // 客房英文名称
	BasisRoomId     int    `json:"basisroomid"`     // 基础房型ID
	BasisRoomCn     string `json:"basisroomCn"`     // 基础房型中文名
	BedType         string `json:"bedType"`         // 床型
	Maximize        int    `json:"maximize"`        // 最大入住人数
	Acreage         string `json:"acreage"`         // 房间面积
	BedWidth        string `json:"bedWidth"`        // 床大小
	FloorDistribute string `json:"floorDistribute"` // 楼层
	Facilities      string `json:"facilities"`      // 房型设施
	ExtraBedState   int    `json:"extraBedtState"`  // 是否允许加床
	BedCount        int    `json:"bedCount"`        // 加床数量
}

// 查询酒店详情
func QueryHotelDetail(requestData QueryHotelDetailRequestData) (*QueryHotelDetailResponse, error) {

	// 查询酒店详情的URL
	var queryHotelDetailUrl = baseURL + "/hotel/queryHotelDetail.json"

	// 创建 resty 客户端
	client := resty.New()

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	sign := generateSign(secretKey, appKey, timestamp)

	// 创建请求体
	req := Request[QueryHotelDetailRequestData]{
		Head: RequestHead{
			AppKey:    appKey,
			Timestamp: timestamp,
			Sign:      sign,
			Version:   "3.0.1",
		},
		Data: requestData,
	}

	// 将请求转换为JSON
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %v", err)
	}

	// 使用resty发送请求
	var response Response[QueryHotelDetailResponse]
	resp, err := client.R().
		SetQueryParam("reqData", string(jsonData)).
		SetResult(&response).
		Get(queryHotelDetailUrl)

	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v", err)
	}

	// 检查响应状态
	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("request failed with status code: %d", resp.StatusCode())
	}

	// 处理响应结果
	if response.Code == 0 {
		return &response.Result, nil
	} else {
		return nil, fmt.Errorf("Request failed with code %d: %s", response.Code, response.ErrorMsg)
	}
}
