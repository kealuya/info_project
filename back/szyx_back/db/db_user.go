package db

import (
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity"
)

func MessageLogin(phone string) (userInfo []entity.LoginUserInfo, msg error) {
	defer common.RecoverHandler(func(rcErr error) {
		userInfo = []entity.LoginUserInfo{}
		msg = rcErr
	})
	userInfos := []entity.LoginUserInfo{}
	dbHandler := db_handler.NewDbHandler()
	var param []interface{}
	param = append(param, phone)
	param = append(param, "1")
	var str = "where USER.MOBILE = ? AND b.serviceId = ? "
	dbAuthResult, err := dbHandler.SelectList(db_handler.Select_LoginUser+str, param...)
	if len(dbAuthResult) > 0 {
		decoder := ObtainDecoderConfig(&userInfos)
		err := decoder.Decode(dbAuthResult)
		common.ErrorHandler(err, "用户信息转换发生错误！")
	}
	common.ErrorHandler(err, "数据库操作异常")
	return userInfos, err
}
