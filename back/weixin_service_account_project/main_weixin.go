package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"weixin_service_account_project/module"
	_ "weixin_service_account_project/util"
)

func init() {
	_ = logs.SetLogger(logs.AdapterConsole)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
}

func main() {
	// 用于注册微信返回地址
	beego.Router("/wechat", &module.ManagerController{}, "get:VerifyToken")
	// 用于接受用户提问，通过大模型回答
	beego.Router("/wechat", &module.WeChatController{}, "post:HandleMessage")

	// Start beego server
	beego.Run()
}
