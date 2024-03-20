package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/mojocn/base64Captcha"
	"szyx_back/common"
	. "szyx_back/entity"
	"szyx_back/models"
)

type SysLoginController struct {
	beego.Controller
}

var store = base64Captcha.DefaultMemStore

// @Title 获取验证码
// @Tags CreateCaptcha
// @Summary 获取验证码
// @Security token
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @router /createCaptcha [post]
func (sysLoginCrtl *SysLoginController) CreateCaptcha() {
	resJson := NewJsonStruct(nil)
	defer func() {
		sysLoginCrtl.Data["json"] = string(common.Marshal(resJson))
		sysLoginCrtl.ServeJSON()
	}()

	// 字符,公式,验证码配置
	// 生成默认数字的driver
	//driver := base64Captcha.NewDriverDigit(20, 50, 4, 0.7, 80)
	driver := base64Captcha.DefaultDriverDigit
	//cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))   // 使用redis
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("验证码获取失败")
	} else {
		captcha := new(Captcha)
		captcha.CaptchaId = id
		captcha.PicPath = b64s
		resJson.Success = true
		resJson.Msg = fmt.Sprintf("验证码获取成功")
		resJson.Data = captcha
	}
}

// @Title 用户登录
// @Tags Login
// @Summary 用户登录
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.LoginRequestPar true "用户名,密码,验证码"
// @Success 200 {object} system.ResultInfo
// @router /login [post]
func (sysLoginCrtl *SysLoginController) Login() {
	resJson := NewJsonStruct(nil)
	defer func() {
		sysLoginCrtl.Data["json"] = string(common.Marshal(resJson))
		sysLoginCrtl.ServeJSON()
	}()

	loginRequestPar := new(LoginRequestPar)
	var jsonByte = sysLoginCrtl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &loginRequestPar)

	//验证码对比
	if store.Verify(loginRequestPar.CaptchaId, loginRequestPar.Captcha, true) {

		user, err := models.UserLogin(loginRequestPar)
		if err == nil {
			resJson.Success = true
			resJson.Data = user
			resJson.Msg = fmt.Sprintf("登录成功")
		} else {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("用户名或密码错误")
		}
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("验证码错误！")
	}
}

// @Title 修改用户密码
// @Tags SetUserInfo
// @Summary 修改用户密码
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.SysUser true "用户Id,用户密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"设置成功"}"
// @router /setUserInfo [post]
func (userCtrl *SysLoginController) SetUserInfo() {
	resJson := NewJsonStruct(nil)
	defer func() {
		userCtrl.Data["json"] = string(common.Marshal(resJson))
		userCtrl.ServeJSON()
	}()

	setUserInfo := new(UserInfo)
	var jsonByte = userCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &setUserInfo)
	user, err := models.SetUserInfo(setUserInfo)
	if err == nil {
		resJson.Success = true
		resJson.Data = user
		resJson.Data = "设置成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("设置用户信息失败")
	}
}
