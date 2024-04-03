package common

import (
	"bytes"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

//kdxf音频转换地址
var kdxf_audio = "http://localhost:7003/uploadAudio"

//kdxf音频转换项目，测试地址
var test_audio = "http://localhost:7003/hello"

//kdxf音频转换
func DoHttpPost_Audio(fileUrl string, filename string, options ...map[string]string) (respBody string) {

	errMsg := `{"msg":"文件上传失败","success":"false"}`

	// 打开文件
	file, err := os.Open(fileUrl)
	if err != nil {
		ErrorHandler(err, "文件未找到"+fileUrl)
		return errMsg
	}
	defer file.Close()

	// 准备表单数据
	var buffer bytes.Buffer
	multipartWriter := multipart.NewWriter(&buffer)

	// 创建文件表单项
	fileWriter, err := multipartWriter.CreateFormFile("file", filename)
	if err != nil {
		ErrorHandler(err, "文件未找到"+fileUrl)
		return errMsg
	}

	// 将文件内容复制到表单项
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		ErrorHandler(err, "请求表单构建失败"+fileUrl)
		return errMsg
	}

	// 关闭 multipart writer
	multipartWriter.Close()

	// 创建请求
	request, err := http.NewRequest("POST", kdxf_audio, &buffer)
	if err != nil {
		ErrorHandler(err, "请求构建失败"+fileUrl)
		return errMsg
	}

	// 设置请求头，包括内容类型
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// 发送请求
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		ErrorHandler(err, "请求失败"+fileUrl)
		return errMsg
	}
	defer response.Body.Close()

	// 处理响应
	if response.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(response.Body)
		if err == nil {
			logs.Info(string(body))
			return string(body)
		} else {
			return errMsg
		}

	} else {
		return errMsg
	}

}

//kdxf音频转换
func DoHttpPost_Audio_err(fileUrl string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_audio)
	req.Header("Content-Type", "multipart/form-data")
	req.Header("accept", "*/*")
	req.Header("accept-encoding", "gzip, deflate, br")
	req.Header("connection", "keep-alive")

	//设置连接超时为5分钟，读写超时为5分钟
	req.SetTimeout(1*time.Minute, 1*time.Minute)
	req.PostFile("file", fileUrl)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Warn(req.String())
	return respBody
}

//http Post 请求
func DoHttpPost_audio(url string, param string, options ...map[string]string) (respBody string) {
	req := httplib.Post(test_audio)
	req.Header("Content-Type", "application/json")
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)

	logs.Warn(respBody)
	return respBody
}
