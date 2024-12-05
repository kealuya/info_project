package repository

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"my_pilot/config"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

func init() {
	InitDbEngine()
}

var dbEngine *xorm.Engine

func InitDbEngine() {
	dbConfig := configs.GetConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["database_shop"]["username"],
		dbConfig["database_shop"]["password"],
		dbConfig["database_shop"]["host"],
		dbConfig["database_shop"]["port"],
		dbConfig["database_shop"]["name"],
	)
	myEngine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logs.Error(err.Error())
		log.Panicln(err)
	}
	// 设置连接池参数
	myEngine.SetMaxOpenConns(500)                   // 最大连接数
	myEngine.SetMaxIdleConns(100)                   // 最大空闲连接数
	myEngine.SetConnMaxLifetime(3600 * time.Second) // 连接最大生命周期
	myEngine.SetConnMaxIdleTime(1800 * time.Second) // 空闲连接最大生命周期

	// 性能优化设置
	myEngine.DatabaseTZ = time.Local // 设置数据库时区
	myEngine.TZLocation = time.Local // 设置本地时区
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
