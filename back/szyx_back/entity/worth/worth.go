package worth

import "szyx_back/entity"

//价值
type Worth struct {
	entity.Base
	WorthTitle string `json:"worthTitle" description:"价值title" `
	WorthScore string `json:"worthScore" description:"价值评分" `
	Status     string `json:"status" description:"状态" `
	Money      string `json:"money" description:"钱数" `
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
