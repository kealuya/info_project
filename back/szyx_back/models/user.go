package models

import (
	"fmt"
	. "szyx_back/entity"
)

//用户登录
func UserLogin(user *LoginRequestPar) (UserInfo, error) {
	fmt.Println(user)
	return UserInfo{Corp: "szht", EmpCode: "18892206664", Username: "测试1", Tel: "13612332123", IdCard: "120103199010102345", Dept: "信息学院"}, nil
}

//修改用户密码
func SetUserInfo(user *UserInfo) (UserInfo, error) {
	fmt.Println(user)
	return UserInfo{Corp: "szht", EmpCode: "18892206664", Username: "测试1", Tel: "13612332123", IdCard: "120103199010102345", Dept: "信息学院"}, nil
}
