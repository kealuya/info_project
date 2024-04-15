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
完成任务，生成价值数据
*/
func FinishMyTask(myTask_Param *task.MyTask) (err error) {
	err = db.FinishMyTask(myTask_Param)
	return err
}
