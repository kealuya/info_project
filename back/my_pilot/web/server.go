package web

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"log"
	configs "my_pilot/config"
)

func ServerStart() {

	config := configs.GetConfig()
	addr := config["server"]["port"]
	mode := config["server"]["mode"]
	gin.SetMode(mode)

	logs.Info("服务启动 - 端口:" + addr + " 模式:" + mode)

	r := gin.New()
	// logger 中间件
	r.Use(ginLogger)
	r.Use(gin.Recovery())
	server := SetupRouter(r)

	err := server.Run(addr)
	if err != nil {
		logs.Error("服务器启动失败:", err)
		log.Panicln(err)
	}
}
