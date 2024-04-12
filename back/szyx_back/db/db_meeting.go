package db

import (
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity/meeting"
	"time"
)

//创建会议
func CreateMeeting(info *meeting.Meeting) (msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.MeetingId)
	Param = append(Param, info.MeetingTitle)
	Param = append(Param, info.MeetingTime)
	Param = append(Param, info.MeetingCity)
	Param = append(Param, info.MeetingAddress)
	Param = append(Param, info.MeetingPeople)
	Param = append(Param, info.MeetingAudioFileUrl)
	Param = append(Param, info.CorpName)
	Param = append(Param, info.CorpCode)
	Param = append(Param, currentTime)
	Param = append(Param, info.Creater)
	Param = append(Param, info.Bz1)
	Param = append(Param, info.Bz2)
	Param = append(Param, info.Bz3)

	num, err := dbHandler.Insert(db_handler.InsertMeeting_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "创建会议发生错误!")
	}
	return err
}

//查看会议列表
func GetMeetingList(info *meeting.MeetingList_Param) (res meeting.MeetingList_Result, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	var Param []interface{}
	Param = append(Param, info.CorpCode)

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)

	selRes, err := dbHandler.SelectList(db_handler.GetMeetingList_sql, Param...)

	meetingList := []meeting.Meeting{}

	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&meetingList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "会议列表信息转换发生错误!")
	}

	meetingListCount := []meeting.Meeting{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetMeetingListCount_sql, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&meetingListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "会议列表分页信息转换发生错误!")
	}

	res.MeetingList = meetingList
	res.TotalCount = int64(len(meetingListCount))

	return res, err
}

//修改会议
func ModifyMeeting(info *meeting.Meeting) (res meeting.Meeting, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	//currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.MeetingAudioFileUrl)
	Param = append(Param, info.MeetingId)

	num, err := dbHandler.Update(db_handler.ModifyMeetingAudioFileUrl_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "修改会议录音地址发生错误!")
	} else {
		//查询会议，返回前台展示
		res1, err1 := dbHandler.SelectOne(db_handler.GetMeetingById_sql, info.MeetingId)
		if len(res1) > 0 && err1 == nil {
			decoder := ObtainDecoderConfig(&res)
			err := decoder.Decode(res1)
			common.ErrorHandler(err, "会议展示信息转换发生错误!")
		} else {
			err = err1
		}
	}

	return res, err
}
