package task

import (
	"szyx_back/entity"
	"szyx_back/entity/meeting"
)

//============================================我的任务======================

//我的任务
type MyTask struct {
	//entity.Base
	TaskId      string            `json:"taskId" description:"任务ID" `
	TaskTitle   string            `json:"taskTitle" description:"任务标题" `
	TaskData    string            `json:"taskData" description:"任务介绍" `
	TaskContent string            `json:"taskContent" description:"任务内容" `
	TaskType    string            `json:"taskType" description:"任务类型" `
	TaskStatus  string            `json:"taskStatus" description:"任务状态" `
	TaskImg     string            `json:"taskImg" description:"任务图片" `
	UserId      string            `json:"userId" description:"用户ID" `
	UserName    string            `json:"userName" description:"用户姓名" `
	UserMobile  string            `json:"userMobile" description:"用户手机号" `
	Flag        string            `json:"flag" description:"完成状态" ` // 任务状态    0.待完成    1.已完成
	FinishTime  string            `json:"finishTime" description:"完成时间" `
	MeetingId   string            `json:"meetingId" description:"关联的会议ID，用于完成任务关联会议下的会议文件" `
	MeetingList []meeting.Meeting `json:"MeetingList" description:"完成任务，关联会议" `
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
type MyTaskList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
	UserId        string `json:"userId" description:"用户ID" `
	Flag          string `json:"flag" description:"完成状态" `
}

//获取任务列表-返回值
type MyTaskList_Result struct {
	entity.Paging          //分页信息
	MyTaskList    []MyTask `json:"myTaskList" description:"我的任务list" `
}
