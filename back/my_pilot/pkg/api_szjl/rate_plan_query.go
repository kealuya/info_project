package api_szjl

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

// RoomGroup 房间入住信息
type RoomGroup struct {
	Adults          int              `json:"adults,omitempty"`    // 成人数
	Children        int              `json:"children,omitempty"`  // 儿童数
	ChildAges       string           `json:"childAges,omitempty"` // 儿童年龄
	CheckInPersions []CheckInPersion `json:"checkInPersions"`     // 儿童年龄
}

// QueryRatePlanRequestData 报价查询请求参数
type QueryRatePlanRequestData struct {
	HotelId      int         `json:"hotelId"`      // 酒店编号
	CheckInDate  string      `json:"checkInDate"`  // 入住日期
	CheckOutDate string      `json:"checkOutDate"` // 离店日期
	RoomGroups   []RoomGroup `json:"roomGroups"`   // 房间信息
	QueryType    int         `json:"queryType"`    // 数据类型
}

// NightlyRate 间夜价格信息
type NightlyRate struct {
	FormulaType     int     `json:"formulaType"`     // 配额类型
	Date            string  `json:"date"`            // 日期
	Cost            float64 `json:"cose"`            // 价格
	Status          int     `json:"status"`          // 房态
	CurrentAlloment int     `json:"currentAlloment"` // 库存
	Breakfast       int     `json:"breakfast"`       // 早餐
	BookingRuleId   string  `json:"bookingRuleId"`   // 预订条款编号
	RefundRuleId    string  `json:"refundRuleId"`    // 取消条款编号
}

// Promotion 礼包信息
type Promotion struct {
	MarketingName     string  `json:"marketingName"`     // 促销名称
	StartDate         string  `json:"startDate"`         // 礼包开始时间
	EndDate           string  `json:"endDate"`           // 礼包结束时间
	Description       string  `json:"description"`       // 礼包描述
	MarketingPrice    float64 `json:"marketingPrice"`    // 礼包价值
	TimeLimitType     int     `json:"timeLimitType"`     // 时段类型
	WeekEffectiveDate string  `json:"weekEffectiveDate"` // 周有效期
	PromotionMethod   int     `json:"promotionMethod"`   // 促销方式
	PromotionType     int     `json:"promotionType"`     // 促销类型
}

// RatePlan 产品信息
type RatePlan struct {
	KeyId         string        `json:"keyId"`         // 产品编号
	SupplierId    int           `json:"supplierId"`    // 供应商ID
	KeyName       string        `json:"keyName"`       // 房型名称
	BedName       string        `json:"bedName"`       // 床型名称
	MaxOccupancy  int           `json:"maxOccupancy"`  // 最大入住人数
	Currency      string        `json:"currency"`      // 币种
	RateTypeId    string        `json:"rateTypeId"`    // 价格类型编号
	PaymentType   int           `json:"paymentType"`   // 付款类型
	Breakfast     int           `json:"breakfast"`     // 早餐
	IfInvoice     int           `json:"ifInvoice"`     // 是否开票
	NightlyRates  []NightlyRate `json:"nightlyRates"`  // 间夜价格数组
	BookingRuleId string        `json:"bookingRuleId"` // 预订条款编号
	RefundRuleId  string        `json:"refundRuleId"`  // 取消条款编号
	Market        string        `json:"market"`        // 适用市场
	BanOta        string        `json:"banota"`        // 禁售渠道
	Promotions    []Promotion   `json:"promotions"`    // 礼包信息
}

// Room 房型信息
type Room struct {
	RoomTypeId string     `json:"roomTypeId"` // 房型编号
	RatePlans  []RatePlan `json:"ratePlans"`  // 产品数组
}

// BookingRule 预订条款信息
type BookingRule struct {
	BookingRuleId  string `json:"bookingRuleId"`  // 预订条款编号
	StartDate      string `json:"startDate"`      // 开始日期
	EndDate        string `json:"endDate"`        // 结束日期
	MinAmount      int    `json:"minAmount"`      // 预订最少数量
	MaxAmount      int    `json:"maxAmount"`      // 预订最多数量
	MinDays        int    `json:"minDays"`        // 最少入住天数
	MaxDays        int    `json:"maxDays"`        // 最多入住天数
	MinAdvHours    int    `json:"minAdvHours"`    // 最少提前预订时间
	MaxAdvHours    int    `json:"maxAdvHours"`    // 最大提前预订时间
	WeekSet        string `json:"weekSet"`        // 有效星期
	StartTime      string `json:"startTime"`      // 每日开始销售时间
	EndTime        string `json:"endTime"`        // 每日结束销售时间
	BookingNotices string `json:"bookingNotices"` // 预订说明
}

// RefundRule 取消条款信息
type RefundRule struct {
	RefundRuleId              string `json:"refundRuleId"`              // 取消条款编号
	RefundRuleType            int    `json:"refundRuleType"`            // 取消条款规则
	ThirtyMinFreeCancelPolicy bool   `json:"thirtyMinFreeCancelPolicy"` // 是否支持30分钟免费取消
	RefundRuleHours           int    `json:"refundRuleHours"`           // 入住前n小时
	DeductType                int    `json:"deductType"`                // 取消客人罚金
}

// HotelRatePlan 酒店报价信息
type HotelRatePlan struct {
	HotelId      int           `json:"hotelId"`      // 酒店编号
	Rooms        []Room        `json:"rooms"`        // 房型数组
	BookingRules []BookingRule `json:"bookingRules"` // 预订条款数组
	RefundRules  []RefundRule  `json:"refundRules"`  // 取消条款数组
}

// QueryRatePlanResult 报价查询响应结果
type QueryRatePlanResult struct {
	HotelRatePlanList []HotelRatePlan `json:"hotelRatePlanList"` // 酒店产品数组
}

func QueryRatePlan(requestData QueryRatePlanRequestData) (*QueryRatePlanResult, error) {
	var queryRatePlanUrl = baseURL + "/hotel/queryRatePlan.json"

	// 创建 resty 客户端
	client := resty.New()

	// 生成时间戳
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)

	// 生成签名
	sign := generateSign(secretKey, appKey, timestamp)

	// 创建请求体
	req := Request[QueryRatePlanRequestData]{
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
	var response Response[QueryRatePlanResult]
	resp, err := client.R().
		SetQueryParam("reqData", string(jsonData)).
		SetResult(&response).
		Get(queryRatePlanUrl)

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
