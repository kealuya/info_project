package models

import (
	"szyx_back/common"
	"szyx_back/db"
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
	res, err = db.GetMeetingList(meetingDto)
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
func CreateMeetingMminutes(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {
	defer common.RecoverHandler(func(err error) {
		err = err
	})

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
