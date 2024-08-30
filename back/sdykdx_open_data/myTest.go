package main

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"sdykdx_open_data/common"
)

func init2() {
	LogConfigInit2()
	//TokenInit()
}
func LogConfigInit2() {
	_ = logs.SetLogger(logs.AdapterConsole)
	_ = logs.SetLogger(logs.AdapterFile, `{"filename":"logs/my.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":365,"color":true,"separate":["error", "warning", "info", "debug"]}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出log
	//logs.Async()
}
func main2() {
	init2()
	defer common.RecoverHandler(func(err error) {

	})

	common.ErrorHandler(fmt.Errorf("errrrrr"), "22222")

}
