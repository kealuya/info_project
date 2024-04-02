package common

import (
	"github.com/astaxie/beego/httplib"
)

//kdxf音频转换地址
var kdxf_audio = "http://localhost:8080/uploadAudio"

//kdxf音频转换
func DoHttpPost_Audio(fileUrl string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_audio)
	req.Header("Content-Type", "form-data")
	req.PostFile("file", fileUrl)
	req.Debug(true)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}
