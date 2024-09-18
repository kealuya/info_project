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
	web.CtrlGet("/organization", (*controllers.BusinessController).Organization)
	web.CtrlGet("/clinical_teacher", (*controllers.BusinessController).ClinicalTeacher)
	web.CtrlGet("/faculty", (*controllers.BusinessController).Faculty)
	web.CtrlGet("/graduate_student", (*controllers.BusinessController).GraduateStudent)
	web.Run()
}
