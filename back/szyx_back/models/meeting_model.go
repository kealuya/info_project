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
	defer common.RecoverHandler(func(err error) {
		err = err
	})
	//创建会议
	err = db.CreateMeeting(meetingDto)
	return err
}

/**
会议列表
*/
func GetMeetingList(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {

	return res, err
}

/**
会议修改
*/
func ModifyMeeting(meetingDto *meeting.Meeting) (res meeting.Meeting, err error) {

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
