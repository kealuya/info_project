package models

import (
	"fmt"
	"sdykdx_open_data/common"
	"sdykdx_open_data/internal"
)

const OrganizationServiceID = "1810234926816927744"

func ObtainOrganization(dwdm string) (organization *Organization, bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		organization = nil
		return
	})

	// 创建请求体数据
	requestData := internal.RequestBody{
		ServiceID: OrganizationServiceID,
		Params: map[string]interface{}{
			"DWDM": dwdm,
			//"page": 1,
			//"size": 20,
		},
	}
	// 发送请求
	response := internal.CommonRequest[Organization](requestData)
	if !response.Success {
		return nil, fmt.Errorf("api接口调用失败")
	}
	if len(response.Data.Data) == 0 {
		return nil, fmt.Errorf("查询无结果")
	}
	if len(response.Data.Data) != 1 {
		return nil, fmt.Errorf("返回数据长度不为1")
	}

	return &response.Data.Data[0], nil
}

type Organization struct {
	WID  string `json:"WID"`  // 主键
	DWDM string `json:"DWDM"` // 单位代码
	DWMC string `json:"DWMC"` // 单位名称
	FID  string `json:"FID"`  // 父级ID
	CC   string `json:"CC"`   // 层次
	PXH  string `json:"PXH"`  // 排序号
}
