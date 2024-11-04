package main

import (
	"github.com/beego/beego/v2/core/logs"
	_ "my_pilot/config"
	"my_pilot/web"
)

func main() {

	logs.Info("系统启动...")
	web.ServerStart()

}
