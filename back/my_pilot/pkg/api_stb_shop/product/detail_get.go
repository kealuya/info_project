package product

import (
	"fmt"
	"my_pilot/pkg/api_stb_shop"
)

// GetDetail 查询商品详情
func GetDetail(sku string) (*api_stb_shop.Response[SkuDetailResult], error) {

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

// SkuDetailResult 商品详情结果结构体
type SkuDetailResult struct {
	// Product 商品信息结构体
	Sku          string `json:"sku"`          // 商品编号（带单位）
	SkuCode      string `json:"skuCode"`      // 商品编号(不带单位)
	Weight       string `json:"weight"`       // 重量
	ImagePath    string `json:"imagePath"`    // 主图地址
	State        string `json:"state"`        // 上下架状态(1为上架，0为下架)
	ProductArea  string `json:"productArea"`  // 产地
	BrandID      string `json:"brandId"`      // 品牌编码
	BrandName    string `json:"brandName"`    // 品牌
	Name         string `json:"name"`         // 商品名称
	UPC          string `json:"upc"`          // 条形码
	SaleUnit     string `json:"saleUnit"`     // 销售单位
	Category     string `json:"category"`     // 最后一级分类ID
	Introduction string `json:"Introduction"` // 详情介绍
	Param        string `json:"param"`        // 规格参数(JSON格式)
	WareQD       string `json:"wareQD"`       // 包装清单
	GoodsModel   string `json:"goodsModel"`   // 商品型号
	Subtitle     string `json:"subtitle"`     // 副标题
	PmTaxSorts   string `json:"pmTaxSorts"`   // 税收分类编码
}
