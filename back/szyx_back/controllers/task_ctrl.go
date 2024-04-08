package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"szyx_back/common"
	"szyx_back/entity/task"
	"szyx_back/models"
)

type TaskCtrl struct {
	beego.Controller
}

// @Title 任务发布
// @Tags CreateTask
// @Summary 任务发布
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "Task struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发布成功"}"
// @router /createTask [post]
func (TaskCtrl *TaskCtrl) CreateTask() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	task := new(task.Task)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &task)
	logs.Info("任务发布入参：" + string(jsonByte))
	//业务处理
	err := models.CreateTask(task)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "任务发布成功"
		//resJson.Data = meetingRes
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务发布失败::%s", err)
	}
}

// @Title 任务列表
// @Tags GetTaskList
// @Summary 任务列表
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "Task struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getTaskList [post]
func (TaskCtrl *TaskCtrl) GetTaskList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	task_param := new(task.Task)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &task_param)
	logs.Info("查询任务列表入参：" + string(jsonByte))
	//业务处理
	res, err := models.GetTaskList(task_param)
	if err == nil {
		resJson.Success = true
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务查询失败::%s", err)
	}
}
