package main

import (
	_ "my_pilot/config"
	"my_pilot/internal/handler/batch"
	_ "net/http/pprof"
)

func main() {

	//logs.Info("系统启动...")
	//web.ServerStart()
	batch.Main()
}
