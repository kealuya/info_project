package models

import (
	"szyx_back/db"
	. "szyx_back/entity"
)

//短信验证码登录
func MessageLogin(mobile string) (loginResp LoginResponseData, err error) {
	userInfo, err := db.MessageLogin(mobile)
	if len(userInfo) > 0 {
		loginResp.UserInfo = userInfo[0]
		loginResp.Result = true
	} else {
		loginResp.Msg = "未找到该用户，请重新登陆!"
		loginResp.Result = false

	}
	return loginResp, err
}

//用户信息统计 任务数  会议数  价值数
func GetUserInfoCount(user *UserInfo) (userInfoCount UserInfoCount_Result, err error) {
	userInfoCount, err = db.GetUserInfoCount(user)
	return userInfoCount, err
}
