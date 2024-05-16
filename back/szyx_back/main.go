package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/astaxie/beego/toolbox"
	_ "szyx_back/configs"
	_ "szyx_back/db/handler"
	_ "szyx_back/models"
	_ "szyx_back/routers"
	"szyx_back/schedule"
)

func main() {

	//beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	//定时任务
	schedule.InitTask()
	toolbox.StartTask()
	defer toolbox.StopTask()

	beego.Run()

	//kdxf.Test()
}
