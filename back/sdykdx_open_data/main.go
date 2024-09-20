package main

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"sdykdx_open_data/controllers"
	"sdykdx_open_data/internal"
	"time"
)

func init() {
	LogConfigInit()
	TokenInit()
}
func TokenInit() {
	internal.ObtainToken()
	ticker := time.NewTicker(90 * time.Minute)
	go func() {
		for range ticker.C {
			internal.ObtainToken()
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
	web.Run()
}
