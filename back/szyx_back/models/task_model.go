package models

import (
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
func GetTaskList(taskList_Param *task.TaskList_Param) (res task.Task, err error) {

	return res, err
}
