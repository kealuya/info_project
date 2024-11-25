package product

import (
	"fmt"
	"github.com/gohouse/t"
	"my_pilot/pkg/api_jd_iop"
)

// QuerySkuByPage   查询商品池编号
func QuerySkuByPage(pageNum string, pageSize int, offset string) (*api_jd_iop.Response[SkuByPageResult], error) {

	config := api_jd_iop.GetJdIopConfig()
	url := config["jd_iop"]["base_url"] + "product/querySkuByPage"
	token := config["token"]["access_token"]
	client := api_jd_iop.GlobalClient.R()

	// 构造请求参数
	formData := map[string]string{
		"token":    token,
		"pageNum":  pageNum,
		"pageSize": t.New(pageSize).String(), //每页大小，默认20，最大1000。建议 50 ~200
		"offset":   offset,                   //偏移量，池id的首次查询传0，相同池ID的上次请求结果中skuIds中的最后一个skuId
	}

	resultResp := &api_jd_iop.Response[SkuByPageResult]{}
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
	Skus   []string `json:"skus"`   // skuId集合
	Offset string   `json:"offset"` // 在下一次查询时使用（偏移量）
}
