package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

func LogConfigInit2() {
	_ = logs.SetLogger(logs.AdapterConsole)
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/my.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365,"color":true,"separate":["error", "warning", "info", "debug"]}`)
	if err != nil {
		panic(fmt.Sprintf("初始化日志失败: %v", err))
	}

	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出log
	//logs.Async()
	fmt.Println("初始化日志成功")

}
func KpHandler2() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root",
		"renhao666",
		"124.71.171.166",
		"3306",
		"hotel_sales",
	)
	myEngine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		logs.Error(err)
		return
	}

	defer func(myEngine *xorm.Engine) {
		err := myEngine.Close()
		if err != nil {
			logs.Error(err)
		}
	}(myEngine)

	// 测试数据库是否连通
	var result string
	_, err = myEngine.SQL("SELECT 1").Get(&result)
	if err != nil {
		logs.Error("Failed to connect to database:", err)
		return
	}

	logs.Info("Database connection successful")

}

func main() {
	//Execute()
	LogConfigInit2()
	KpHandler2()
}
