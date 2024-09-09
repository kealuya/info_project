package models

import (
	"fmt"
	"sdykdx_open_data/common"
	"sdykdx_open_data/internal"
)

const ClinicalTeacherServiceID = "1819266813802614784"

func ObtainClinicalTeacher(zgh string) (clinicalTeacher *ClinicalTeacher, bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		clinicalTeacher = nil
		return
	})

	// 创建请求体数据
	requestData := internal.RequestBody{
		ServiceID: ClinicalTeacherServiceID,
		Params: map[string]interface{}{
			"ZGH": zgh,
			//"page": 1,
			//"size": 20,
		},
	}
	// 发送请求
	response := internal.CommonRequest[ClinicalTeacher](requestData)
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

type ClinicalTeacher struct {
	ZGH   string `json:"ZGH"`   // 职工号
	XM    string `json:"XM"`    // 姓名
	XB    string `json:"XB"`    // 性别
	XYDM  string `json:"XYDM"`  // 学院代码
	XYMC  string `json:"XYMC"`  // 学院名称
	KSJYS string `json:"KSJYS"` // 科室/教研室
	RYFL  string `json:"RYFL"`  // 人员分类
	DQZT  string `json:"DQZT"`  // 当前状态
}
