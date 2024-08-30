package internal

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"
	"log"
	"sync"
)

const (
	baseUrl   = "https://data.ccmu.edu.cn/open-data/"
	TokenUrl  = baseUrl + "getAccessToken"
	ApiUrl    = baseUrl + "api"
	appId     = "20a6f1bdfa4c444da1502843ce743e1b"
	appSecret = "68748af104534cefb319df5ecce54d39"
)

type ResponseToken struct {
	Success   bool   `json:"success"`
	Code      int    `json:"code"`
	Data      string `json:"data"`
	Timestamp string `json:"timestamp"`
}

var token = ""
var rwMutex sync.RWMutex

func GetToken() string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return token
}

func ObtainToken() string {
	rwMutex.Lock()
	client := resty.New()
	// 发起 GET 请求以获取accessToken
	tokenStruct := ResponseToken{}
	resp, err := client.R().
		SetQueryParam("appId", appId).
		SetQueryParam("appSecret", appSecret).
		SetResult(&tokenStruct).
		Get(TokenUrl)
	if err != nil {
		logs.Error(err)
		log.Panic(err)
	}
	if resp.IsError() {
		logs.Error(resp.String())
		log.Panic(resp.String())
	}

	if !tokenStruct.Success {
		logs.Error(fmt.Sprintf("%#v", tokenStruct))
		log.Panic(resp.String())
	}
	logs.Info("获取token::", fmt.Sprintf("%+v", tokenStruct))

	token = tokenStruct.Data
	rwMutex.Unlock()
	return tokenStruct.Data

}
