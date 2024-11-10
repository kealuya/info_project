package api_szjl

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"sync"
	"time"
)

// QueryHotelDetailRequestData 请求体结构
type QueryHotelDetailRequestData struct {
	HotelId int    `json:"hotelId"`
	Params  string `json:"params,omitempty"` // 1,2
}

// QueryHotelDetailResponse 酒店详情结构
type QueryHotelDetailResponse struct {
	HotelDetailList []struct {
		HotelId      int             `json:"hotelId"`
		HotelInfo    HotelStaticInfo `json:"hotelInfo"`
		RoomTypeList []RoomType      `json:"roomTypeList,omitempty"`
		RateTypeList []RateType      `json:"rateTypeList"`
		ImageList    []Image         `json:"imageList"`
	} `json:"hotelDetailList"`
}

// RateType 价格类型信息
type RateType struct {
	RateTypeId int    `json:"rateTypeId"` // 价格类型编号
	RateTypeCn string `json:"rateTypeCn"` // 价格类型中文名
	RateTypeEn string `json:"rateTypeEn"` // 价格类型英文名
}

// Image 图片信息
type Image struct {
	ImageId     int    `json:"imageId"`     // 图片编号
	Type        string `json:"type"`        // 图片类型，见常量列表
	RoomTypeIds string `json:"roomTypeIds"` // 图片关联房型，多个逗号分隔
	ThumbUrl    string `json:"thumbUrl"`    // 缩略图地址
	ImageUrl    string `json:"imageUrl"`    // 图片地址
	ImageLogo   int    `json:"imageLogo"`   // 是否带水印(0:有水印logo 1:无水印logo)
	ImageSize   int    `json:"imageSize"`   // 图片规格(0:未分类、1:350*350、2:550*412、3:640*960)
}

// HotelStaticInfo 酒店信息结构
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

// RoomType 房型信息结构
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

var (
	// 创建全局 resty 客户端，避免重复创建
	globalClient = resty.New().
			SetTimeout(20 * time.Second). // 设置超时
			SetRetryCount(2).             // 设置重试次数
			SetRetryWaitTime(1000 * time.Millisecond).
			SetRetryMaxWaitTime(5000 * time.Millisecond).
			SetTransport(&http.Transport{
			MaxIdleConnsPerHost:   100,              // 每个host最大空闲连接数
			MaxConnsPerHost:       200,              // 每个host最大连接数
			IdleConnTimeout:       90 * time.Second, // 空闲连接超时时间
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    false, // 如果不需要压缩可以禁用
		})
	// 初始化 sonic 编码器
	jsonEncoder = sonic.ConfigDefault
	// 使用对象池来减少内存分配
	requestPool = sync.Pool{
		New: func() interface{} {
			return &Request[QueryHotelDetailRequestData]{}
		},
	}
)

// QueryHotelDetail 查询酒店详情
func QueryHotelDetail(requestData QueryHotelDetailRequestData) (*QueryHotelDetailResponse, error) {

	// 查询酒店详情的URL
	var queryHotelDetailUrl = baseURL + "/hotel/queryHotelDetail.json"

	// 创建 resty 客户端
	//client := resty.New()

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	sign := generateSign(secretKey, appKey, timestamp)
	// 从对象池获取请求对象
	req := requestPool.Get().(*Request[QueryHotelDetailRequestData])
	defer requestPool.Put(req)
	// 创建请求体
	*req = Request[QueryHotelDetailRequestData]{
		Head: RequestHead{
			AppKey:    appKey,
			Timestamp: timestamp,
			Sign:      sign,
			Version:   "3.0.1",
		},
		Data: requestData,
	}

	// 使用 sonic 序列化
	jsonData, err := jsonEncoder.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON: %v", err)
	}

	// 使用resty发送请求
	var response Response[QueryHotelDetailResponse]
	resp, err := globalClient.R().
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
	}
	return nil, fmt.Errorf("Request failed with code %d: %s", response.Code, response.ErrorMsg)

}
