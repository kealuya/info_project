package product

import (
	"fmt"
	"github.com/gohouse/t"
	"my_pilot/pkg/api_jd_vop"
)

// QuerySkuByPage   查询商品池编号
func QuerySkuByPage(pageNum string, pageNo int) (*api_jd_vop.Response[SkuByPageResult], error) {

	config := api_jd_vop.GetJdVopConfig()
	url := config["jd_vop"]["base_url"] + "product/getSkuByPage"
	token := config["token"]["access_token"]
	client := api_jd_vop.GlobalClient.R()

	// 构造请求参数
	formData := map[string]string{
		"token":   token,
		"pageNum": t.New(pageNum).String(),
		"pageNo":  t.New(pageNo).String(),
	}

	resultResp := &api_jd_vop.Response[SkuByPageResult]{}
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

// SkuByPageResult   查询商品池编号
type SkuByPageResult struct {
	SkuIds    []int `json:"skuIds"`    // skuId集合
	PageCount int   `json:"pageCount"` //
}
