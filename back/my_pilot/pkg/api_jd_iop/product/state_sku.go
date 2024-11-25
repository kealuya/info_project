package product

import (
	"fmt"
	"my_pilot/pkg/api_jd_iop"
	"strings"
)

// StateSku  查询商品上下架状态
func StateSku(sku []string) (*api_jd_iop.Response[[]SkuStateResult], error) {

	config := api_jd_iop.GetJdIopConfig()
	url := config["jd_iop"]["base_url"] + "product/skuState"
	token := config["token"]["access_token"]
	client := api_jd_iop.GlobalClient.R()

	// 构造请求参数
	formData := map[string]string{
		"token": token,
		"sku":   strings.Join(sku, ","),
	}

	resultResp := &api_jd_iop.Response[[]SkuStateResult]{}
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

// SkuStateResult 商品实物结果结构体
type SkuStateResult struct {
	// 1：上架，0：下架
	State int `json:"state"`
	// 商品编号
	Sku int64 `json:"sku"`
}
