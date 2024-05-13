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
创建会议，保存到数据库中
*/
func CreateMeeting(meetingDto *meeting.Meeting) (err error) {
	//创建会议
	err = db.CreateMeeting(meetingDto)
	return err
}

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
会议修改
*/
func ModifyMeeting(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {
	res, err = db.ModifyMeeting(meetingDto)
	return res, err
}

/**
语音  会议转译
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
	responseStr := common.DoHttpPost_Audio(audioFullPath, speechDto.FileName)

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
*/
func CreateAudioMeetingSummary(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByFileId(speechDto.FileId)
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

需要先 生成摘要 才能调用这个方法

语音转译后  生成会议纪要
*/
func CreateAudioMeetingMinutes(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByFileId(speechDto.FileId)
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
*/
func CreateAudioMeetingBrainMap(speechDto *kdxf.Kdxf_audio_param) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 根据文件ID，得到orderId
	kdxfSpeech, err := db.GetOrderIdByFileId(speechDto.FileId)
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
	kdxfSpeech, err := db.GetOrderIdByFileId(speechDto.FileId)
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
