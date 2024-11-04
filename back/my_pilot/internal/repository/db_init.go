package repository

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"my_pilot/config"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

var dbEngine *xorm.Engine

func InitDbEngine() {
	dbConfig := configs.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["database"]["username"],
		dbConfig["database"]["password"],
		dbConfig["database"]["host"],
		dbConfig["database"]["port"],
		dbConfig["database"]["name"],
	)
	myEngine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logs.Error(err.Error())
		log.Panicln(err)
	}

	// 创建 SimpleLogger，将beego logger writer配置给xorm，用来打印日志
	logWriter := xlog.NewSimpleLogger(logs.GetLogger().Writer())
	myEngine.SetLogger(logWriter)
	myEngine.ShowSQL(false)
	err = myEngine.Ping()
	if err != nil {
		logs.Error(err.Error())
		log.Panicln(err)
	}
	logs.Info("初始化数据库...")
	dbEngine = myEngine
}
