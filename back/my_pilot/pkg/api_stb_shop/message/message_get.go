package message

import (
	"fmt"
	"my_pilot/pkg/api_stb_shop"
)

func GetMessage(del int, msgType string) (*api_stb_shop.Response[SkuDetailResult], error) {

	config := api_stb_shop.GetStbConfig()
	url := config["stb_shop"]["base_url"] + "goodsModel/product/getDetail"
	token := config["token"]["access_token"]
	client := api_stb_shop.GlobalClient.R()

	formData := map[string]string{
		"token": token,
		"sku":   sku,
	}

	resultResp := &api_stb_shop.Response[SkuDetailResult]{}
	resp, errClient := client.SetHeader("Content-Type", "application/json").
		SetQueryParams(formData).
		SetResult(resultResp).
		Post(url)
	if errClient != nil {
		return nil, errClient
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}
	return resultResp, nil
}

// PushMessage 推送消息基础结构体
type PushMessage struct {
	ID     string      `json:"id"`     // 推送ID
	Result interface{} `json:"result"` // 推送结果，根据type不同使用不同的结构体
	Type   int         `json:"type"`   // 推送类型
	Time   string      `json:"time"`   // 推送时间
}

// OrderSplitResult 订单拆分变更 type=1
type OrderSplitResult struct {
	POrderID string `json:"pOrderId"` // 父订单ID
}

// PriceChangeResult 商品价格变更 type=2,3
type PriceChangeResult struct {
	SkuID string `json:"skuId"` // 商品编号
}

// ProductStateResult 商品上下架变更 type=4
type ProductStateResult struct {
	SkuID string `json:"skuId"` // 商品编号
	State int    `json:"state"` // 0是下架，1是上架
}

// OrderDeliveryResult 订单妥投状态 type=5,16
type OrderDeliveryResult struct {
	OrderID string `json:"orderId"` // 史泰博订单编号
	State   string `json:"state"`   // 1是妥投，2是拒收
}

// ProductPoolResult 商品池变更 type=6
type ProductPoolResult struct {
	SkuID string `json:"skuId"` // 商品编号
	State string `json:"state"` // 1添加，2删除，3更新
}

// InvoiceResult 发票推送 type=7
type InvoiceResult struct {
	MarkID        string `json:"markId"`        // 第三方申请开发票的唯一标识
	Success       bool   `json:"success"`       // 是否成功
	ResultMessage string `json:"resultMessage"` // 开票结果信息
	InvoiceID     string `json:"invoiceId"`     // 发票号
}

// OrderStatusResult 订单状态变更 type=8,12,13
type OrderStatusResult struct {
	OrderID string `json:"orderId"` // 订单编号
}

// ProductAreaResult 商品可售区域变更 type=9,10
type ProductAreaResult struct {
	SkuID string `json:"skuId"` // 商品编号
}

// ProductSeriesResult 商品规格变更 type=15
type ProductSeriesResult struct {
	SeriesID string `json:"seriesID"` // 商品组号
	State    string `json:"state"`    // 1添加，2更新，3删除
}
