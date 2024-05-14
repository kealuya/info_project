package db

import (
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity/kdxf"
)

//根据MeetingId查询orderId
func GetOrderIdByMeetingId(meetingId string) (res kdxf.Kdxf_speech, msg error) {
	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}
	Param = append(Param, meetingId)
	selRes, err := dbHandler.SelectOne(db_handler.SelectSpeechById_sql, Param...)
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&res)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "kdxf录音文件转换发生错误!")
	}
	return res, err
}
