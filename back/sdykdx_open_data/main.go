package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"sdykdx_open_data/controllers"
	"sdykdx_open_data/internal"
	"sdykdx_open_data/internal/qy_weixin"
	"time"
)

func init() {
	logs.Info("init service")
	LogConfigInit()
	TokenInit()
	QyWeixinAccessTokenInit()
}
func TokenInit() {
	internal.ObtainToken()
	ticker := time.NewTicker(100 * time.Minute)
	go func() {
		for range ticker.C {
			internal.ObtainToken()
		}
	}()
}
func QyWeixinAccessTokenInit() {
	qy_weixin.ObtainQyWeixinAccessToken(qy_weixin.TokenUrl)
	ticker := time.NewTicker(110 * time.Minute)
	go func() {
		for range ticker.C {
			qy_weixin.ObtainQyWeixinAccessToken(qy_weixin.RefreshTokenUrl)
		}
	}()
}
func LogConfigInit() {
	_ = logs.SetLogger(logs.AdapterConsole)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/my.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365,"color":true,"separate":["error", "warning", "info", "debug"]}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出log
	//logs.Async()
}
func main() {
	BeeGoRun()
}

func BeeGoRun() {
	logs.Info("start service")
	//web.BConfig.CopyRequestBody = true
	//web.BConfig.Listen.HTTPPort = 8080
	// Organization 组织架构信息?dwdm=2222222
	web.CtrlGet("/organization", (*controllers.BusinessController).Organization)
	// ClinicalTeacher 临床教师信息?zgh=2222222
	web.CtrlGet("/clinical_teacher", (*controllers.BusinessController).ClinicalTeacher)
	// Faculty 教职工信息?zgh=2222222
	web.CtrlGet("/faculty", (*controllers.BusinessController).Faculty)
	// GraduateStudent 研究生基本信息?xh=2222222
	web.CtrlGet("/graduate_student", (*controllers.BusinessController).GraduateStudent)

	/*
		企业微信登录，页面部分判断用户是否已经登录（storage判断），未登录的场合，跳转企业微信地址，并填写redirect再回调到当前地址
		移动校园域名/uc/api/oauth/index?redirect=回调地址&appid=200240918103235340&state=XXXX&qrcode=1
		appid可以在页面填好，但更建议从后台请求，获取后拼接调用
		state是自己识别用的字段
		用户授权成功后将跳转至`redirect?state=XXXX&code=CODE
		页面部分，拿到code，并回传后台
	*/
	// 返回：{"success":true,"msg":"","data":{"appId":"222222222"}}
	web.CtrlGet("/qy_weixin_appid", (*controllers.QyWeixinController).GetAppId)
	/*
			请求：{"code":"用户code"}
			返回：
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
	*/
	web.CtrlPost("/qy_weixin_userinfo", (*controllers.QyWeixinController).GetUserByCode)
	/*
		请求：{"numbers":"学号","description":"消息描述","title":"消息标题","url":"跳转链接"}
		返回：{"success":true,"msg":"","data":null}
	*/
	web.CtrlPost("/qy_weixin_send_message", (*controllers.QyWeixinController).SendQyWeixinMessage)
	web.CtrlGet("/portal_cas", (*controllers.CasController).PortalCas)
	web.CtrlGet("/test_cas", (*controllers.CasController).TestCas)
	web.Run()
}
