package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
	"strings"
)

type LogInfoController struct {
	beego.Controller
}

func (logInfo *LogInfoController) GetLogInfo() {
	str := getLogInfo()
	logInfo.Data["json"] = str
	logInfo.ServeJSON()
}

// @Title 查看日志文件
// @Tags getLogInfo
// @Summary 查看日志文件
// @Security token
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @router /getLogInfo [get]
func getLogInfo() (result string) {
	f, err := os.Open("D:/Git/szyx_back/logs/project.log") //日志文件路径

	if err != nil {
		fmt.Println("文件读取失败")
		return "文件读取失败~"
	}

	defer f.Close()
	buf := make([]byte, 1024)
	var strContent = ""
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break
		}
		//strContent += string(buf[0:n]);
		//strContent += strings.Replace(string(buf[0:n]), "\n", "", -1)
		strContent += strings.Replace(string(buf[0:n]), "\u001b", "", -1)
	}

	var strContentd = strings.Replace(strContent, "\n2021", "", -1)
	//return strings.Replace(strContent, "\n", "", -1);
	return strContentd
}
