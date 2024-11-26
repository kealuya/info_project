package product

import (
	"fmt"
	"my_pilot/pkg/api_stb_shop"
)

// GetSkus 获取商品编号sku
func GetSkus(pageNum string) (*api_stb_shop.Response[string], error) {

	config := api_stb_shop.GetStbConfig()
	url := config["stb_shop"]["base_url"] + "goods/product/getSku"
	token := config["token"]["access_token"]
	client := api_stb_shop.GlobalClient.R()

	formData := map[string]string{
		"token":   token,
		"pageNum": pageNum,
	}

	resultResp := &api_stb_shop.Response[string]{}
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
