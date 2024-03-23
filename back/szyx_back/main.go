package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/toolbox"
	_ "szyx_back/configs"
	_ "szyx_back/db/handler"
	_ "szyx_back/db/redis"
	_ "szyx_back/models"
	_ "szyx_back/routers"
)

func main() {

	//beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	beego.Run()
	//kdxf.Test()
}
