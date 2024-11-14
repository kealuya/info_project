package models_qy_weixin

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
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
func convertToFormDataFormat(items []string) map[string]string {
	result := make(map[string]string)
	for i, item := range items {
		key := fmt.Sprintf("numbers[%d]", i)
		result[key] = item
	}
	return result
}
func SendQyWeixinNewsMessage(numbers []string, title string, description string, url string) (bizErr error) {

	defer common.RecoverHandler(func(err error) {
		bizErr = err
		logs.Error(err)
		return
	})

	// 发送请求
	numbersStrMap := convertToFormDataFormat(numbers)
	sendMessageResponse := SendMessageResponse{}

	m := map[string]string{
		"access_token":                            qy_weixin.GetQyWeixinAccessToken(),
		"content[msgtype]":                        "news",
		"content[news][articles][0][title]":       title,
		"content[news][articles][0][description]": description,
		"content[news][articles][0][url]":         url,
		"content[news][articles][0][picurl]":      "",
		"wid":                                     "22",
		"isall":                                   "0",
	}
	for key, value := range numbersStrMap {
		m[key] = value
	}

	logs.Info(fmt.Sprintf("SendQyWeixinNewsMessage -- requestUrl:%s , requestFormData: %#v", qy_weixin.MessageUrl, m))

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetFormData(m).SetResult(&sendMessageResponse).
		Post(qy_weixin.MessageUrl)

	logs.Info(fmt.Sprintf("SendQyWeixinNewsMessage -- response: %#v", resp.String()))
	common.ErrorHandler(err)
	if sendMessageResponse.E != 0 {
		common.ErrorHandler(fmt.Errorf("code:%d,%s", sendMessageResponse.E, sendMessageResponse.M))
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
	UCSID     int64    `json:"ucs_id"`     // 消息推送ID
	ErrorData []string `json:"error_data"` // 错误信息
}

// 定义请求结构体，完全匹配文档要求
type Article struct {
	Title       string `json:"title"`       // 必填，文章标题
	Description string `json:"description"` // 必填，文章描述
	URL         string `json:"url"`         // 必填，文章链接
	PicURL      string `json:"picurl"`      // 必填，封面图片链接
}

type News struct {
	Articles []Article `json:"articles"` // 必填，文章列表
}

type Content struct {
	MsgType string `json:"msgtype"` // 必填，固定为"news"
	News    News   `json:"news"`    // 必填
}

type PostData struct {
	AccessToken    string   `json:"access_token"`               // 必填，接口调用凭证
	Numbers        []string `json:"numbers"`                    // 必填，学工号列表
	IsAll          int      `json:"isall"`                      // 必填，是否全部推送
	Content        Content  `json:"content"`                    // 必填，消息体
	Wid            string   `json:"wid"`                        // 必填，企业微信应用ID
	Time           string   `json:"time,omitempty"`             // 可选，定时发送时间
	UcsType        *int     `json:"ucs_type,omitempty"`         // 可选，消息类别id
	IsPushIdentity *int     `json:"is_push_identity,omitempty"` // 可选，是否推送到用户的所有身份账号
}
