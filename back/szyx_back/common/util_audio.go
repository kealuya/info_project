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
)

//kdxf音频转换地址
var kdxf_audio = "http://localhost:7003/uploadAudio"

//kdxf音频转换地址---入参meetingId
var kdxf_audio_translation_meeting = "http://localhost:7003/audioMeetingTranslation"

//kdxf 获取音频转译结果
var kdxf_audio_translation_getResult = "http://localhost:7003/getResultByMeetingId"

//kdxf音频会议总结
var kdxf_CreateMeetingSummary = "http://localhost:7003/convertAudioToSummary"

//kdxf音频会议纪要
var kdxf_CreateMeetingMinutes = "http://localhost:7003/convertAudioToMeetingMinute"

//kdxf音频会议脑图
var kdxf_CreateMeetingBrainMap = "http://localhost:7003/convertAudioToBrainMap"

//==================================文档操作=====================

//kdxf 文档  会议纪要生成
var kdxf_CreateMeetingMinutes_document = "http://localhost:7003/convertDocumentToSummary"

//kdxf音频转换项目，测试地址
var test_audio = "http://localhost:7003/hello"

//kdxf音频转换
func DoHttpPost_Audio(fileUrl string, filename string, meetingId string, orderId string, options ...map[string]string) (respBody string) {

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

	// 添加表单字段
	fileWriter, err = multipartWriter.CreateFormField("meetingId")
	fileWriter.Write([]byte(meetingId))

	fileWriter, err = multipartWriter.CreateFormField("orderId")
	fileWriter.Write([]byte(orderId))

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

//kdxf 语音翻译后  生成会议总结
func DoHttpPost_kdxf_audio_summary(orderId string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_CreateMeetingSummary)
	req.Header("Content-Type", "application/json")
	req.Param("orderId", orderId)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Debug(respBody)
	return respBody
}

//kdxf 语音翻译后  生成会议纪要
func DoHttpPost_kdxf_audio_minutes(orderId string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_CreateMeetingMinutes)
	req.Header("Content-Type", "application/json")
	req.Param("orderId", orderId)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Warn("Java系统的返回值：" + respBody)
	return respBody
}

//kdxf 语音翻译后  生成会议脑图  BrainMap
func DoHttpPost_kdxf_audio_brainMap(orderId string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_CreateMeetingBrainMap)
	req.Header("Content-Type", "application/json")
	req.Param("orderId", orderId)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Warn(respBody)
	return respBody
}

//kdxf 语音转译 meetingId 入参
func DoHttpPost_kdxf_audio_translation_meeting(meetingId string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_audio_translation_meeting)
	req.Header("Content-Type", "application/json")
	req.Param("meetingId", meetingId)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Warn(respBody)
	return respBody
}

//kdxf 获取转译结果
func DoHttpPost_kdxf_audio_translation_getResult(meetingId string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_audio_translation_getResult)
	req.Header("Content-Type", "application/json")
	req.Param("meetingId", meetingId)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Warn(respBody)
	return respBody
}

//==========================================文档操作=======================

//kdxf 文档生成会议记录
func DoHttpPost_kdxf_document_minutes(meetingId string, fileUrl string, options ...map[string]string) (respBody string) {
	req := httplib.Post(kdxf_CreateMeetingMinutes_document)
	req.Header("Content-Type", "application/json")
	//req.Body(param)
	req.Param("meetingId", meetingId)
	req.Param("fileUrl", fileUrl)
	respBody, err := req.String()
	ErrorHandler(err)
	logs.Warn(respBody)
	return respBody
}
