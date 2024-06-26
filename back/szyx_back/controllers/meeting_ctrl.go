package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"path"
	"szyx_back/common"
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
	err,mag := models.CreateMeeting(metting)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "会议创建成功" + mag
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

	defer func() {
		MeetingCtrl.Data["json"] = resJson
		logs.Warn(resJson)
		MeetingCtrl.ServeJSON()
	}()

	file, h, _ := MeetingCtrl.GetFile("file")       //获取上传的文件
	meetingId := MeetingCtrl.GetString("meetingId") //获取会议ID
	userId := MeetingCtrl.GetString("userId")       //用户ID
	corpCode := MeetingCtrl.GetString("corpCode")   //企业code
	audioTime := MeetingCtrl.GetString("audioTime") //音频时长

	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求   mp3/blob
	AllowExtMap := map[string]bool{
		".mp3":  true,
		".blob": true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		resJson.Success = false
		resJson.Msg = "文件类型不正确"
		return
	}
	//这是上传本地代码 改为上传obs
	/*uploadDir := "static/upload/"
	//构造上传路径+文件名
	filePath := uploadDir + h.Filename
	defer file.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况

	err := MeetingCtrl.SaveToFile("file", filePath)
	if err != nil {
		MeetingCtrl.Ctx.WriteString(fmt.Sprintf("%v"))
		flag = false
	}*/
	//调用华为云OBS上传音频文件 //音频传入audioFile文件夹
	obsPutFileRep, err := common.ObsPutFile(corpCode, h.Filename, file,"audioFile")
	if obsPutFileRep.Success {
		//上传的音频文件基础信息存表。用于选择业务内容关联展示
		meetingFile := new(meeting.MeetingFile)
		meetingFile.MeetingId = meetingId
		meetingFile.FileUrl = obsPutFileRep.ImgURL
		meetingFile.FileName = h.Filename
		meetingFile.Creater = userId
		meetingFile.CorpCode = corpCode
		meetingFile.FileType = "mp3" //音频文件
		meetingFile.AudioTime = audioTime
		err2 := models.AddMeetingFileInfo(meetingFile)
		if err2 == nil {
			fileStruct := new(meeting.MeetingFile_Result)
			fileStruct.FileUrl = obsPutFileRep.ImgURL
			resJson.Success = true
			resJson.Msg = "录音上传OBS成功"
			resJson.Data = fileStruct
		} else {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("会议文件信息保存失败", err2)
		}
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("上传华为obs失败::%s", err)
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

	defer func() {
		MeetingCtrl.Data["json"] = resJson
		logs.Warn(resJson)
		MeetingCtrl.ServeJSON()
	}()

	file, h, _ := MeetingCtrl.GetFile("file")       //获取上传的文件
	meetingId := MeetingCtrl.GetString("meetingId") //获取会议ID
	userId := MeetingCtrl.GetString("userId")       //用户ID
	corpCode := MeetingCtrl.GetString("corpCode")   //企业code

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
		resJson.Success = false
		resJson.Msg = "文件类型不正确"
		return
	}
	//调用华为云OBS上传音频文件  //文档传入documentFile文件夹
	obsPutFileRep, err := common.ObsPutFile(corpCode, h.Filename, file,"documentFile")
	if obsPutFileRep.Success {
		//上传文件的基础信息存表。用于选择业务内容关联展示
		meetingFile := new(meeting.MeetingFile)
		meetingFile.MeetingId = meetingId
		meetingFile.FileUrl = obsPutFileRep.ImgURL
		meetingFile.FileName = h.Filename
		meetingFile.Creater = userId
		meetingFile.CorpCode = corpCode
		meetingFile.FileType = "word" //文档文件
		err2 := models.AddMeetingFileInfo(meetingFile)
		if err2 == nil {
			fileStruct := new(meeting.MeetingFile_Result)
			fileStruct.FileUrl = obsPutFileRep.ImgURL
			resJson.Success = true
			resJson.Msg = "文件上传OBS成功"
			resJson.Data = fileStruct
		} else {
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

// @Title 会议录音生成 会议摘要
// @Tags AudioMeeting_Ai_Abstract
// @Summary 会议录音生成 会议摘要
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Ai生成成功"}"
// @router /audioMeeting_Ai_Abstract [post]
func (MeetingCtrl *MeetingCtrl) AudioMeeting_Ai_Abstract() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	speech := new(kdxf.Kdxf_audio_param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("生成会议摘要入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &speech)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.AudioMeeting_Ai_Abstract(speech.MeetingId)
	if true {
		resJson.Success = true
		resJson.Msg = "正在生成摘要，请稍等..."
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("Ai生成内容报错:%s", err)
	}
}

// @Title 会议录音生成 会议纪要 + 会议脑图
// @Tags AudioMeeting_Ai_Summary_BrainMap
// @Summary 会议录音生成 会议纪要 + 会议脑图
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"Ai生成成功"}"
// @router /audioMeeting_Ai_Summary_BrainMap [post]
func (MeetingCtrl *MeetingCtrl) AudioMeeting_Ai_Summary_BrainMap() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	speech := new(kdxf.Kdxf_audio_param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("生成会议纪要、脑图入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &speech)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.AudioMeeting_Ai_Summary_BrainMap(speech.MeetingId)
	if true {
		resJson.Success = true
		resJson.Msg = "纪要、脑图获取成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("Ai生成内容报错:%s", err)
	}
}


// TODO 【暂时没调用这个接口，把这块逻辑写在了 创建会议 CreateMeeting接口中】
// @Title 音频会议 录音文件转译文字 【暂时没调用这个接口，把这块逻辑写在了 创建会议接口中】
// @Tags CreateMeetingTranslation
// @Summary 音频会议 录音文件转译文字 【暂时没调用这个接口，把这块逻辑写在了 创建会议接口中】
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"语音会议转译成功"}"
// @router /createMeetingTranslation [post]
func (MeetingCtrl *MeetingCtrl) CreateMeetingTranslation() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	speech := new(kdxf.Kdxf_audio_param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("语音会议转译入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &speech)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateMeetingTranslation(speech)

	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("语音会议转译失败::%s", err)
	}
}

// TODO 暂时没调用这个接口，把这块逻辑写在了同一个接口AudioMeeting_Ai_Abstract_Summary_BrainMap方法中
// @Title 语音会议 摘要 生成
// @Tags CreateAudioMeetingSummary
// @Summary 语音 会议摘要
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议摘要生成成功"}"
// @router /createAudioMeetingSummary [post]
func (MeetingCtrl *MeetingCtrl) CreateAudioMeetingSummary() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	speech := new(kdxf.Kdxf_audio_param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("生成会议摘要入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &speech)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateAudioMeetingSummary(speech)

	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议摘要生成失败::%s", err)
	}
}

// TODO 暂时没调用这个接口，把这块逻辑写在了同一个接口AudioMeeting_Ai_Abstract_Summary_BrainMap方法中
// @Title 语音会议纪要生成
// @Tags CreateAudioMeetingMminutes
// @Summary 会议纪要
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议纪要生成成功"}"
// @router /createAudioMeetingMinutes [post]
func (MeetingCtrl *MeetingCtrl) CreateAudioMeetingMinutes() {
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
	res, err := models.CreateAudioMeetingMinutes(speech)

	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议纪要生成失败::%s", err)
	}
}

// TODO 暂时没调用这个接口，把这块逻辑写在了同一个接口AudioMeeting_Ai_Abstract_Summary_BrainMap方法中
// @Title 语音 会议脑图，用于根据会议纪要得到会议脑图
// @Tags CreateAudioMeetingBrainMap
// @Summary 语音 会议脑图
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议脑图生成成功"}"
// @router /createAudioMeetingBrainMap [post]
func (MeetingCtrl *MeetingCtrl) CreateAudioMeetingBrainMap() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	speech := new(kdxf.Kdxf_audio_param)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("生成语音会议脑图 入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &speech)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateAudioMeetingBrainMap(speech)

	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("语音会议脑图生成失败::%s", err)
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

//===================================================文档会议操作=====================

// @Title 文档会议纪要生成
// @Tags CreateDocumentMeetingMinutes
// @Summary 会议纪要
// @accept application/json
// @Produce application/json
// @Param data body kdxf.Kdxf_audio_param true "Kdxf_audio_param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"会议纪要生成成功"}"
// @router /createDocumentMeetingMinutes [post]
func (MeetingCtrl *MeetingCtrl) CreateDocumentMeetingMinutes() {
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
	res, err := models.CreateDocumentMeetingMinutes(speech)

	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("会议纪要生成失败::%s", err)
	}
}

// @Title 删除音频会议上传的 音频文件
// @Tags DeleteAudioMeetingFile
// @Summary 删除音频 
// @accept application/json
// @Produce application/json
// @Param data body meeting.Meeting true "Meeting struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @router /deleteAudioMeetingFile [post]
func (MeetingCtrl *MeetingCtrl) DeleteAudioMeetingFile() {
	resJson := NewJsonStruct(nil)
	defer func() {
		MeetingCtrl.Data["json"] = resJson
		MeetingCtrl.ServeJSON()
	}()
	meetingFile := new(meeting.MeetingFile)
	var jsonByte = MeetingCtrl.Ctx.Input.RequestBody
	logs.Info("删除音频文件入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &meetingFile)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.DeleteAudioMeetingFile(meetingFile)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "删除成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("删除音频文件失败::%s", err)
	}
}
