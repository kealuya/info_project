package product

import (
	"fmt"
	"github.com/shopspring/decimal"
	"my_pilot/pkg/api_lxwl_shop"
)

// GetDetail 查询商品详情
func GetDetail(skus []string) (*api_lxwl_shop.Response[[]SkuDetailResult], error) {

	config := api_lxwl_shop.GetLxwlConfig()
	url := config["lxwl_shop"]["base_url"] + "product/detail"
	token := config["token"]["access_token"]
	client := api_lxwl_shop.GlobalClient.R()

	formData := map[string]string{
		"accessToken": token,
	}

	resultResp := &api_lxwl_shop.Response[[]SkuDetailResult]{}
	resp, errClient := client.SetHeader("Content-Type", "application/json").
		SetQueryParams(formData).
		SetResult(resultResp).
		SetBody(skus).
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
	SkuNo                 string              `json:"skuNo"`                 // 商品编号
	SkuName               string              `json:"skuName"`               // 商品名称
	BrandName             string              `json:"brandName"`             // 品牌名称
	TaxRate               decimal.Decimal     `json:"taxRate"`               // 税率
	TaxCode               string              `json:"taxCode"`               // 税收分类编码
	CategoryCode          string              `json:"categoryCode"`          // 领先分类编码
	CategoryName          string              `json:"categoryName"`          // 领先分类名称
	ThirdCategoryCode     string              `json:"thirdCategoryCode"`     // 三方分类编码
	ThirdCategoryName     string              `json:"thirdCategoryName"`     // 三方分类名称
	WebsitePrice          decimal.Decimal     `json:"websitePrice"`          // 官网售价（含税）
	MallPrice             decimal.Decimal     `json:"mallPrice"`             // 商城售价（含税）
	Inventory             int64               `json:"inventory"`             // 库存
	DiscountRate          decimal.Decimal     `json:"discountRate"`          // 折扣率
	Barcode               string              `json:"barcode"`               // 69码/条形码
	Unit                  string              `json:"unit"`                  // 单位
	PrimaryAttributeName  string              `json:"primaryAttributeName"`  // 型号/规格
	PrimaryAttributeValue string              `json:"primaryAttributeValue"` // 型号/规格值
	Weight                float64             `json:"weight"`                // 重量（含包装）
	Origin                string              `json:"origin"`                // 产地
	ShelfLife             int                 `json:"shelfLife"`             // 保质期
	ShelfLifeUnit         string              `json:"shelfLifeUnit"`         // 保质期单位（天，月，年）
	Warranty              string              `json:"warranty"`              // 质保
	ExecutiveStandard     string              `json:"executiveStandard"`     // 执行标准
	PackingList           string              `json:"packingList"`           // 包装清单
	Slogan                string              `json:"slogan"`                // 广告语
	Length                int                 `json:"length"`                // 长
	Width                 int                 `json:"width"`                 // 宽
	Height                int                 `json:"height"`                // 高
	PackingSpecification  string              `json:"packingSpecification"`  // 包装规格
	SkuAttributeGroupList []SkuAttributeGroup `json:"skuAttributeGroupList"` // 商品属性组集合
	MainPicture           SkuImage            `json:"mainPicture"`           // 主图
	DetailPicture         []SkuImage          `json:"detailPicture"`         // 详情图
	SlidePicture          []SkuImage          `json:"slidePicture"`          // 轮播图
}

// SkuAttributeGroup 商品属性组
type SkuAttributeGroup struct {
	AttributeGroupName string         `json:"attributeGroupName"` // 属性组名字
	AttributeList      []SkuAttribute `json:"attributeList"`      // 属性值集合
}

// SkuAttribute 商品属性
type SkuAttribute struct {
	AttributeName  string `json:"attributeName"`  // 属性名
	AttributeValue string `json:"attributeValue"` // 属性值
}

// SkuImage 商品图片
type SkuImage struct {
	Path string `json:"path"` // 图片路径
}
