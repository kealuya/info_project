package internal

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"
	"sdykdx_open_data/common"
)

// RequestBody 请求和返回结构体定义
type RequestBody struct {
	ServiceID string                 `json:"serviceId"`
	Params    map[string]interface{} `json:"params"`
}
type Response[T any] struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	Data    struct {
		Page       int `json:"page"`
		Size       int `json:"size"`
		TotalRows  int `json:"totalRows"`
		TotalPages int `json:"totalPages"`
		Data       []T `json:"data"`
	} `json:"data"`
	Timestamp string `json:"timestamp"`
}

func CommonRequest[T any](requestData RequestBody) Response[T] {
	response := Response[T]{}
	// 创建 Resty 客户端
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetHeader("accessToken", GetToken()).
		SetBody(requestData).
		SetResult(&response).
		Post(ApiUrl)

	common.ErrorHandler(err)

	logs.Info(fmt.Sprintf("request::%+v -- response::%+v", requestData, response))
	return response
}
