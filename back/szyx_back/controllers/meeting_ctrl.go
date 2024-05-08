package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"path"
	"szyx_back/entity/kdxf"
	"szyx_back/entity/meeting"
	"szyx_back/models"
)

type MeetingCtrl struct {
	beego.Controller
}

// @Title 会议创建
// @Tags CreateMeeting
// @Summary 会议创建
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "Meeting struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @router /createMeeting [post]
func (MeetingCtrl *MeetingCtrl) CreateMeeting() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	metting := new(meeting.Meeting)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("创建会议入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &metting)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.CreateMeeting(metting)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "会议创建成功"
		//resJson.Data = meetingRes
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议创建失败::%s", err)
	}
}

// @Title 会议音频上传
// @Tags UploadMeetingAudioFile
// @Summary 会议音频上传
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "file文件,会议ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @router /uploadMeetingAudioFile [post]
func (MeetingCtrl *MeetingCtrl) UploadMeetingAudioFile() {

	//设置允许跨域请求
	MeetingCtrl.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	resJson := NewJsonStruct(nil)
	flag := true

	defer func() {
		MeetingCtrl.Data["json"] = resJson
		logs.Warn(resJson)
		MeetingCtrl.ServeJSON()
	}()

	file, h, _ := MeetingCtrl.GetFile("file")       //获取上传的文件
	meetingId := MeetingCtrl.GetString("meetingId") //获取会议ID
	userId := MeetingCtrl.GetString("userId")       //用户ID

	fmt.Println("------------" + meetingId)
	fmt.Println("------------" + userId)

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
		return
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
		fileStruct := new(meeting.MeetingFile_Result)
		fileStruct.FileUrl = fpath
		resJson.Success = true
		resJson.Msg = "上传成功"
		resJson.Data = fileStruct
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议录音上传失败::%s", err)
	}

}

// @Title 会议文件上传
// @Tags UploadMeetingFile
// @Summary 会议文件上传
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "file文件,会议ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"上传成功"}"
// @router /uploadMeetingFile [post]
func (MeetingCtrl *MeetingCtrl) UploadMeetingFile() {

	//设置允许跨域请求
	MeetingCtrl.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	resJson := NewJsonStruct(nil)
	flag := true

	defer func() {
		MeetingCtrl.Data["json"] = resJson
		logs.Warn(resJson)
		MeetingCtrl.ServeJSON()
	}()

	file, h, _ := MeetingCtrl.GetFile("file")       //获取上传的文件
	meetingId := MeetingCtrl.GetString("meetingId") //获取会议ID
	meetingTitle := MeetingCtrl.GetString("meetingTitle") //会议标题
	userId := MeetingCtrl.GetString("userId")       //用户ID
	corpCode := MeetingCtrl.GetString("corpCode")     //企业code

	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求   doc/txt
	AllowExtMap := map[string]bool{
		".doc":  true,
		".docx": true,
		".txt":  true,
		".xls":  true,
		".xlsx": true,
		".pdf":  true,
	}
	//判读所属类型
	if _, ok := AllowExtMap[ext]; !ok {
		flag = false
		resJson.Success = false
		resJson.Msg = "文件类型不正确"
		return
	}
	uploadDir := "static/upload/"
	//构造上传路径+文件名
	filePath := uploadDir + h.Filename
	defer file.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况

	err := MeetingCtrl.SaveToFile("file", filePath)
	if err != nil {
		MeetingCtrl.Ctx.WriteString(fmt.Sprintf("%v"))
		flag = false
	}
	if flag {
		//上传文件的基础信息存表。用于选择业务内容关联展示
		meetingFile := new(meeting.MeetingFile)
		meetingFile.MeetingId = meetingId
		meetingFile.MeetingTitle = meetingTitle
		meetingFile.FileUrl = filePath //TODO 先取构造的上传路径 日后有了文件obs地址在改
		meetingFile.FileName = h.Filename
		meetingFile.Creater = userId
		meetingFile.CorpCode = corpCode
		err2 := models.AddMeetingFileInfo(meetingFile)
		if err2 == nil {
			fileStruct := new(meeting.MeetingFile_Result)
			fileStruct.FileUrl = filePath
			resJson.Success = true
			resJson.Msg = "上传成功"
			resJson.Data = fileStruct
		}else{
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("会议文件信息保存失败", err2)
		}
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议文件上传失败::%s", err)
	}

}

// @Title 会议列表
// @Tags GetMeetingList
// @Summary 会议列表
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "Meeting struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getMeetingList [post]
func (MeetingCtrl *MeetingCtrl) GetMeetingList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	metting_param := new(meeting.MeetingList_Param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("查询会议列表入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &metting_param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetMeetingList(metting_param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "获取会议列表成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议列表查询失败::%s", err)
	}
}

// @Title 会议修改，用于修改录音文件地址
// @Tags ModifyMeeting
// @Summary 会议修改
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "Meeting struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"修改会议成功"}"
// @router /modifyMeeting [post]
func (MeetingCtrl *MeetingCtrl) ModifyMeeting() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	metting := new(meeting.Meeting)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("修改会议入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &metting)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.ModifyMeeting(metting)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议修改失败::%s", err)
	}
}

// @Title 会议纪要，用于根据录音得到会议纪要
// @Tags CreateMeetingMminutes
// @Summary 会议纪要
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议纪要生成成功"}"
// @router /createMeetingMminutes [post]
func (MeetingCtrl *MeetingCtrl) CreateMeetingMminutes() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	speech := new(kdxf.Kdxf_audio_param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("生成会议纪要入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &speech)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateMeetingMminutes(speech)

	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议纪要生成失败::%s", err)
	}
}

// @Title 会议脑图，用于根据会议纪要得到会议脑图
// @Tags CreateMeetingBrainMap
// @Summary 会议脑图
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "Meeting struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议脑图生成成功"}"
// @router /createMeetingBrainMap [post]
func (MeetingCtrl *MeetingCtrl) CreateMeetingBrainMap() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	metting := new(meeting.Meeting)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("生成会议脑图入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &metting)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateMeetingBrainMap(metting)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议脑图生成失败::%s", err)
	}
}

// @Title 会议详情
// @Tags GetMeetingDetails
// @Summary 会议详情
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "Meeting struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议详情查询成功"}"
// @router /getMeetingDetails [post]
func (MeetingCtrl *MeetingCtrl) GetMeetingDetails() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	metting := new(meeting.Meeting)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("会议详情入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &metting)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetMeetingDetails(metting)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议详情获取失败::%s", err)
	}
}
