package product

import (
	"fmt"
	"my_pilot/pkg/api_jd_iop"
	"strings"
)

// IsCheckProduct  验证商品可售性
func IsCheckProduct(skuId []string, queryExts string) (*api_jd_iop.Response[DetailIsCheckResult], error) {

	config := api_jd_iop.GetJdIopConfig()
	url := config["jd_iop"]["base_url"] + "product/check"
	token := config["token"]["access_token"]
	client := api_jd_iop.GlobalClient.R()

	// 构造请求参数
	//queryExts 扩展参数：英文逗号间隔输入
	//noReasonToReturn //无理由退货类型
	//thwa //无理由退货文案类型
	//isSelf // 是否自营
	//isJDLogistics // 是否京东配送
	//taxInfo //商品税率
	formData := map[string]string{
		"token":     token,
		"skuIds":    strings.Join(skuId, ","),
		"queryExts": queryExts,
	}

	resultResp := &api_jd_iop.Response[DetailIsCheckResult]{}
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

// DetailIsCheckResult 商品可售状态结构体
type DetailIsCheckResult struct {
	// 商品编号
	SkuID int64 `json:"skuId"`
	// 商品名称
	Name string `json:"name"`
	// 是否可售，1：是，0：否
	SaleState int `json:"saleState"`
	// 是否可开专票，1：支持，0：不支持
	IsCanVAT int `json:"isCanVAT"`
	// 无理由退货类型
	// 0、3：不支持7天无理由退货
	// 1、5、8或null：支持7天无理由退货
	// 2：支持90天无理由退货
	// 4、7：支持15天无理由退货
	// 6：支持30天无理由退货
	NoReasonToReturn int `json:"noReasonToReturn"`
	// 无理由退货文案类型
	// null/0：文案空
	// 1：支持7天无理由退货
	// 2：支持7天无理由退货（拆封后不支持）
	// 3：支持7天无理由退货（激活后不支持）
	// 4：支持7天无理由退货（使用后不支持）
	// 5：支持7天无理由退货（安装后不支持）
	// 12：支持15天无理由退货
	// 13：支持15天无理由退货（拆封后不支持）
	// 14：支持15天无理由退货（激活后不支持）
	// 15：支持15天无理由退货（使用后不支持）
	// 16：支持15天无理由退货（安装后不支持）
	// 22：支持30天无理由退货
	// 23：支持30天无理由退货（安装后不支持）
	// 24：支持30天无理由退货（拆封后不支持）
	// 25：支持30天无理由退货（使用后不支持）
	// 26：支持30天无理由退货（激活后不支持）
	Thwa int `json:"thwa"`
}
