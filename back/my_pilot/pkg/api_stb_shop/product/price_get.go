package product

import (
	"fmt"
	"my_pilot/pkg/api_stb_shop"
)

// GetPrice 查询商品价格
func GetPrice(sku string) (*api_stb_shop.Response[[]PriceResult], error) {

	config := api_stb_shop.GetStbConfig()
	url := config["stb_shop"]["base_url"] + "goods/product/getPrice"
	token := config["token"]["access_token"]
	client := api_stb_shop.GlobalClient.R()

	formData := map[string]string{
		"token": token,
		"sku":   sku,
	}

	resultResp := &api_stb_shop.Response[[]PriceResult]{}
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

// PriceResult 商品价格结果结构体
type PriceResult struct {
	SkuID         string `json:"skuId"`         // 商品编号
	BizPrice      string `json:"bizPrice"`      // 企业价(协议裸价+协议税价)
	SupplierPrice string `json:"supplierPrice"` // 史泰博市场价
	BizNakedPrice string `json:"bizNakedPrice"` // 协议裸价
	BizTaxPrice   string `json:"bizTaxPrice"`   // 协议税价（仅税额部分）
	TaxRate       string `json:"taxRate"`       // 税率（保留两位小数）
}
