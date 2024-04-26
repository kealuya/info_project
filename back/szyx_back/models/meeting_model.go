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
会议纪要
*/
func CreateMeetingMminutes(speechDto *kdxf.Kdxf_speech) (res kdxf.Kdxf_speech, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	//FIXME 语音转会议纪要
	//拼接路径
	//audioFullPath := "/Users/zhanbaohua/webStorm_work/github.com/info_project/back/szyx_back/" + fpath
	//audioFullPath := "/Users/zhanbaohua//github.com/info_project/back/kdxf_speech/src/main/audio/合成音频.wav"
	//audioFullPath := "/Users/zhanbaohua/Downloads/demo/audio/3344555.m4a"
	audioFullPath := configs.FILE_UPLOAD_URL + speechDto.FileUrl
	//调用的路径
	logs.Info("调用路径" + audioFullPath)
	////调用语音识别
	responseStr := common.DoHttpPost_Audio(audioFullPath, speechDto.FileName)

	kdxf_audio_result := new(kdxf.Kdxf_audio_result)
	common.Unmarshal([]byte(responseStr), &kdxf_audio_result)
	if kdxf_audio_result.Success == "true" {
		res.FileName = kdxf_audio_result.FileName
	} else {
		err = errors.New(fmt.Sprintf("会议纪要生成失败::%s", kdxf_audio_result.Msg))
	}
	return res, err
}

/**
会议脑图
*/
func CreateMeetingBrainMap(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

	return res, err
}

/**
会议详情
*/
func GetMeetingDetails(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {
	res, err = db.GetMeetingDetails(meetingDto)
	return res, err
}
