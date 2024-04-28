package worth

import "szyx_back/entity"

//价值
type Worth struct {
	WorthId      string `json:"worthId" description:"价值Id" `
	WorthTitle   string `json:"worthTitle" description:"价值title" `
	WorthScore   string `json:"worthScore" description:"任务价值评分" `
	WorthData    string `json:"worthData" description:"价值概况" `
	WorthContent string `json:"worthContent" description:"价值内容" `
	ApplyTime    string `json:"applyTime" description:"价值申请时间" `
	Status       string `json:"status" description:"状态" ` //状态 价值状态 0：未申请   1：已申请
	Money        string `json:"money" description:"价值申请金额" `
	UserId       string `json:"userId" description:"用户ID" `
	UserName     string `json:"userName" description:"用户姓名" `
	UserMobile   string `json:"userMobile" description:"用户手机号" `
	CorpName     string `json:"corpName" description:"企业名称" `
	CorpCode     string `json:"corpCode" description:"企业code"`
	CreateTime   string `json:"createTime" description:"创建时间"`
	Creater      string `json:"creater" description:"创建人"`
	Bz1          string `json:"bz1" description:"备注1"`
	Bz2          string `json:"bz2" description:"备注2"`
	Bz3          string `json:"bz3" description:"备注3"`
}

//获取价值列表-入参
type WorthList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
	Status        string `json:"status" description:"状态" `
	UserId        string `json:"userId" description:"用户ID" `
}

//获取价值列表-返回值
type WorthList_Result struct {
	entity.Paging         //分页信息
	WorthList     []Worth `json:"worthList" description:"价值list" `
}
