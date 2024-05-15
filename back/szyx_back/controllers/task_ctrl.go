package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"szyx_back/entity/task"
	"szyx_back/models"
)

/**
任务模块
*/
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
	task_param := new(task.Task)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("任务发布入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &task_param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.CreateTask(task_param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "任务发布成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务发布失败::%s", err)
	}
}

// @Title 任务池列表
// @Tags GetTaskPoolList
// @Summary 任务池列表
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "Task struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getTaskPoolList [post]
func (TaskCtrl *TaskCtrl) GetTaskPoolList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	taskList_Param := new(task.TaskList_Param)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	paramerr := jsoniter.Unmarshal(jsonByte, &taskList_Param)
	logs.Info("查询任务池入参：" + string(jsonByte))
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	logs.Info("查询任务池列表入参：" + string(jsonByte))
	//业务处理
	res, err := models.GetTaskPoolList(taskList_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务查询失败::%s", err)
	}
}

// @Title 任务池 任务详情
// @Tags TaskPoolDetails
// @Summary 用户完成任务详情
// @accept application/json
// @Produce application/json
// @Param data body task.MyTask true "MyTask struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /taskPoolDetails [post]
func (TaskCtrl *TaskCtrl) TaskPoolDetails() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	myTask_Param := new(task.MyTask)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info(" 任务池详情入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &myTask_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.TaskPoolDetails(myTask_Param.TaskId)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "用户任务详情查询成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("用户任务详情查询失败::%s", err)
	}
}


// @Title 用户个人任务列表
// @Tags GetTaskList
// @Summary 用户任务列表
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
	myTaskList_Param := new(task.MyTaskList_Param)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("查询任务列表入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &myTaskList_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetTaskList(myTaskList_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "操作成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务查询失败::%s", err)
	}
}

// @Title  校验是否能参加任务，一个任务只能参加一次
// @Tags CheckIsJoinTask
// @Summary 校验是否能参加任务，一个任务只能参加一次
// @accept application/json
// @Produce application/json
// @Param data body task.MyTask true "MyTask struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /checkIsJoinTask [post]
func (TaskCtrl *TaskCtrl) CheckIsJoinTask() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	myTask_Param := new(task.MyTask)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("校验是否能参加任务入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &myTask_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "checkIsJoinTask接口入参有误!"
		return
	}
	//业务处理
	result,msg,err := models.CheckIsJoinTask(myTask_Param)
	if err == nil {
		resJson.Success = true
		resJson.Data = result
		resJson.Msg = msg
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("checkIsJoinTask接口调用失败::%s", err)
	}
}

// @Title 申请参加任务
// @Tags ApplyJoinTask
// @Summary 申请参加任务
// @accept application/json
// @Produce application/json
// @Param data body task.MyTask true "MyTask struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /applyJoinTask [post]
func (TaskCtrl *TaskCtrl) ApplyJoinTask() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	myTask_Param := new(task.MyTask)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("申请参加任务入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &myTask_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.ApplyJoinTask(myTask_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "参加成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务参加失败::%s", err)
	}
}

// @Title 完成我的任务
// @Tags FinishMyTask
// @Summary 完成我的任务
// @accept application/json
// @Produce application/json
// @Param data body task.MyTask true "MyTask struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /finishMyTask [post]
func (TaskCtrl *TaskCtrl) FinishMyTask() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	myTask_Param := new(task.MyTask)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("完成任务入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &myTask_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.FinishMyTask(myTask_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "任务完成成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("任务完成失败::%s", err)
	}
}

// @Title 用户任务详情
// @Tags MyTaskDetails
// @Summary 用户完成任务详情
// @accept application/json
// @Produce application/json
// @Param data body task.MyTask true "MyTask struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /myTaskDetails [post]
func (TaskCtrl *TaskCtrl) MyTaskDetails() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	myTask_Param := new(task.MyTask)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("用户任务详情入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &myTask_Param)
	if paramerr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.MyTaskDetails(myTask_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "用户任务详情查询成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("用户任务详情查询失败::%s", err)
	}
}

// @Title 用户放弃未完成的任务
// @Tags GiveUpTask
// @Summary 用户放弃未完成的任务
// @accept application/json
// @Produce application/json
// @Param data body task.Task true "Task struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @router /giveUpTask [post]
func (TaskCtrl *TaskCtrl)  GiveUpTask() {
	resJson := NewJsonStruct(nil)
	defer func() {
		TaskCtrl.Data["json"] = resJson
		TaskCtrl.ServeJSON()
	}()
	task_param := new(task.Task)
	var jsonByte = TaskCtrl.Ctx.Input.RequestBody
	logs.Info("放弃未完成的任务入参：" + string(jsonByte))
	paramerr := jsoniter.Unmarshal(jsonByte, &task_param)

	if paramerr != nil {
		resJson.Success = false


		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.GiveUpTask(task_param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "删除任务成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("放弃任务失败::%s", err)
	}
}