package api_lxwl_shop

import (
	"github.com/bytedance/sonic"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

var (
	GlobalClient *resty.Client
	JsonEncoder  sonic.API
)

func CreateRestyClient() {

	// 创建全局 resty 客户端，避免重复创建
	GlobalClient = resty.New().
		SetTimeout(20 * time.Second). // 设置超时
		SetRetryCount(2).             // 设置重试次数
		SetRetryWaitTime(1000 * time.Millisecond).
		SetRetryMaxWaitTime(5000 * time.Millisecond).
		SetTransport(&http.Transport{
			MaxIdleConnsPerHost:   100,              // 每个host最大空闲连接数
			MaxConnsPerHost:       200,              // 每个host最大连接数
			IdleConnTimeout:       90 * time.Second, // 空闲连接超时时间
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			DisableCompression:    false, // 如果不需要压缩可以禁用
		})
	// 初始化 sonic 编码器
	JsonEncoder = sonic.ConfigDefault

}

// Response   响应结构体
type Response[T any] struct {
	Success bool   `json:"success"`
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}
