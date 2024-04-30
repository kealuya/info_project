package db

import (
	"fmt"
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity"
)

func MessageLogin(phone string) (userInfo []entity.UserInfo, msg error) {
	defer common.RecoverHandler(func(rcErr error) {
		userInfo = []entity.UserInfo{}
		msg = rcErr
	})
	userInfos := []entity.UserInfo{}
	dbHandler := db_handler.NewDbHandler()
	var param []interface{}
	param = append(param, phone)
	var str = " where userMobile = ? "
	dbAuthResult, err := dbHandler.SelectList(db_handler.Select_LoginUser+str, param...)
	if len(dbAuthResult) > 0 {
		decoder := ObtainDecoderConfig(&userInfos)
		err := decoder.Decode(dbAuthResult)
		common.ErrorHandler(err, "用户信息转换发生错误！")
	}
	common.ErrorHandler(err, "数据库操作异常")
	return userInfos, err
}

//用户信息计数
func GetUserInfoCount(info *entity.UserInfo) (res entity.UserInfoCount_Result, msg error) {

	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}
	Param = append(Param, info.UserId)
	Param = append(Param, info.CorpCode)

	//会议条数
	meetingSelRes, err := dbHandler.SelectOne(db_handler.Select_userMeetingCount, Param...)
	//任务条数
	taskSelRes, err := dbHandler.SelectOne(db_handler.Select_userTaskCount, Param...)
	var Param2 []interface{}
	Param2 = append(Param2, info.UserId)
	Param2 = append(Param2, info.CorpCode)
	Param2 = append(Param2, common.MY_WORTH_APPLY_FLAG_KEY_1)

	//价值金额
	worthSelRes, err := dbHandler.SelectOne(db_handler.Select_userWorthCount, Param2...)

	fmt.Println(meetingSelRes["meetingCount"])
	fmt.Println(taskSelRes["taskCount"])
	fmt.Println(worthSelRes["worthCount"])
	//类型转换
	meetingCountInt64, _ := meetingSelRes["meetingCount"].(int64)
	taskCountInt64, _ := taskSelRes["taskCount"].(int64)
	worthCountFloat64, _ := worthSelRes["worthCount"].(float64)
	res.MeetingCount = meetingCountInt64
	res.TaskCount = taskCountInt64
	res.WorthCount = worthCountFloat64
	return res, err
}
