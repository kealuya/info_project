package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"llm/chat"
	"time"
	"xorm.io/xorm"
)

func LogConfigInit() {
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

func KpHandler() {
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
	}
	defer myEngine.Close()

	queryString, err := myEngine.QueryString("select * from kp order by id")
	if err != nil {
		logs.Error(err)
	}
	//m := queryString[0]

	for _, m := range queryString {

		bt := m["bt"]
		id := m["id"]

		start := time.Now()

		c := chat.NewInstance(string(""))
		br := c.Conversation(bt)
		elapsed := time.Since(start)
		seconds := elapsed.Seconds()
		logs.Info("%v 耗时: %.2f秒\n", m["ywid"], seconds) // 保留2位小数
		//fmt.Printf("%s", br.Answer)
		_, err = myEngine.Exec("update kp set llm_nr = ?,elapsed=?,token=? where id = ?", br.Answer, fmt.Sprintf("%.2f", seconds), br.Metadata.Usage.TotalTokens, id)
		if err != nil {
			logs.Error(err)
		}
		logs.Info("id:%v更新成功", id)
		//if i == 2 {
		//	break
		//}
	}
}

func main() {
	//Execute()
	LogConfigInit()
	KpHandler()
}
