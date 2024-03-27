package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"path"
	"szyx_back/common"
)

type MeetingCtrl struct {
	beego.Controller
}

//// @Title 会议创建
//// @Tags CreateMeeting
//// @Summary 会议创建
//// @accept application/json
//// @Produce application/json
//// @Param data body meeting.Meeting true "Meeting struct"
//// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
//// @router /createMeeting [post]
//func (MeetingCtrl *MeetingCtrl) CreateMeeting() {
//	resJson := NewJsonStruct(nil)
//	defer func() {
//		MeetingCtrl.Data["json"] = string(common.Marshal(resJson))
//		MeetingCtrl.ServeJSON()
//	}()
//	metting := new(meeting.Meeting)
//	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
//	common.Unmarshal(jsonByte, &metting)
//	logs.Info("创建会议入参：" + string(jsonByte))
//	//业务处理
//	meetingRes, err := models.CreateMeeting(metting)
//	if err == nil {
//		resJson.Success = true
//		resJson.Data = meetingRes
//	} else {
//		resJson.Success = false
//		resJson.Msg = fmt.Sprintf("会议创建失败::%s", err)
//	}
//}

// @Title 会议音频上传
// @Tags UploadMeetingAudioFile
// @Summary 会议音频上传
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "file文件,会议ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @router /uploadMeetingAudioFile [post]
func (MeetingCtrl *MeetingCtrl) UploadMeetingAudioFile() {

	MeetingCtrl.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")

	resJson := NewJsonStruct(nil)

	flag := true

	defer func() {
		MeetingCtrl.Data["json"] = string(common.Marshal(resJson))
		MeetingCtrl.ServeJSON()
	}()

	file, h, _ := MeetingCtrl.GetFile("file") //获取上传的文件
	//var userCorpCode = MeetingCtrl.GetString("meetingId") //获取会议ID

	ext := path.Ext(h.Filename)
	//fmt.Println("------------" + ext)

	//验证后缀名是否符合要求   mp3/blob
	AllowExtMap := map[string]bool{
		".mp3":  true,
		".blob": true,
	}

	if _, ok := AllowExtMap[ext]; !ok {
		flag = false
		resJson.Success = false
		resJson.Msg = "文件类型不正确"
	}

	uploadDir := "static/upload/"

	////创建目录 不需要
	//uploadDir := "static/upload/"
	//err := os.MkdirAll(uploadDir, 777)
	//if err != nil {
	//	MeetingCtrl.Ctx.WriteString(fmt.Sprintf("%v"))
	//	return
	//}

	//构造上传路径+文件名
	fpath := uploadDir + h.Filename
	defer file.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况

	err := MeetingCtrl.SaveToFile("file", fpath)
	if err != nil {
		MeetingCtrl.Ctx.WriteString(fmt.Sprintf("%v"))
		flag = false
	}

	if flag {
		resJson.Success = true
		resJson.Msg = "上传成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议录音上传失败::%s", err)
	}

}

//TODO 会议列表

//TODO 会议修改

//TODO 会议纪要

//TODO 会议脑图
