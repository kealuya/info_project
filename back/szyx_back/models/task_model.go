package models

import (
	"szyx_back/db"
	"szyx_back/entity/task"
)

/**
任务发布
*/
func CreateTask(task *task.Task) (err error) {
	return err
}

/**
获取任务池列表
*/
func GetTaskPoolList(taskList_Param *task.TaskList_Param) (res task.TaskList_Result, err error) {
	res, err = db.GetTaskPoolList(taskList_Param)
	return res, err
}

/**
获取我的任务列表
*/
func GetTaskList(myTaskList_Param *task.MyTaskList_Param) (res task.MyTaskList_Result, err error) {
	res, err = db.GetTaskList(myTaskList_Param)
	return res, err
}

/**
参加任务
*/
func ApplyJoinTask(myTask_Param *task.MyTask) (err error) {
	err = db.ApplyJoinTask(myTask_Param)
	return err
}

/**
校验是否能参加任务，一个任务只能参加一次
*/
func CheckIsJoinTask(myTask_Param *task.MyTask) (result string,err error) {
	result,err = db.CheckIsJoinTask(myTask_Param)
	return result,err
}

/**
完成任务，生成价值数据
*/
func FinishMyTask(myTask_Param *task.MyTask) (err error) {
	//完成任务
	err = db.FinishMyTask(myTask_Param)
	return err
}

/**
查询用户任务详情
*/
func MyTaskDetails(myTask_Param *task.MyTask) (res task.MyTask, err error) {
	res, err = db.MyTaskDetails(myTask_Param)
	return res, err
}
