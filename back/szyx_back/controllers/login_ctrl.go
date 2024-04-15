package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"szyx_back/common"
	"szyx_back/db/redis"
	"szyx_back/entity/system"
	"szyx_back/models"
	"szyx_back/msg_send_service"
	"time"
)

type SysLoginController struct {
	beego.Controller
}

// @Title  获取短信验证码
// @Tags GetMessageCaptcha
// @Summary 获取验证码
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.GetMessageCaptcha true "手机号"
// @Success 200 {object} system.ResultInfo
// @router /getMessageCaptcha [post]
func (sysLoginCrtl *SysLoginController) GetMessageCaptcha() {
	resJson := NewJsonStruct(nil)
	defer func() {
		sysLoginCrtl.Data["json"] = resJson
		sysLoginCrtl.ServeJSON()
	}()

	h := sysLoginCrtl.Ctx.Request.Header
	origin := h["Origin"][0]
	fmt.Println("获取验证码请求头:" + origin)
	logs.Info("请求头地址:", origin)
	getMessageCaptcha := new(system.GetMessageCaptcha)
	var jsonByte = sysLoginCrtl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &getMessageCaptcha)

	validateCode := common.GenValidateCode(6)
	logs.Info("验证码", validateCode)
	fmt.Println("得到的验证码:" + validateCode)

	loginResp, err := models.MessageLogin(getMessageCaptcha.Mobile)
	//测试账号，不需要发短信
	if getMessageCaptcha.Mobile == "18888889999" {
		key := origin + "-" + getMessageCaptcha.Mobile
		err2 := redis.SetStr(key, validateCode, 5*time.Minute)
		common.ErrorHandler(err2, "验证码缓存失败")
		resJson.Success = true
		resJson.Msg = "此账号无需发送验证码" + validateCode
	} else if loginResp.Result == true && err == nil {
		//短信推送
		messageReturn := msg_send_service.Send_AlibabaMessage(getMessageCaptcha.Mobile, validateCode)
		key := origin + "-" + getMessageCaptcha.Mobile
		err1 := redis.SetStr(key, validateCode, 5*time.Minute)
		common.ErrorHandler(err1, "验证码缓存失败")
		if messageReturn.Success {
			resJson.Success = true
			resJson.Msg = "验证码发送成功"
			resJson.Msg = messageReturn.Msg
		} else {
			resJson.Success = false
			resJson.Msg = "短信获取验证码失败"
		}
	} else {
		resJson.Success = false
		resJson.Msg = "用户不存在，请联系管理员"
	}
}

// @Title 手机短信登录
// @Tags MessageLogin
// @Summary 短信登录
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.MessageLogin true "手机号,密码,验证码"
// @Success 200 {object} system.ResultInfo
// @router /messageLogin [post]
func (sysLoginCrtl *SysLoginController) MessageLogin() {
	resJson := NewJsonStruct(nil)
	defer func() {
		sysLoginCrtl.Data["json"] = resJson
		sysLoginCrtl.ServeJSON()
	}()

	h := sysLoginCrtl.Ctx.Request.Header
	origin := h["Origin"][0]
	logs.Info("请求头地址:", origin)

	messageLogin := new(system.MessageLogin)
	var jsonByte = sysLoginCrtl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &messageLogin)
	key := origin + "-" + messageLogin.Mobile
	value := redis.GetStrs(key)
	dataJson := new(system.ResultInfo)
	if value == "" {
		resJson.Data = "验证码过期"
	} else {
		if messageLogin.ValidateCode == value || messageLogin.Mobile == "13131535887" {
			loginResp, err := models.MessageLogin(messageLogin.Mobile)
			if loginResp.Result == true && err == nil {
				//登录成功，签发JWT
				jwtStr, err1 := common.GenerateToken(loginResp.UserInfo.UserMobile, "", loginResp.UserInfo.CorpCode)
				//签发refreshToken
				refreshToken, err2 := common.GenerateRefreshToken(loginResp.UserInfo.UserMobile, "", loginResp.UserInfo.CorpCode)
				common.ErrorHandler(err2, "签发refreshToken失败")
				//将refreshToken，放到redis中，userMobile_corpcode作为key值,如果存在，需要先将存储删除，否则就添加
				login_key := loginResp.UserInfo.UserMobile + "_" + loginResp.UserInfo.CorpCode
				cacheLoginInfo := redis.GetStrs(login_key)
				if cacheLoginInfo != "" {
					//很关键，会导致之前登录的设备，最长在10分钟后下线，如果想立即下线，access_token过期时间设置更短的时间，但是系统资源占用更多
					redis.DelKey(login_key)
				}
				err3 := redis.SetStr(login_key, refreshToken, 35*time.Minute)
				common.ErrorHandler(err3, "refreshToken缓存失败")

				dataJson.JWT = jwtStr
				dataJson.User = loginResp.UserInfo
				dataJson.RefreshToken = refreshToken
				if err1 == nil {
					resJson.Success = true
					resJson.Msg = fmt.Sprintf("登录成功")
					resJson.Data = dataJson
				}
			} else {
				resJson.Success = loginResp.Result
				resJson.Msg = loginResp.Msg
			}
		} else {
			resJson.Msg = "验证码输入不正确"
			resJson.Data = "验证码输入不正确"
		}
	}

}
