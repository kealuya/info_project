package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"szyx_back/common"
	"szyx_back/configs"
	"szyx_back/db"
	"szyx_back/entity/kdxf"
	"szyx_back/entity/meeting"
)

/**
会议列表
*/
func GetMeetingList(meetingDto *meeting.MeetingList_Param) (res meeting.MeetingList_Result, err error) {
	// 原集合
	res, err = db.GetMeetingList(meetingDto)
	//遍历会议list 查会议文件 组装fileList到集合结构中
	for k, v := range res.MeetingList {
		// 创建目标集合
		var meetingListNew []meeting.MeetingFile
		fileList, _ := db.GetMeetingFileList(v.MeetingId)
		meetingListNew = append(meetingListNew, fileList...)
		res.MeetingList[k].MeetingFile = meetingListNew
	}
	return res, err
}

/**
创建会议，保存到数据库中。 调用语音转译接口 把录音转成文字
*/
func CreateMeeting(meetingDto *meeting.Meeting) (err error) {
	//创建会议
	err = db.CreateMeeting(meetingDto)
	//调用语音转译接口 把录音转成文字
	if err == nil || meetingDto.MeetingType == "audio" {
		//遍历音频会议 上传的所有录音文件
		fileList, _ := db.GetMeetingFileList(meetingDto.MeetingId)
		orderId := ""
		for _, v := range fileList {
			//调用语音识别接口
			responseStr := common.DoHttpPost_Audio(v.FileUrl, v.FileName,v.MeetingId,orderId)
			fmt.Println(responseStr)
			//TODO 这块逻需要调用java在联调
			//TODO 目前传了MeetingId和获取的orderId。 需要确认java那边同一orderId的录音，返回转译的问题
			orderId = "假装是获取到的orderId"
		}
	}
	return err
}

/**
1.音频会议通过转译接口【创建会议】把录音文件转成文字。
2.用户调用此接口对会议内容进行生成 会议摘要、会议纪要、会议脑图结构化字符串
3.调用顺序必须是  会议摘要 - 会议纪要 - 会议脑图
*/
func AudioMeeting_Ai_Abstract_Summary_BrainMap(meetingId string) (err error) {

	defer common.RecoverHandler(func(err error) {
		err = err
	})
	//根据会议ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByMeetingId(meetingId)
	//1.会议摘要
	responseZY := common.DoHttpPost_kdxf_audio_summary(kdxfSpeech.OrderId)
	logs.Info("Ai生成摘要返回：" + responseZY)
	ZY_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseZY), &ZY_result)
	//判断获取摘要成功
	if ZY_result.Success == true {
		//2.会议纪要
		responseJY := common.DoHttpPost_kdxf_audio_minutes(kdxfSpeech.OrderId)
		logs.Info("Ai生成纪要返回：" + responseJY)
		JY_result := new(kdxf.Kdxf_audio_result)
		common.Unmarshal([]byte(responseJY), &JY_result)
		//判断获取纪要成功
		if JY_result.Success == true {
			//3.会议脑图
			responseNT := common.DoHttpPost_kdxf_audio_brainMap(kdxfSpeech.OrderId)
			logs.Info("Ai生成脑图返回：" + responseNT)
			NT_result := new(kdxf.Kdxf_audio_result)
			common.Unmarshal([]byte(responseNT), &NT_result)
			//判断获取脑图成功
			if NT_result.Success != true {
				err = errors.New(fmt.Sprintf("会议脑图生成失败:%s", NT_result.Msg))
			}
		} else {
			err = errors.New(fmt.Sprintf("会议纪要生成失败:%s", JY_result.Msg))
		}
	} else {
		err = errors.New(fmt.Sprintf("会议摘要生成失败:%s", ZY_result.Msg))
	}
	return err
}



/**
//TODO 【暂时没调用这个接口，把这块逻辑写在了 创建会议 CreateMeeting接口中】
音频会议 录音文件转译文字
*/
func CreateMeetingTranslation(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})
	//FIXME 语音转文字
	//拼接路径
	//audioFullPath := "/Users/zhanbaohua/webStorm_work/github.com/info_project/back/szyx_back/" + fpath
	audioFullPath := configs.FILE_UPLOAD_URL + speechDto.FileUrl
	//调用的路径
	logs.Info("调用路径" + audioFullPath)
	////调用语音识别
	responseStr := common.DoHttpPost_Audio(audioFullPath, speechDto.FileName,"","")

	kdxf_audio_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseStr), &kdxf_audio_result)
	if kdxf_audio_result.Success == true {
		res.FileName = kdxf_audio_result.FileName
	} else {
		err = errors.New(fmt.Sprintf("语音会议转译生成失败::%s", kdxf_audio_result.Msg))
	}
	return res, err
}



/**
语音转译后  生成会议摘要
//TODO 【暂时没调用这个接口，把这块逻辑写在了同一个接口AudioMeeting_Ai_Abstract_Summary_BrainMap方法中】
*/
func CreateAudioMeetingSummary(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByMeetingId(speechDto.MeetingId)
	responseStr := common.DoHttpPost_kdxf_audio_summary(kdxfSpeech.OrderId)

	Kdxf_audio_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseStr), &Kdxf_audio_result)
	if Kdxf_audio_result.Success == true {
		res.MeetingId = speechDto.MeetingId
	} else {
		err = errors.New(fmt.Sprintf("会议摘要生成失败::%s", Kdxf_audio_result.Msg))
	}
	return res, err
}

/**
需要先 生成摘要 才能调用这个方法 语音转译后  生成会议纪要
//TODO  【暂时没调用这个接口，把这块逻辑写在了同一个接口AudioMeeting_Ai_Abstract_Summary_BrainMap方法中】

*/
func CreateAudioMeetingMinutes(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByMeetingId(speechDto.MeetingId)
	responseStr := common.DoHttpPost_kdxf_audio_minutes(kdxfSpeech.OrderId)

	Kdxf_audio_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseStr), &Kdxf_audio_result)
	if Kdxf_audio_result.Success == true {
		res.MeetingId = speechDto.MeetingId
	} else {
		err = errors.New(fmt.Sprintf("会议纪要生成失败::%s", Kdxf_audio_result.Msg))
	}
	return res, err
}

/**
语音  会议脑图
//TODO 【暂时没调用这个接口，把这块逻辑写在了同一个接口AudioMeeting_Ai_Abstract_Summary_BrainMap方法中】
*/
func CreateAudioMeetingBrainMap(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByMeetingId(speechDto.MeetingId)
	responseStr := common.DoHttpPost_kdxf_audio_brainMap(kdxfSpeech.OrderId)

	Kdxf_audio_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseStr), &Kdxf_audio_result)
	if Kdxf_audio_result.Success == true {
		res.MeetingId = speechDto.MeetingId
	} else {
		err = errors.New(fmt.Sprintf("会议纪要生成失败::%s", Kdxf_audio_result.Msg))
	}
	return res, err
}

//==========================================文档会议 相关操作 ============================

/**
文档会议   生成会议纪要
*/
func CreateDocumentMeetingMinutes(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	////FIXME 会议文件地址
	responseStr := common.DoHttpPost_kdxf_document_minutes(speechDto.MeetingId, speechDto.FileUrl)

	kdxf_document_result := new(kdxf.Kdxf_document_result)
	common.Unmarshal([]byte(responseStr), &kdxf_document_result)
	if kdxf_document_result.Success == true {
		res.MeetingId = speechDto.MeetingId
	} else {
		err = errors.New(fmt.Sprintf("会议纪要生成失败::%s", kdxf_document_result.Msg))
	}
	return res, err
}

/**
文档  会议 会议脑图
*/
func CreateDocumentMeetingBrainMap(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByMeetingId(speechDto.MeetingId)
	responseStr := common.DoHttpPost_kdxf_audio_minutes(kdxfSpeech.OrderId)

	Kdxf_audio_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseStr), &Kdxf_audio_result)
	if Kdxf_audio_result.Success == true {
		res.MeetingId = speechDto.MeetingId
	} else {
		err = errors.New(fmt.Sprintf("会议纪要生成失败::%s", Kdxf_audio_result.Msg))
	}
	return res, err
}

/**
会议详情
*/
func GetMeetingDetails(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {
	res, err = db.GetMeetingDetails(meetingDto)
	return res, err
}

/**
上传文件，文件基础信息存表。用于选择业务内容关联展示
*/
func AddMeetingFileInfo(meetingFile *meeting.MeetingFile) (err error) {
	//创建会议
	err = db.AddMeetingFileInfo(meetingFile)
	return err
}

/**
会议修改
*/
func ModifyMeeting(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {
	res, err = db.ModifyMeeting(meetingDto)
	return res, err
}

/**
删除音频会议上传的 音频文件
*/
func DeleteAudioMeetingFile(meetingFile *meeting.MeetingFile) (err error) {
	err = db.DeleteAudioMeetingFile(meetingFile)
	return err
}
