package models

import (
	"fmt"
	"sdykdx_open_data/common"
	"sdykdx_open_data/internal"
)

const GraduateStudentServiceID = "1819266356241797120"

func ObtainGraduateStudent(xh string) (graduateStudent *GraduateStudent, bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		graduateStudent = nil
		return
	})

	// 创建请求体数据
	requestData := internal.RequestBody{
		ServiceID: GraduateStudentServiceID,
		Params: map[string]interface{}{
			"XH": xh,
			//"page": 1,
			//"size": 20,
		},
	}
	// 发送请求
	response := internal.CommonRequest[GraduateStudent](requestData)
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

type GraduateStudent struct {
	XH   string `json:"XH"`   // 学号
	XM   string `json:"XM"`   // 姓名
	XB   string `json:"XB"`   // 性别
	NJ   string `json:"NJ"`   // 年级
	YXDM string `json:"YXDM"` // 院系代码
	YXMC string `json:"YXMC"` // 院系名称
	ZYH  string `json:"ZYH"`  // 专业号
	ZY   string `json:"ZY"`   // 专业名称
	BH   string `json:"BH"`   // 班号
	BJMC string `json:"BJMC"` // 班级名称
	PYCC string `json:"PYCC"` // 培养层次
	PYLB string `json:"PYLB"` // 培养类别
	RXFS string `json:"RXFS"` // 入校方式
	ZXZT string `json:"ZXZT"` // 在校状态
}
