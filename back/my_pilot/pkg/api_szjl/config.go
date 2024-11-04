package api_szjl

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/viper"
	"log"
)

var (
	szjlConfig map[string]map[string]string
)

// GetSzjlConfig 获取配置
func GetSzjlConfig() map[string]map[string]string {
	return szjlConfig
}
func init() {
	viper.SetConfigName("api_szjl_config") // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")            // 如果配置文件的名称中没有扩展名，则需要配置此项

	// 添加配置文件路径
	viper.AddConfigPath("./pkg/api_szjl/") // 相对于项目根目录的路径
	viper.AddConfigPath("../api_szjl/")    // 相对于当前包的上一级目录
	viper.AddConfigPath(".")               // 当前目录

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("读取配置文件失败: %w", err)
	}

	// 将配置解析到结构体中
	// 创建新的配置实例
	szjlConfig = make(map[string]map[string]string)
	if err := viper.Unmarshal(&szjlConfig); err != nil {
		log.Panicf("解析配置文件失败: %w", err)
	}

	appKey = szjlConfig["szjl"]["app_key"]
	secretKey = szjlConfig["szjl"]["secret_key"]
	baseURL = szjlConfig["szjl"]["base_url"]
}

// 测试用账号及地址
var (
	appKey    string
	secretKey string
	baseURL   string
)

// Generate MD5 signature
func generateSign(secretKey, appKey, timestamp string) string {
	firstMd5 := md5.Sum([]byte(secretKey + appKey))
	firstHash := hex.EncodeToString(firstMd5[:])

	secondMd5 := md5.Sum([]byte(firstHash + timestamp))
	return hex.EncodeToString(secondMd5[:])
}

// Request  structures
type RequestHead struct {
	AppKey    string `json:"appKey"`
	Timestamp string `json:"timestamp"`
	Sign      string `json:"sign"`
	Version   string `json:"version"`
}
type Request[T any] struct {
	Head RequestHead `json:"head"`
	Data T           `json:"data"`
}

type Response[T any] struct {
	Code     int    `json:"code"`
	ErrorMsg string `json:"errorMsg"`
	Result   T      `json:"result"`
	RespId   string `json:"respId"`
}
