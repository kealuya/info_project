package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"szyx_back/entity"
	"szyx_back/models"
)

type UserCtrl struct {
	beego.Controller
}

// @Title 个人信息统计
// @Tags GetUserInfoCount
// @Summary 个人信息统计
// @accept application/json
// @Produce application/json
// @Param data body entity.UserInfo true "UserInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getUserInfoCount [post]
func (UserCtrl *UserCtrl) GetUserInfoCount() {
	resJson := NewJsonStruct(nil)
	defer func() {
		UserCtrl.Data["json"] = resJson
		UserCtrl.ServeJSON()
	}()
	userInfoCount_Param := new(entity.UserInfo)
	var jsonByte = UserCtrl.Ctx.Input.RequestBody
	logs.Info("用户统计信息入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &userInfoCount_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetUserInfoCount(userInfoCount_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "用户统计信息查询成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("用户统计信息查询失败::%s", err)
	}
}
