package main

import (
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	_ "github.com/go-sql-driver/mysql"
	"llm/chat"
	"slices"
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

	// 设置连接池参数
	myEngine.SetMaxIdleConns(10)           // 最大空闲连接数
	myEngine.SetMaxOpenConns(100)          // 最大连接数
	myEngine.SetConnMaxLifetime(time.Hour) // 连接最大生存时间

	defer myEngine.Close()

	queryString, err := myEngine.QueryString("select * from kp order by id")
	if err != nil {
		logs.Error(err)
	}
	//m := queryString[0]

	for i, m := range queryString {

		nr := m["nr"]
		id := m["id"]

		replayIds := []string{"187", "188", "189", "191",
			"192",
			"193",
			"194",
			"195",
			"196"}

		if !slices.Contains(replayIds, id) {
			continue
		}

		start := time.Now()

		c := chat.NewInstance(string(""))
		br, errConversation := c.Conversation(nr)
		if errConversation != nil {
			logs.Error(m["id"], errConversation)
			continue
		}
		elapsed := time.Since(start)
		seconds := elapsed.Seconds()
		logs.Info("%v 耗时: %.2f秒\n", m["ywid"], seconds) // 保留2位小数

		logs.Info(br.Answer)

		mm := make(map[string]string)

		errJson := json.Unmarshal([]byte(br.Answer), &mm)
		if errJson != nil {
			logs.Error(id, errJson)
			continue
		}

		//fmt.Printf("%s", br.Answer)
		_, err = myEngine.Exec("update kp set llm_nr = ?,llm_bt = ?,elapsed=?,token=? where id = ?", mm["summary"], mm["title"], fmt.Sprintf("%.2f", seconds), br.Metadata.Usage.TotalTokens, id)
		if err != nil {
			logs.Error(err)
		}
		logs.Info("id:%v更新成功", id)
		if i == 2 {
			//break
		}
	}
}

func main() {
	//Execute()
	LogConfigInit()
	KpHandler()
}
