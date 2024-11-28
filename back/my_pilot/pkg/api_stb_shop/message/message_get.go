package message

import (
	"fmt"
	"my_pilot/pkg/api_stb_shop"
)

func GetMessage(del, msgTypes string) (*api_stb_shop.Response[[]GetMessageResult], error) {

	config := api_stb_shop.GetStbConfig()
	url := config["stb_shop"]["base_url"] + "message/get"
	token := config["token"]["access_token"]
	client := api_stb_shop.GlobalClient.R()

	formData := map[string]string{
		"token": token,
		"del":   del,
		"type":  msgTypes,
	}

	resultResp := &api_stb_shop.Response[[]GetMessageResult]{}
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

// GetMessageResult 推送消息基础结构体
type GetMessageResult struct {
	ID     string `json:"id"`     // 推送ID
	Result any    `json:"result"` // 推送结果，根据type不同使用不同的结构体
	Type   int    `json:"type"`   // 推送类型
	Time   string `json:"time"`   // 推送时间
}
type PushMessage struct {
	// 订单拆分变更 type=1
	OrderSplit func(OrderSplitResult) error

	// 商品协议价格、史泰博价格变更 type=2
	PriceUpdate func(PriceUpdateResult) error

	// 商品库存变更 type=3
	StockUpdate func(StockUpdateResult) error

	// 商品上下架变更 type=4
	ProductStateChange func(ProductStateResult) error

	// 订单妥投状态 type=5
	OrderDelivery func(OrderDeliveryResult) error

	// 商品池变更 type=6
	ProductPool func(ProductPoolResult) error

	// 发票推送 type=7
	Invoice func(InvoiceResult) error

	// 挂起订单状态变更 type=8
	OrderStatusHold func(OrderStatusHoldResult) error

	// 商品可售区域变更 type=9
	ProductAreaChange func(ProductAreaChangeResult) error

	// 商品最小起订量变更 type=10
	ProductMinOrderChange func(ProductMinOrderResult) error

	// 订单核销消息 type=12
	OrderVerification func(OrderVerificationResult) error

	// 取消订单通知消息 type=13
	OrderCancel func(OrderCancelResult) error

	// 商品规格变更(联通福利商城) type=15
	ProductSeriesChange func(ProductSeriesResult) error

	// 中船订单妥投消息 type=16
	ShipOrderDelivery func(ShipOrderDeliveryResult) error
}

func (p PushMessage) RunLogicWithGetMessageResult(resp GetMessageResult) error {
	data := resp.Result
	switch resp.Type {
	case 1: // 订单拆分变更
		if result, ok := data.(OrderSplitResult); ok && p.OrderSplit != nil {
			return p.OrderSplit(result)
		}

	case 2: // 商品协议价格、史泰博价格变更
		if result, ok := data.(PriceUpdateResult); ok && p.PriceUpdate != nil {
			return p.PriceUpdate(result)
		}

	case 3: // 商品库存变更
		if result, ok := data.(StockUpdateResult); ok && p.StockUpdate != nil {
			return p.StockUpdate(result)
		}

	case 4: // 商品上下架变更
		if result, ok := data.(ProductStateResult); ok && p.ProductStateChange != nil {
			return p.ProductStateChange(result)
		}

	case 5: // 订单妥投状态
		if result, ok := data.(OrderDeliveryResult); ok && p.OrderDelivery != nil {
			return p.OrderDelivery(result)
		}

	case 6: // 商品池变更
		if result, ok := data.(ProductPoolResult); ok && p.ProductPool != nil {
			return p.ProductPool(result)
		}

	case 7: // 发票推送
		if result, ok := data.(InvoiceResult); ok && p.Invoice != nil {
			return p.Invoice(result)
		}

	case 8: // 挂起订单状态变更
		if result, ok := data.(OrderStatusHoldResult); ok && p.OrderStatusHold != nil {
			return p.OrderStatusHold(result)
		}

	case 9: // 商品可售区域变更
		if result, ok := data.(ProductAreaChangeResult); ok && p.ProductAreaChange != nil {
			return p.ProductAreaChange(result)
		}

	case 10: // 商品最小起订量变更
		if result, ok := data.(ProductMinOrderResult); ok && p.ProductMinOrderChange != nil {
			return p.ProductMinOrderChange(result)
		}

	case 12: // 订单核销消息
		if result, ok := data.(OrderVerificationResult); ok && p.OrderVerification != nil {
			return p.OrderVerification(result)
		}

	case 13: // 取消订单通知消息
		if result, ok := data.(OrderCancelResult); ok && p.OrderCancel != nil {
			return p.OrderCancel(result)
		}

	case 15: // 商品规格变更(联通福利商城)
		if result, ok := data.(ProductSeriesResult); ok && p.ProductSeriesChange != nil {
			return p.ProductSeriesChange(result)
		}

	case 16: // 中船订单妥投消息
		if result, ok := data.(ShipOrderDeliveryResult); ok && p.ShipOrderDelivery != nil {
			return p.ShipOrderDelivery(result)
		}

	default:
		return fmt.Errorf("unknown message type: %d", resp.Type)
	}

	return nil
}

// OrderSplitResult 订单拆分变更 type=1
type OrderSplitResult struct {
	POrderID string `json:"pOrderId"` // 父订单ID
}

// PriceUpdateResult 商品协议价格、史泰博价格变更 type=2
type PriceUpdateResult struct {
	SkuID string `json:"skuId"` // 商品编号
}

// StockUpdateResult 商品库存变更 type=3
type StockUpdateResult struct {
	SkuID string `json:"skuId"` // 商品编号
}

// ProductStateResult 商品上下架变更 type=4
type ProductStateResult struct {
	SkuID string `json:"skuId"` // 商品编号
	State int    `json:"state"` // 0是下架，1是上架
}

// OrderDeliveryResult 订单妥投状态 type=5
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

// OrderStatusHoldResult 挂起订单状态变更 type=8
type OrderStatusHoldResult struct {
	OrderID string `json:"orderId"` // 订单编号
}

// ProductAreaChangeResult 商品可售区域变更 type=9
type ProductAreaChangeResult struct {
	SkuID string `json:"skuId"` // 商品编号
}

// ProductMinOrderResult 商品最小起订量变更 type=10
type ProductMinOrderResult struct {
	SkuID string `json:"skuId"` // 商品编号
}

// OrderVerificationResult 订单核销消息 type=12
type OrderVerificationResult struct {
	OrderID string `json:"orderId"` // 订单编号
}

// OrderCancelResult 取消订单通知消息 type=13
type OrderCancelResult struct {
	OrderID string `json:"orderId"` // 订单编号
}

// ProductSeriesResult 商品规格变更(联通福利商城) type=15
type ProductSeriesResult struct {
	SeriesID string `json:"seriesID"` // 商品组号
	State    string `json:"state"`    // 1添加，2更新，3删除
}

// ShipOrderDeliveryResult 中船订单妥投消息 type=16
type ShipOrderDeliveryResult struct {
	OrderID string `json:"orderId"` // 史泰博订单编号
	State   string `json:"state"`   // 1是妥投，2是拒收
}
