package main

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"runtime"
	"weixin_service_account_project/module"
)

func init() {
	_ = logs.SetLogger(logs.AdapterConsole)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	config()
}

func config() {
	// 获取当前系统的 CPU 核心数
	numCPU := runtime.NumCPU()
	logs.Info("Number of CPUs: %d\n", numCPU)

	// 设置 GOMAXPROCS 为 CPU 核心数
	prevGOMAXPROCS := runtime.GOMAXPROCS(numCPU)
	logs.Info("Previous GOMAXPROCS setting: %d\n", prevGOMAXPROCS)
	logs.Info("Current GOMAXPROCS setting: %d\n", runtime.GOMAXPROCS(numCPU))
}

func main() {
	// 用于注册微信返回地址
	beego.Router("/wechat", &module.ManagerController{}, "get:VerifyToken")
	// 用于接受用户提问，通过大模型回答
	beego.Router("/wechat", &module.WeChatController{}, "post:HandleMessage")
	beego.Router("/", &module.QyWeixinController{}, "get:Test")

	// Start beego server
	beego.Run()
}
