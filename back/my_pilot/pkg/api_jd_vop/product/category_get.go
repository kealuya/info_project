package product

import (
	"fmt"
	"github.com/gohouse/t"
	"my_pilot/pkg/api_jd_vop"
)

// GetCategory   查询分类信息
func GetCategory(cid string) (*api_jd_vop.Response[CategoryResult], error) {

	config := api_jd_vop.GetJdVopConfig()
	url := config["jd_vop"]["base_url"] + "product/getCategory"
	token := config["token"]["access_token"]
	client := api_jd_vop.GlobalClient.R()

	// 构造请求参数
	formData := map[string]string{
		"token": token,
		"cid":   t.New(cid).String(),
	}

	resultResp := &api_jd_vop.Response[CategoryResult]{}
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

// CategoryResult   分类信息
type CategoryResult struct {
	CatId    int    `json:"catId"`    //分类ID
	ParentId int    `json:"parentId"` //父分类ID
	Name     string `json:"name"`     //分类名称
	CatClass int    `json:"catClass"` //0：一级分类；1：二级分类；2：三级分类；
	State    int    `json:"state"`    //1：有效；0：无效；
}
