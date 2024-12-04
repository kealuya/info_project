package product

import (
	"fmt"
	"github.com/gohouse/t"
	"my_pilot/pkg/api_jd_vop"
)

// GetDetail  查询商品详情
func GetDetail(sku int, queryExts string) (*api_jd_vop.Response[DetailResult], error) {

	config := api_jd_vop.GetJdVopConfig()
	url := config["jd_vop"]["base_url"] + "product/getDetail"
	token := config["token"]["access_token"]
	client := api_jd_vop.GlobalClient.R()

	// 构造请求参数
	// queryExts 以下为商品维度扩展字段，当入参输入某个扩展字段后，出参会返回该字段对应的出参。
	//可以根据需要选用。
	//nappintroduction//移动端商品详情大字段
	//nintroduction//PC端商品详情大字段
	//wxintroduction//微信小程序商品详情大字段，仅提供图片地址，需要客户添加显示逻辑
	//contractSkuExt(已废弃)//获取客户侧分类编号，需要京东运营维护京东SKU与客户分类编号的映射
	//isFactoryShip//是否厂直商品  materielCode//物料编码
	//isEnergySaving//是否节能环保商品 taxCode//京东侧税收分类编码
	//LowestBuy//商品最低起购量
	//capacity//容量单位转换（例如：油品单位桶转升）
	//spuId//京东侧模拟SPU号
	//pName//SPU名称
	//isJDLogistics// "是否京东配送"
	//taxInfo//"商品税率"
	//upc69//"商品69条码"
	//ChinaCatalog//中国法分类（仅限图书商品使用）
	//productFeatures//图书产品特色
	//seoModel//规格参数
	//paramDetailJson//获取结构化商品属性数据(同京东官网样式
	//paramJson//转商详接口出参param为json格式(只解析原出参param)
	//wserve//质保信息 isSelf//是否自营
	//categoryAttrs//商品属性
	//saleAttr//商品颜色尺码
	//sizeDesc//尺寸描述：长：length、宽：width、高：height
	//例：appintroduce,nappintroduction,nintroduction,wxintroduction,shouhou,contractSkuExt,
	//isEnergySaving,taxCode,LowestBuy,capacity,spuId,pName,isJDLogistics,taxInfo,upc69,
	//ChinaCatalog,contractSkuPoolExt,productFeatures,seoModel,paramDetailJson,paramJson,
	//wserve,categoryAttrs,isSelf,isFactoryShip,materielCode,customerAttrs,standBrandId,
	//model,qualityLevel,encapsForm,originInfo,warranty

	formData := map[string]string{
		"token":     token,
		"sku":       t.New(sku).String(),
		"queryExts": queryExts,
	}

	resultResp := &api_jd_vop.Response[DetailResult]{}
	resp, errClient := client.SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		SetResult(resultResp).
		ForceContentType("application/json").
		Post(url)
	if errClient != nil {
		return nil, errClient
	}
	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}
	return resultResp, nil
}

// DetailResult 商品实物结果结构体
type DetailResult struct {
	// 售卖单位
	SaleUnit string `json:"saleUnit"`
	// 重量
	Weight string `json:"weight"`
	// 产地
	ProductArea string `json:"productArea"`
	// 包装清单
	WareQD string `json:"wareQD"`
	// 主图
	ImagePath string `json:"imagePath"`
	// 规格参数
	Param string `json:"param"`
	// 主站上下架状态 (1上架 0下架)
	State int `json:"state"`
	// 商品编号
	Sku int `json:"sku"`
	// 品牌名称
	BrandName string `json:"brandName"`
	// UPC码区分实物、图书、音像、三种场景
	UPC string `json:"upc"`
	// 分类示例"670;729;4837"
	Category string `json:"category"`
	// 商品名称
	Name string `json:"name"`
	// 商品详情页大字段
	Introduction string `json:"introduction"`
}
