package configs

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

func init() {
	LogConfigInit()
}
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
