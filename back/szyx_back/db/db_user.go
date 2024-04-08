package db

import (
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
