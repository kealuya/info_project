package product

import (
	"fmt"
	"my_pilot/pkg/api_stb_shop"
)

// GetPageNum 获取品目池
func GetPageNum() (*api_stb_shop.Response[[]PageNumResult], error) {

	config := api_stb_shop.GetStbConfig()
	url := config["stb_shop"]["base_url"] + "goods/getPageNum"
	token := config["token"]["access_token"]
	client := api_stb_shop.GlobalClient.R()

	formData := map[string]string{
		"token": token,
	}

	resultResp := &api_stb_shop.Response[[]PageNumResult]{}
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

// PageNumResult 品目池结果结构体
type PageNumResult struct {
	// 品目名称
	Name string `json:"name"`
	// 品目编号
	PageNum string `json:"page_num"`
}
