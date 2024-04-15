package task

import "szyx_back/entity"

//任务池
type Task struct {
	//entity.Base
	TaskId      string `json:"taskId" description:"任务ID" `
	TaskTitle   string `json:"taskTitle" description:"任务标题" `
	TaskTarget  string `json:"taskTarget" description:"任务目标" `
	TaskContent string `json:"taskContent" description:"任务内容" `
	TaskType    string `json:"taskType" description:"任务类型" `
	TaskStatus  string `json:"taskStatus" description:"任务状态" `
	TaskImg     string `json:"taskImg" description:"任务图片" `
	//FIXME base 字段，上面的写法，在查询db后，不能赋值
	CorpName   string `json:"corpName" description:"企业名称" `
	CorpCode   string `json:"corpCode" description:"企业code"`
	CreateTime string `json:"createTime" description:"创建时间"`
	Creater    string `json:"creater" description:"创建人"`
	Bz1        string `json:"bz1" description:"备注1"`
	Bz2        string `json:"bz2" description:"备注2"`
	Bz3        string `json:"bz3" description:"备注3"`
}

//获取任务列表-入参
type TaskList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
	Status        string `json:"status" description:"状态" `
	UserId        string `json:"userId" description:"用户ID" `
	Flag          string `json:"flag" description:"完成状态" `
}

//获取任务列表-返回值
type TaskList_Result struct {
	entity.Paging        //分页信息
	TaskList      []Task `json:"taskList" description:"任务list" `
}
