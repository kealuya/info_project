package db

import (
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity/meeting"
	"time"
)

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

	num, err := dbHandler.Insert(db_handler.InsertMeeting, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "创建会议发生错误!")
	}
	return err
}
