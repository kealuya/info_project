package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"szyx_back/entity/evidenceChain"
	"szyx_back/models"
)

type EvidenceChainCtrl struct {
	beego.Controller
}

// @Title 生成证据链模块 - 获取计划下的所有会议，生成会议主题
// @Tags CreateMeetingEvidenceChain
// @Summary 生成证据链模块 - 获取计划下的所有会议，生成会议主题
// @accept application/json
// @Produce application/json
// @Param data body evidenceChain.EvidenceChain_Param true "EvidenceChain_Param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /createMeetingEvidenceChain [post]
func (EvidenceChainCtrl *EvidenceChainCtrl) CreateMeetingEvidenceChain() {
	resJson := NewJsonStruct(nil)
	defer func() {
		EvidenceChainCtrl.Data["json"] = resJson
		EvidenceChainCtrl.ServeJSON()
	}()
	params := new(evidenceChain.EvidenceChain_Param)
	var jsonByte = EvidenceChainCtrl.Ctx.Input.RequestBody
	logs.Info("获取计划下的所有会议，生成会议主题的入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &params)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateMeetingEvidenceChain(params)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "获取计划下的所有会议，生成会议主题数据获取成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("查询失败:", err)
	}
}

// @Title 生成证据链模块 - 获取计划下的新闻 生成截图
// @Tags CreateNewsEvidenceChain
// @Summary 生成证据链模块 - 获取计划下的新闻 生成截图
// @accept application/json
// @Produce application/json
// @Param data body evidenceChain.EvidenceChain_Param true "EvidenceChain struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /createNewsEvidenceChain [post]
func (EvidenceChainCtrl *EvidenceChainCtrl) CreateNewsEvidenceChain() {
	resJson := NewJsonStruct(nil)
	defer func() {
		EvidenceChainCtrl.Data["json"] = resJson
		EvidenceChainCtrl.ServeJSON()
	}()
	params := new(evidenceChain.EvidenceChain_Param)
	var jsonByte = EvidenceChainCtrl.Ctx.Input.RequestBody
	logs.Info("获取计划下的新闻 生成截图的入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &params)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateNewsEvidenceChain(params)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "获取计划下的新闻成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("查询失败:", err)
	}
}

// @Title 生成证据链模块 - 根据时间范围、需要截图的张数 获取智能问答id
// @Tags CreateWDEvidenceChain
// @Summary 生成证据链模块 - 根据时间范围、需要截图的张数 获取智能问答id
// @accept application/json
// @Produce application/json
// @Param data body evidenceChain.EvidenceChainWD_Param true "EvidenceChain struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /createWDEvidenceChain [post]
func (EvidenceChainCtrl *EvidenceChainCtrl) CreateWDEvidenceChain() {
	resJson := NewJsonStruct(nil)
	defer func() {
		EvidenceChainCtrl.Data["json"] = resJson
		EvidenceChainCtrl.ServeJSON()
	}()
	params := new(evidenceChain.EvidenceChainWD_Param)
	var jsonByte = EvidenceChainCtrl.Ctx.Input.RequestBody
	logs.Info("根据时间范围、需要截图的张数 获取智能问答id的入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &params)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateWDEvidenceChain(params)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "获取问答id列表成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("查询失败:", err)
	}
}

// @Title 生成证据链模块 - 获取计划下的问卷回答id,生成问卷回答截图
// @Tags CreateQuestionAnswerEvidenceChain
// @Summary 生成证据链模块 - 获取计划下的问卷回答id,生成问卷回答截图
// @accept application/json
// @Produce application/json
// @Param data body evidenceChain.EvidenceChain_Param true "EvidenceChain_Param struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /createQuestionAnswerEvidenceChain [post]
func (EvidenceChainCtrl *EvidenceChainCtrl) CreateQuestionAnswerEvidenceChain() {
	resJson := NewJsonStruct(nil)
	defer func() {
		EvidenceChainCtrl.Data["json"] = resJson
		EvidenceChainCtrl.ServeJSON()
	}()
	params := new(evidenceChain.EvidenceChain_Param)
	var jsonByte = EvidenceChainCtrl.Ctx.Input.RequestBody
	logs.Info("获取计划下的问卷回答id的入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &params)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.CreateQuestionAnswerEvidenceChain(params)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "获取计划下的问卷回答id数据获取成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("查询失败:", err)
	}
}
