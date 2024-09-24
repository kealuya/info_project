package qy_weixin

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"
	"log"
	"sync"
)

const (
	baseUrl         = "移动校园域名"
	TokenUrl        = baseUrl + "/api/third/get-token"
	RefreshTokenUrl = baseUrl + "/api/third/refresh-token"
	UserInfoUrl     = baseUrl + "/uc/api/oauth/user-by-code"
	MessageUrl      = baseUrl + "/uc/api/ucs/qywx"
	AppId           = "xxx"
	AppSecret       = "xxxx"
)

var token = ""
var rwMutex sync.RWMutex

func GetQyWeixinAccessToken() string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return token
}

func ObtainQyWeixinAccessToken(url string) string {
	rwMutex.Lock()
	client := resty.New()
	// 发起 GET 请求以获取accessToken
	tokenStruct := ResponseToken{}
	resp, err := client.R().
		SetQueryParam("appid", AppId).
		SetQueryParam("appsecret", AppSecret).
		SetHeader("content-type", "application/x-www-form-urlencoded").
		SetResult(&tokenStruct).
		Get(url)
	if err != nil {
		logs.Error(err)
		log.Panic(err)
	}
	if resp.IsError() {
		logs.Error(resp.String())
		log.Panic(resp.String())
	}

	if tokenStruct.E != 0 {
		logs.Error(fmt.Sprintf("%#v", tokenStruct))
		log.Panic(resp.String())
	}
	logs.Info("获取QyWeixinAccessToken::", fmt.Sprintf("%+v", tokenStruct))

	token = tokenStruct.D.AccessToken
	rwMutex.Unlock()
	return tokenStruct.D.AccessToken

}

type ResponseToken struct {
	E int      `json:"e"` // 返回码，成功返回0，其他为异常
	M string   `json:"m"` // 返回描述，用于说明成功或错误信息
	D struct { // 返回内容，包含实际的数据
		AccessToken string `json:"access_token"` // 接口调用凭证，需要保存以进行后续请求
		ExpiresIn   int    `json:"expires_in"`   // 凭证的有效期，以秒为单位
	} `json:"d"`
}
