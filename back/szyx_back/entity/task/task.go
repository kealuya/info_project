package task

import "szyx_back/entity"

//任务
type Task struct {
	entity.Base
	TaskName string `json:"TaskName" description:"任务名称" `
}

//获取任务列表-入参
type TaskList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
	Status        string `json:"status" description:"状态" `
	UserId        string `json:"userId" description:"用户ID" `
}

//获取任务列表-返回值
type TaskList_Result struct {
	entity.Paging        //分页信息
	TaskList      []Task `json:"taskList" description:"任务list" `
}
