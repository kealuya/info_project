package product

import (
	"fmt"
	"my_pilot/pkg/api_lxwl_shop"
)

// GetSkus 获取商品编号sku
func GetSkus(categoryNo string) (*api_lxwl_shop.Response[[]string], error) {

	config := api_lxwl_shop.GetLxwlConfig()
	url := config["lxwl_shop"]["base_url"] + "product/skus"
	token := config["token"]["access_token"]
	client := api_lxwl_shop.GlobalClient.R()

	formData := map[string]string{
		"accessToken": token,
		"categoryNo":  categoryNo,
	}

	resultResp := &api_lxwl_shop.Response[[]string]{}
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
