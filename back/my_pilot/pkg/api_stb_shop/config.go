package api_stb_shop

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var (
	stbConfig map[string]map[string]string

	myViper *viper.Viper
	rwMutex sync.RWMutex
)

func init() {
	myViper = viper.New()
	myViper.SetConfigName("api_stb_shop_config") // 配置文件名称(无扩展名)
	myViper.SetConfigType("yaml")                // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 添加配置文件路径
	myViper.AddConfigPath("./pkg/api_stb_shop/") // 相对于项目根目录的路径
	myViper.AddConfigPath("../api_stb_shop/")    // 相对于当前包的上一级目录
	myViper.AddConfigPath(".")                   // 当前目录

	logs.Info("init api_stb_shop_config.yaml")
	//读取配置文件
	if err := myViper.ReadInConfig(); err != nil {
		log.Panicf("读取配置文件失败: %w", err)
	}
	ReadStbConfig()
	// token处理逻辑
	TokenHandler()
	// 创建全局resty.Client
	CreateRestyClient()
}

// GetStbConfig   获取配置
func GetStbConfig() map[string]map[string]string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return stbConfig
}
func ReadStbConfig() {

	// 重新写入 jdIopConfig
	rwMutex.Lock()
	stbConfig = make(map[string]map[string]string)
	if err := myViper.Unmarshal(&stbConfig); err != nil {
		rwMutex.Unlock()
		log.Panicf("解析配置文件失败: %w", err)
	}
	rwMutex.Unlock()
}

// TokenResponse 响应结构体
type TokenResponse struct {
	Response[[]TokenResult]
}

type TokenResult struct {
	UID                 string `json:"uid"`
	AccessToken         string `json:"access_token"`  // 匹配文档中的 access_token
	RefreshToken        string `json:"refresh_token"` // 匹配文档中的 refresh_token
	Time                string `json:"time"`
	ExpiresIn           string `json:"expires_in"`            // 匹配文档中的 expires_in
	RefreshTokenExpires string `json:"refresh_token_expires"` // 匹配文档中的 refresh_token_expires
}

// refreshAccessToken 获取access token
func refreshAccessToken(clientID, clientSecret, refreshToken string) (*TokenResponse, error) {
	// 创建resty客户端
	client := resty.New()

	// 构造请求参数
	formData := map[string]string{
		"refresh_token": refreshToken,
		"client_id":     clientID,
		"client_secret": clientSecret,
	}

	config := GetStbConfig()
	baseUrl := config["stb_shop"]["base_url"] + "refresh_token"

	// 发送请求并获取响应
	var result TokenResponse
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(formData).
		SetResult(&result).
		Post(baseUrl)

	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}

	return &result, nil
}

// getAccessToken 获取access token
func getAccessToken(clientID, clientSecret, code, redirectUrl string) (*TokenResponse, error) {
	// 创建resty客户端
	client := resty.New()

	// 构造请求参数
	formData := map[string]string{
		"grant_type":    "authorization_code",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"redirect_uri":  redirectUrl,
		"code":          code,
	}

	config := GetStbConfig()

	baseUrl := config["stb_shop"]["base_url"] + "access_token"

	// 发送请求并获取响应
	var result TokenResponse
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParams(formData).
		SetResult(&result).
		Post(baseUrl)

	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}

	return &result, nil
}
func TokenHandler() {
	// 启动执行，然后每10小时执行一次
	err := runToken()
	if err != nil {
		log.Panicf("panic runToken: %v", err)
	}
	ticker := time.NewTicker(time.Hour * 10)
	go func() {
		for _ = range ticker.C {
			errN := runToken()
			if errN != nil {
				logs.Error("panic runToken: %v", errN)
			} else {
				logs.Info("执行 AccessToken 刷新正常 ")
			}
		}
	}()

}

// IsTokenExpired 判断access_token是否过期
func IsTokenExpired(expire int64) bool {
	return time.Now().UnixMilli() >= expire
}

func runToken() (bizErr error) {

	defer func() {
		if r := recover(); r != nil {
			logs.Error("panic runToken: %v", r)
			bizErr = fmt.Errorf("panic runToken: %v", r)
		}
	}()

	config := GetStbConfig()

	clientID := config["stb_shop"]["client_id"]
	clientSecret := config["stb_shop"]["client_secret"]
	resp, err := refreshAccessToken(clientID, clientSecret, config["token"]["refresh_token"])
	if err != nil {
		log.Panicf("refreshAccessToken失败: %v\n", err)
	}

	if !resp.Success {
		logs.Info("获取token失败: %s %s 。重新获取。\n", resp.ResultCode, resp.ResultMessage)
		code := config["stb_shop"]["code"]
		redirectUri := config["stb_shop"]["redirect_uri"]
		// 获取refresh_token
		respGetAccessToken, errGetAccessToken := getAccessToken(clientID, clientSecret, code, redirectUri)
		if errGetAccessToken != nil {
			log.Panicf("获取token失败: %v\n", errGetAccessToken)
		}
		if !respGetAccessToken.Success {
			log.Panicf("获取token失败: %s\n", respGetAccessToken.ResultMessage)
		}
		logs.Info("执行 accessToken 正常:: %+v", respGetAccessToken)

		errCopier := copier.Copy(&resp, respGetAccessToken)
		if errCopier != nil {
			log.Panicf("copier.Copy失败: %s\n", errCopier)
		}
	} else {
		logs.Info("执行 refreshAccessToken 正常:: %+v", resp)
	}

	// 更新内存中的配置
	myViper.Set("token.access_token", resp.Result[0].AccessToken)
	myViper.Set("token.access_token_time", resp.Result[0].Time)
	myViper.Set("token.refresh_token", resp.Result[0].RefreshToken)
	myViper.Set("token.refresh_token_expire", resp.Result[0].RefreshTokenExpires)

	ReadStbConfig()

	err = myViper.WriteConfig()
	if err != nil {
		log.Panicf("本地token写入失败: %v\n", err)
	}
	return nil
}
