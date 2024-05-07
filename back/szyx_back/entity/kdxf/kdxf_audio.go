package kdxf

import "szyx_back/entity"

//语音识别入参-生成会议记录
type Kdxf_audio_param struct {
	UserId    string `json:"userId" description:"用户ID" `
	CorpCode  string `json:"corpCode" description:"企业ID" `
	MeetingId string `json:"meetingId" description:"会议ID" `
	FileId    string `json:"fileId" description:"文件ID" `
	FileUrl   string `json:"fileUrl" description:"文件地址" `
	FileName  string `json:"fileName" description:"文件名称" `
}

//语音识别返回值
type Kdxf_audio_result struct {
	Msg      string `json:"msg" description:"msg" `
	FileName string `json:"fileName" description:"文件名" `
	OrderId  string `json:"orderId" description:"ID" `
	Success  string `json:"success" description:"" `
}

//录音文件
type Kdxf_speech struct {
	entity.Base         //预留多公司模式
	OrderId      string `json:"orderId" description:"语音订单唯一编号" `
	No           string `json:"no" description:"用来处理多语言归属同一会议的情况" `
	FileName     string `json:"fileName" description:"存储文件的文件名称" `
	State        string `json:"state" description:"0：进行中；1：完成；2：错误" `
	DateTime     string `json:"dateTime" description:"发生时间" `
	Content      string `json:"content" description:"语音转换内容" `
	Comment      string `json:"comment" description:"备注" `
	RealDuration string `json:"realDuration" description:"语音时长" `
	UserId       string `json:"userId" description:"用户ID" `
	UserName     string `json:"userName" description:"用户姓名" `
	UserMobile   string `json:"userMobile" description:"用户手机号" `
	MeetingId    string `json:"meetingId" description:"会议id" `
	MeetingTitle string `json:"meetingTitle" description:"会议标题" `
	UseFlag      string `json:"useFlag" description:"使用状态" `
	TaskId       string `json:"taskId" description:"任务ID" `
	FileUrl      string `json:"fileUrl" description:"文件存放路径" `
}

//FIXME 考虑文件跟语音分开存储，方便后面ai识别使用
//文档文件  基本上是word,txt等文字性质的文件
type Kdxf_File struct {
	entity.Base         //预留多公司模式
	OrderId      string `json:"orderId" description:"文件订单唯一编号" `
	No           string `json:"no" description:"用来处理多语言归属同一会议的情况" `
	FileName     string `json:"fileName" description:"存储文件的文件名称" `
	State        string `json:"state" description:"0：进行中；1：完成；2：错误" `
	DateTime     string `json:"dateTime" description:"发生时间" `
	Content      string `json:"content" description:"文件提纲" `
	Comment      string `json:"comment" description:"备注" `
	RealDuration string `json:"realDuration" description:"文件大小" `
	UserId       string `json:"userId" description:"用户ID" `
	UserName     string `json:"userName" description:"用户姓名" `
	UserMobile   string `json:"userMobile" description:"用户手机号" `
	MeetingId    string `json:"meetingId" description:"会议id" `
	MeetingTitle string `json:"meetingTitle" description:"会议标题" `
	UseFlag      string `json:"useFlag" description:"使用状态" `
	TaskId       string `json:"taskId" description:"任务ID" `
	FileUrl      string `json:"fileUrl" description:"文件存放路径" `
}
