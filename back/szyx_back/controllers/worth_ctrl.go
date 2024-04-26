package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"szyx_back/entity/worth"
	"szyx_back/models"
)

/**
价值模块
*/
type WorthCtrl struct {
	beego.Controller
}

// @Title 导出申请的价值Excel提供数据js生成excel
// @Tags ExportWorthExcel
// @Summary 导出Excel
// @accept application/json
// @Produce application/json
// @Param data body worth.Worth true "申请的价值"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @router /exportWorthExcel [post]
func (WorthCtrl *WorthCtrl) ExportExcel() {
	resJson := NewJsonStruct(nil)
	defer func() {
		WorthCtrl.Data["json"] = resJson
		WorthCtrl.ServeJSON()
	}()

	worth_param := new(worth.Worth)
	var jsonByte = WorthCtrl.Ctx.Input.RequestBody
	logs.Info("导出申请的价值入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &worth_param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	list, err := models.ExportWorthExcel(worth_param)

	if err == nil {
		resJson.Success = true
		resJson.Data = list
		resJson.Msg = "导出成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("导出excel表失败:%v", err)
	}
}

// @Title 申请价值
// @Tags ApplyWorth
// @Summary 申请价值
// @accept application/json
// @Produce application/json
// @Param data body worth.Worth true "Worth struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"价值申请成功"}"
// @router /applyWorth [post]
func (WorthCtrl *WorthCtrl) ApplyWorth() {
	resJson := NewJsonStruct(nil)
	defer func() {
		WorthCtrl.Data["json"] = resJson
		WorthCtrl.ServeJSON()
	}()
	worth_param := new(worth.Worth)
	var jsonByte = WorthCtrl.Ctx.Input.RequestBody
	logs.Info("申请价值入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &worth_param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.ApplyWorth(worth_param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "申请价值成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("申请价值失败::%v", err)
	}
}

// @Title 价值列表
// @Tags GetWorthList
// @Summary 价值列表
// @accept application/json
// @Produce application/json
// @Param data body worth.Worth true "Worth struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getWorthList [post]
func (WorthCtrl *WorthCtrl) GetWorthList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		WorthCtrl.Data["json"] = resJson
		WorthCtrl.ServeJSON()
	}()
	worthList_Param := new(worth.WorthList_Param)
	var jsonByte = WorthCtrl.Ctx.Input.RequestBody
	logs.Info("价值列表入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &worthList_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	worthList, err := models.GetWorthList(worthList_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "价值列表获取成功"
		resJson.Data = worthList
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("价值列表获取失败::%s", err)
	}
}

// @Title 价值详情
// @Tags GetWorthDetails
// @Summary 价值详情
// @accept application/json
// @Produce application/json
// @Param data body worth.Worth true "Worth struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getWorthDetails [post]
func (WorthCtrl *WorthCtrl) GetWorthDetails() {
	resJson := NewJsonStruct(nil)
	defer func() {
		WorthCtrl.Data["json"] = resJson
		WorthCtrl.ServeJSON()
	}()
	worth_Param := new(worth.Worth)
	var jsonByte = WorthCtrl.Ctx.Input.RequestBody
	logs.Info("价值详情入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &worth_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	worthList, err := models.GetWorthDetails(worth_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "价值详情获取成功"
		resJson.Data = worthList
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("价值详情获取失败::%s", err)
	}
}
