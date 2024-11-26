package product

import (
	"fmt"
	"my_pilot/pkg/api_lxwl_shop"
)

// GetPools 获取品目池
func GetPools() (*api_lxwl_shop.Response[[]CategoryResult], error) {

	config := api_lxwl_shop.GetLxwlConfig()
	url := config["lxwl_shop"]["base_url"] + "product/pools"
	token := config["token"]["access_token"]
	client := api_lxwl_shop.GlobalClient.R()

	formData := map[string]string{
		"accessToken": token,
	}

	resultResp := &api_lxwl_shop.Response[[]CategoryResult]{}
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

// CategoryResult 品目池结果结构体
type CategoryResult struct {
	// 品目名称
	Name string `json:"name"`
	// 品目编号
	No string `json:"no"`
}
