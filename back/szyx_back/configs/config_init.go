package configs

import "github.com/astaxie/beego/logs"

func init() {
	LogConfigInit()
}

func LogConfigInit() {
	_ = logs.SetLogger("console")
	_ = logs.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/project.log",
															"level":7,
															"maxlines":0,
															"maxsize":0,
															"daily":true,
															"maxdays":365,
															"color":true,
															"separate":["error","warning","info","debug"]}`)
	//输出文件名和行号
	logs.EnableFuncCallDepth(true)
	//异步输出log
	//logs.Async()
}
