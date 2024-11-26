package api_jd_iop

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/go-resty/resty/v2"
	"github.com/gohouse/t"
	"github.com/jinzhu/copier"
	"github.com/spf13/viper"
	"log"
	"strings"
	"sync"
	"time"
)

var (
	jdIopConfig map[string]map[string]string

	myViper *viper.Viper
	rwMutex sync.RWMutex
)

func init() {
	myViper = viper.New()
	myViper.SetConfigName("api_jd_iop_config") // 配置文件名称(无扩展名)
	myViper.SetConfigType("yaml")              // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 添加配置文件路径
	myViper.AddConfigPath("./pkg/api_jd_iop/") // 相对于项目根目录的路径
	myViper.AddConfigPath("../api_jd_iop/")    // 相对于当前包的上一级目录
	myViper.AddConfigPath(".")                 // 当前目录

	logs.Info("init api_jd_iop_config.yaml")
	//读取配置文件
	if err := myViper.ReadInConfig(); err != nil {
		log.Panicf("读取配置文件失败: %w", err)
	}
	ReadJdIopConfig()
	// token处理逻辑
	TokenHandler()
	// 创建全局resty.Client
	CreateRestyClient()
}

// GetJdIopConfig GetSzjlConfig 获取配置
func GetJdIopConfig() map[string]map[string]string {
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	return jdIopConfig
}
func ReadJdIopConfig() {

	// 重新写入 jdIopConfig
	rwMutex.Lock()
	jdIopConfig = make(map[string]map[string]string)
	if err := myViper.Unmarshal(&jdIopConfig); err != nil {
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
	UID                 string `json:"uid"`
	AccessToken         string `json:"access_token"`  // 匹配文档中的 access_token
	RefreshToken        string `json:"refresh_token"` // 匹配文档中的 refresh_token
	Time                int64  `json:"time"`
	ExpiresIn           int    `json:"expires_in"`            // 匹配文档中的 expires_in
	RefreshTokenExpires int64  `json:"refresh_token_expires"` // 匹配文档中的 refresh_token_expires
}

// MD5加密函数
func getMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// getAccessToken 获取access token
func refreshAccessToken(clientID, clientSecret, refreshToken string) (*TokenResponse, error) {
	// 创建resty客户端
	client := resty.New()

	// 构造请求参数
	formData := map[string]string{
		"refresh_token": refreshToken,
		"client_id":     clientID,
		"client_secret": clientSecret,
	}

	// 发送请求并获取响应
	var result TokenResponse
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).SetResult(&result).
		ForceContentType("application/json"). // 京东侧返回的默认头不正确【Content-Type: text/plain;charset=UTF-8】，需要强制指定
		Post("https://api-iop.jd.com/oauth2/refreshToken")

	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}

	return &result, nil
}

// getAccessToken 获取access token
func getAccessToken(clientID, clientSecret, username, password string) (*TokenResponse, error) {
	// 创建resty客户端
	client := resty.New()

	// 获取当前时间戳
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	// 对密码进行md5加密（小写）
	passwordMD5 := getMD5(password)

	// 构造签名字符串
	signStr := fmt.Sprintf("%s%s%s%s%s%s%s",
		clientSecret,
		timestamp,
		clientID,
		username,
		passwordMD5,
		"access_token",
		clientSecret)

	// 生成签名（大写）
	sign := strings.ToUpper(getMD5(signStr))

	// 构造请求参数
	formData := map[string]string{
		"grant_type":    "access_token",
		"client_id":     clientID,
		"client_secret": clientSecret,
		"timestamp":     timestamp,
		"username":      username,
		"password":      passwordMD5,
		"sign":          sign,
	}

	// 发送请求并获取响应
	var result TokenResponse
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(formData).
		Post("https://api-iop.jd.com/oauth2/accessToken")

	if err != nil {
		return nil, fmt.Errorf("请求失败: %v", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("HTTP状态码错误: %d", resp.StatusCode())
	}

	// 直接解析响应体
	err = json.Unmarshal(resp.Body(), &result)
	if err != nil {
		return nil, fmt.Errorf("解析响应失败: %v", err)
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

	config := GetJdIopConfig()

	clientID := config["jd_iop"]["client_id"]
	clientSecret := config["jd_iop"]["client_secret"]
	resp, err := refreshAccessToken(clientID, clientSecret, config["token"]["refresh_token"])
	if err != nil {
		log.Panicf("refreshAccessToken失败: %v\n", err)
	}

	if !resp.Success {
		logs.Info("获取token失败: %s %s 。重新获取。\n", resp.ResultCode, resp.ResultMessage)
		username := config["jd_iop"]["username"]
		password := config["jd_iop"]["password"]
		// 获取refresh_token
		respGetAccessToken, errGetAccessToken := getAccessToken(clientID, clientSecret, username, password)
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
	myViper.Set("token.access_token", resp.Result.AccessToken)
	myViper.Set("token.access_token_time", resp.Result.Time)
	myViper.Set("token.access_token_time_format", time.UnixMilli(resp.Result.Time).Format("2006-01-02 15:04:05"))
	myViper.Set("token.refresh_token", resp.Result.RefreshToken)
	myViper.Set("token.refresh_token_expire", resp.Result.RefreshTokenExpires)

	ReadJdIopConfig()

	err = myViper.WriteConfig()
	if err != nil {
		log.Panicf("本地token写入失败: %v\n", err)
	}
	return nil
	//fmt.Printf("%+v\n", GetJdIopConfig())
}

func TokenHandlerNoSupport() {

	config := GetJdIopConfig()
	// accessToken过期时间需要加上一天
	expiredTime := t.New(config["token"]["access_token_time"]).Int64() + (24 * time.Hour).Milliseconds()

	if config["token"]["access_token"] == "" || IsTokenExpired(expiredTime) {
		// access_token 过期
		clientID := config["jd_iop"]["client_id"]
		clientSecret := config["jd_iop"]["client_secret"]
		username := config["jd_iop"]["username"]
		password := config["jd_iop"]["password"]
		var resp *TokenResponse
		var err error
		if config["token"]["refresh_token"] == "" || IsTokenExpired(t.New(config["token"]["refresh_token_expire"]).Int64()) {
			// refresh_token 过期
			resp, err = getAccessToken(clientID, clientSecret, username, password)
			if err != nil {
				log.Panicf("获取token失败: %v\n", err)
			}

			if !resp.Success {
				log.Panicf("获取token失败: %s\n", resp.ResultMessage)
			}
			logs.Info("refresh_token过期，获取正常:: %+v", resp)
		} else {
			// access_token 过期
			resp, err = refreshAccessToken(clientID, clientSecret, config["token"]["refresh_token"])
			if err != nil {
				log.Panicf("refreshAccessToken失败: %v\n", err)
			}

			if !resp.Success {
				log.Panicf("获取token失败: %s\n", resp.ResultMessage)
			}
			logs.Info("access_token过期，获取正常:: %+v", resp)
		}

		// 更新内存中的配置
		myViper.Set("token.access_token", resp.Result.AccessToken)
		myViper.Set("token.access_token_time", resp.Result.Time)
		myViper.Set("token.refresh_token", resp.Result.RefreshToken)
		myViper.Set("token.refresh_token_expire", resp.Result.RefreshTokenExpires)

		err = myViper.WriteConfig()
		if err != nil {
			log.Panicf("本地token写入失败: %v\n", err)
		}
		ReadJdIopConfig()
	} else {
		logs.Info("access_token正常可用")
	}

	fmt.Printf("%+v\n", GetJdIopConfig())

}
