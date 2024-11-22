package main

import (
	"github.com/beego/beego/v2/core/logs"
	_ "my_pilot/config"
	"my_pilot/web"
	_ "net/http/pprof"
)

func main() {

	logs.Info("系统启动...")
	web.ServerStart()
}
