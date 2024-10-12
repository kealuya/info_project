package models_qy_weixin

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"sdykdx_open_data/common"
	"sdykdx_open_data/internal/qy_weixin"
	"strings"
)

func ObtainUserInfoByCode(code string) (userInfo *UserInfo, bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		userInfo = nil
		return
	})

	// 发送请求

	qyWeixinUserInfo := QyWeixinUserInfo{}
	client := resty.New()
	_, err := client.R().
		SetQueryParam("code", code).
		SetQueryParam("access_token", qy_weixin.GetQyWeixinAccessToken()).
		SetHeader("content-type", "application/x-www-form-urlencoded").
		SetResult(&qyWeixinUserInfo).
		Get(qy_weixin.UserInfoUrl)
	common.ErrorHandler(err)
	if qyWeixinUserInfo.E != 0 {
		return nil, fmt.Errorf(qyWeixinUserInfo.M)
	}

	return &qyWeixinUserInfo.D, nil
}

type QyWeixinUserInfo struct {
	E int      `json:"e"` // 通常用于表示错误码或状态码
	M string   `json:"m"` // 通常用于表示错误信息或状态信息
	D UserInfo `json:"d"`
}
type UserInfo struct { // 用户详细信息
	Realname string `json:"realname"` // 用户真实姓名
	Sex      string `json:"sex"`      // 用户性别
	Email    string `json:"email"`    // 邮箱地址
	Mobile   string `json:"mobile"`   // 手机号
	Weixin   string `json:"weixin"`   // 微信号
	Qq       string `json:"qq"`       // QQ 号
	Csrq     string `json:"csrq"`     // 初始日期，格式为 YYYYMMDD
	Mz       string `json:"mz"`       // 民族，使用国标代码表示
	Avatar   string `json:"avatar"`   // 用户头像 URL

	Role struct { // 身份信息
		Number   string `json:"number"`   // 学工号
		Identity string `json:"identity"` // 身份描述
	} `json:"role"`
	Departs json.RawMessage `json:"departs"` // 所在学院或部门名称
	// 部门字段，可能需要根据实际情况调整为合适的表示方式
}

// Helper function to convert []string to "['str1','str2']" format
func convertToJSONArrayString(items []string) string {
	quotedItems := make([]string, len(items))
	for i, item := range items {
		quotedItems[i] = fmt.Sprintf("'%s'", item)
	}
	return fmt.Sprintf("[%s]", strings.Join(quotedItems, ","))
}
func SendQyWeixinNewsMessage(numbers []string, title string, description string, url string) (bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		return
	})

	// 发送请求
	numbersStr := convertToJSONArrayString(numbers)
	sendMessageResponse := SendMessageResponse{}
	client := resty.New()
	_, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(map[string]string{
			"access_token":                            qy_weixin.GetQyWeixinAccessToken(),
			"numbers[]":                               numbersStr,
			"content[msgtype]":                        "news",
			"content[news][articles][0][title]":       title,
			"content[news][articles][0][description]": description,
			"content[news][articles][0][url]":         url,
			"content[news][articles][0][picurl]":      "",
			"wid":                                     "22",
		}).SetResult(&sendMessageResponse).
		Post(qy_weixin.MessageUrl)
	common.ErrorHandler(err)
	if sendMessageResponse.E != 0 {
		return fmt.Errorf(sendMessageResponse.M)
	}

	return nil
}

type SendMessageResponse struct {
	E int    `json:"e"` // 返回码
	M string `json:"m"` // 返回描述
	D Data   `json:"d"` // 返回内容
}

// Data encapsulates the detailed data in the response
type Data struct {
	UCSID     string   `json:"ucs_id"`     // 消息推送ID
	ErrorData []string `json:"error_data"` // 错误信息
}
