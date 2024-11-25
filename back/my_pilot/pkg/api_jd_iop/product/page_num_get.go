package product

import (
	"fmt"
	"my_pilot/pkg/api_jd_iop"
)

// GetPageNum 查询池内商品编号
func GetPageNum() (*api_jd_iop.Response[PageNumResult], error) {

	config := api_jd_iop.GetJdIopConfig()
	url := config["jd_iop"]["base_url"] + "product/getPageNum"
	token := config["token"]["access_token"]
	client := api_jd_iop.GlobalClient.R()

	// 构造请求参数
	formData := map[string]string{
		"token": token,
	}

	resultResp := &api_jd_iop.Response[PageNumResult]{}
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

// PageNumResult 查询商品池编号
type PageNumResult struct {
	Name    string `json:"name"`     // 商品池名称
	PageNum string `json:"page_num"` // 商品池编号

}
