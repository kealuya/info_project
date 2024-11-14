package models

import (
	"fmt"
	"sdykdx_open_data/common"
	"sdykdx_open_data/internal"
)

const FacultyServiceID = "1810234648378056704"

func ObtainFaculty(zgh, xm string) (faculty *[]Faculty, bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		faculty = nil
		return
	})

	// 创建请求体数据
	requestData := internal.RequestBody{
		ServiceID: FacultyServiceID,
		Params: map[string]interface{}{
			"ZGH": zgh,
			"XM":  xm,
			//"page": 1,
			//"size": 20,
		},
	}
	// 发送请求
	response := internal.CommonRequest[Faculty](requestData)
	if !response.Success {
		return nil, fmt.Errorf("api接口调用失败")
	}
	if len(response.Data.Data) == 0 {
		return nil, fmt.Errorf("查询无结果")
	}

	return &response.Data.Data, nil
}

type Faculty struct {
	ZGH    string `json:"ZGH"`    // 职工号
	XM     string `json:"XM"`     // 姓名
	XBDM   string `json:"XBDM"`   // 性别代码
	SZDWBM string `json:"SZDWBM"` // 所在单位代码
	SZDWDW string `json:"SZDWDW"` // 所在单位
	XXBM   string `json:"XXBM"`   // 学系/办公室码
	XXMC   string `json:"XXMC"`   // 学系/办公室
	RYFL   string `json:"RYFL"`   // 人员分类
	RYZTMC string `json:"RYZTMC"` // 当前状态
}
