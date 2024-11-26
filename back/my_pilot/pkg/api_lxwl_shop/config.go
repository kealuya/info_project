package api_lxwl_shop

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	lxwlConfig map[string]map[string]string

	rwMutex sync.RWMutex
	myViper *viper.Viper
)

func init() {
	myViper = viper.New()
	myViper.SetConfigName("api_lxwl_shop_config") // 配置文件名称(无扩展名)
	myViper.SetConfigType("yaml")                 // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 添加配置文件路径
	myViper.AddConfigPath("./pkg/api_lxwl_shop/") // 相对于项目根目录的路径
	myViper.AddConfigPath("../api_lxwl_shop/")    // 相对于当前包的上一级目录
	myViper.AddConfigPath(".")                    // 当前目录

	logs.Info("init api_lxwl_shop_config.yaml")
	//读取配置文件
	if err := myViper.ReadInConfig(); err != nil {
		log.Panicf("读取配置文件失败: %w", err)
	}
	ReadLxwlConfig()
	// token处理逻辑
	TokenHandler()
	// 创建全局resty.Client
	CreateRestyClient()
}

// GetLxwlConfig GetSzjlConfig 获取配置
func GetLxwlConfig() map[string]map[string]string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return lxwlConfig
}
func ReadLxwlConfig() {

	// 重新写入 lxwlConfig
	rwMutex.Lock()
	lxwlConfig = make(map[string]map[string]string)
	if err := myViper.Unmarshal(&lxwlConfig); err != nil {
		rwMutex.Unlock()
		log.Panicf("解析配置文件失败: %w", err)
	}
	rwMutex.Unlock()
}

// TokenResponse 响应结构体
type TokenResponse struct {
	Response[TokenResult]
}

type TokenResult struct {
	AccessToken string `json:"accessToken"` // 匹配文档中的 accessToken
	ExpiresIn   string `json:"expiresAt"`   // 匹配文档中的 expiresAt
}

// MD5加密函数
func getMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// getAccessToken 获取access token
func getAccessToken(username, password string) (*TokenResponse, error) {

	config := GetLxwlConfig()

	baseUrl := config["lxwl_shop"]["base_url"] + "auth2/accessToken"

	// 创建resty客户端
	client := resty.New()
	// 获取当前时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	// 构造签名字符串
	signStr := fmt.Sprintf("%s%s%s%s",
		username,
		password,
		timestamp,
		password,
	)

	// 生成签名（大写）
	sign := strings.ToLower(getMD5(signStr))

	// 构造请求参数
	formData := map[string]string{
		"timestamp": timestamp,
		"username":  username,
		"password":  password,
		"sign":      sign,
	}

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

func runToken() (bizErr error) {

	defer func() {
		if r := recover(); r != nil {
			logs.Error("panic runToken: %v", r)
			bizErr = fmt.Errorf("panic runToken: %v", r)
		}
	}()

	config := GetLxwlConfig()

	username := config["lxwl_shop"]["username"]
	password := config["lxwl_shop"]["password"]
	// 获取 token
	respGetAccessToken, errGetAccessToken := getAccessToken(username, password)
	if errGetAccessToken != nil {
		log.Panicf("获取token失败: %v\n", errGetAccessToken)
	}
	if !respGetAccessToken.Success {
		log.Panicf("获取token失败: %s\n", respGetAccessToken.Message)
	}
	logs.Info("执行 accessToken 正常:: %+v", respGetAccessToken)

	// 更新内存中的配置
	myViper.Set("token.access_token", respGetAccessToken.Result.AccessToken)
	myViper.Set("token.access_token_time_format", respGetAccessToken.Result.ExpiresIn)

	ReadLxwlConfig()

	err := myViper.WriteConfig()
	if err != nil {
		log.Panicf("本地token写入失败: %v\n", err)
	}
	return nil
}
