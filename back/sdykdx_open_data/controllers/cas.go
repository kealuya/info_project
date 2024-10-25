package controllers

import (
	"encoding/xml"
	"fmt"
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/go-resty/resty/v2"
	"github.com/wumansgy/goEncrypt/des"
	"net/url"
)

type CasController struct {
	web.Controller
}

func (c *CasController) TestCas() {
	c.Redirect("https://baidu.com?niubi=232", 302)
}

func (c *CasController) PortalCas() {
	ticket := c.GetString("ticket")

	// 创建一个新的 resty 客户端
	client := resty.New()

	// 构建基础 URL    https://sso.ccmu.edu.cn/serviceValidate?service=http://gzpt.ccmu.edu.cn/personCommon/portal_cas&ticket=xxxxxxxx
	baseURL := "https://sso.ccmu.edu.cn/serviceValidate"

	//u, _ := url.Parse(baseURL)
	//// 构建查询参数
	//queryParams := url.Values{}
	// 读取casurl
	casUrl, _ := config.String("casurl")
	//queryParams.Set("service", casUrl)
	//queryParams.Set("ticket", ticket)
	//u.RawQuery = queryParams.Encode()
	//logs.Info("请求地址get::" + u.String())

	baseURL = baseURL + "?service=" + casUrl + "&ticket=" + ticket
	logs.Info("请求地址get::" + baseURL)
	// 发送 GET 请求
	resp, err := client.R().
		//SetQueryParamsFromValues(queryParams).
		Get(baseURL)

	if err != nil {
		msg := fmt.Sprintf("请求发送失败: %v\n", err)
		logs.Error(msg)
		c.CustomAbort(400, msg)
		return
	}

	// 检查响应状态码
	if resp.StatusCode() != 200 {
		msg := fmt.Sprintf("https://sso.ccmu.edu.cn/serviceValidate 服务器返回非200状态码: %d\n", resp.StatusCode())
		logs.Error(msg)
		c.CustomAbort(400, msg)
		return
	}

	// 解析 XML 响应
	var casResponse CASResponse
	err = xml.Unmarshal(resp.Body(), &casResponse)
	if err != nil {
		msg := fmt.Sprintf("XML 解析失败: %v\n", err)
		logs.Error(msg)
		c.CustomAbort(400, msg)
		return
	}
	// 创建一个 url.Values 对象
	params := url.Values{}
	dataMap := make(map[string]any)
	myResp := &Response{}
	// 处理响应
	if casResponse.Success != nil {
		success := casResponse.Success
		logs.Info(ticket+" - 验证成功::", fmt.Sprintf("用户名: %s\n", success.User), fmt.Sprintf("姓名: %s\n", success.Name), fmt.Sprintf("员工编号: %s\n", success.EmployeeNumber))
		params.Add("isOk", "true")
		params.Add("user", success.User)
		params.Add("name", success.Name)
		params.Add("employee_number", success.EmployeeNumber)
		dataMap["user"] = success.User
		dataMap["name"] = success.Name
		dataMap["employee_number"] = success.EmployeeNumber
		myResp.Success = true
		myResp.Data = dataMap
	} else {
		failure := casResponse.Failure
		reason := ""
		switch failure.Code {
		case InvalidRequest:
			reason = "请求参数错误"
		case InvalidTicket:
			reason = "ticket 验证失败"
		case InvalidService:
			reason = "ticket 验证成功，但是 service 不在白名单中"
		case InternalError:
			reason = "认证服务内部错误"
		default:
			reason = "未知错误类型"
		}
		logs.Info(ticket+" - 验证失败::", fmt.Sprintf("错误代码: %s\n", failure.Code), fmt.Sprintf("错误信息: %s\n", reason))
		params.Add("isOk", "false")
		params.Add("code", failure.Code)
		params.Add("msg", reason)
		myResp.Success = false
		myResp.Msg = reason
	}
	// 读取redirect
	redirect, _ := config.String("redirect")
	u, _ := url.Parse(redirect)
	// 设置查询参数
	desSecretKey := "szhtszht"
	hexText, err := des.DesCbcEncryptHex([]byte(params.Encode()), []byte(desSecretKey), nil)
	if err != nil {
		msg := fmt.Sprintf("参数加密失败: %v\n", err)
		logs.Error(msg)
		c.CustomAbort(400, msg)
		return
	}
	u.RawQuery = "rawQuery=" + hexText
	logs.Info("redirect_meta::" + redirect + "?rawQuery=" + params.Encode())
	logs.Info("redirect_dec::" + u.String())
	c.Redirect(u.String(), 302)
	//c.Data["json"] = myResp
	//_ = c.ServeJSON()
	//return
}

// CASResponse 结构体用于解析 CAS 服务的响应
type CASResponse struct {
	XMLName xml.Name               `xml:"serviceResponse"`
	Success *AuthenticationSuccess `xml:"authenticationSuccess"`
	Failure *AuthenticationFailure `xml:"authenticationFailure"`
}

type AuthenticationSuccess struct {
	User           string `xml:"user"`
	Name           string `xml:"name"`
	EmployeeNumber string `xml:"employeeNumber"`
}

type AuthenticationFailure struct {
	Code    string `xml:"code,attr"`
	Message string `xml:",chardata"`
}

// 错误码常量
const (
	InvalidRequest = "INVALID_REQUEST"
	InvalidTicket  = "INVALID_TICKET"
	InvalidService = "INVALID_SERVICE"
	InternalError  = "INTERNAL_ERROR"
)
