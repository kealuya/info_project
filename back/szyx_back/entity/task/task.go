package task

import "szyx_back/entity"

//任务池
type Task struct {
	entity.Base
	TaskId      string `json:"taskId" description:"任务ID" `
	TaskTitle   string `json:"taskTitle" description:"任务标题" `
	TaskTarget  string `json:"taskTarget" description:"任务目标" `
	TaskContent string `json:"taskContent" description:"任务内容" `
	TaskType    string `json:"taskType" description:"任务类型" `
	TaskStatus  string `json:"taskStatus" description:"任务状态" `
	TaskImg     string `json:"taskImg" description:"任务图片" `
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

//我的任务
type MyTask struct {
	Task              //任务池字段
	UserId     string `json:"userId" description:"用户ID" `
	UserName   string `json:"userName" description:"用户姓名" `
	UserMobile string `json:"userMobile" description:"用户手机号" `
	Flag       string `json:"flag" description:"完成状态" `
	FinishTime string `json:"finishTime" description:"完成时间" `
	TaskData   string `json:"taskData" description:"资料id集合；分割" `
}
