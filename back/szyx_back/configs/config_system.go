package configs

import (
	"github.com/astaxie/beego"
)

//sentry 是否发送
var SENTRY_FLAG = true

//log 是否打印sql(true 打印；false 不打印)
var LOG_SQL_FLAG = false

//数据库连接信息
var DB_dataSource = ""

//JWT 私钥
var JWT_SECRET = ""

func init() {
	SENTRY_FLAG, _ = beego.AppConfig.Bool("SENTRY_FLAG")
	LOG_SQL_FLAG, _ = beego.AppConfig.Bool("LOG_SQL_FLAG")
	DB_dataSource = beego.AppConfig.String("DB_dataSource")
	JWT_SECRET = beego.AppConfig.String("JWT_SECRET")

}
