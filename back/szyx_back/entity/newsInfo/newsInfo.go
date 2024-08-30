package newsInfo

import "szyx_back/entity"

// 新闻媒体列表  入参
type NewsList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
}

// 新闻媒体列表-返回值
type NewsList_Result struct {
	entity.Paging            //分页信息
	NewsInfoList  []NewsInfo `json:"newsInfoList" description:"新闻媒体列表list" `
}

// 新闻推广信息
type NewsInfo struct {
	Id            string `json:"id" description:"id,主键" `
	NewsId        string `json:"newsId" description:"新闻id" `
	NewsName      string `json:"newsName" description:"新闻标题" `
	NewsText      string `json:"newsText" description:"新闻内容" `
	CorpName      string `json:"corpName" description:"企业名称" `
	CorpCode      string `json:"corpCode" description:"企业code"`
	Creater       string `json:"creater" description:"创建人"`
	CreaterName   string `json:"createrName" description:"创建人Name"`
	CreateTime    string `json:"createTime" description:"创建时间"`
	ModifyTime    string `json:"modifyTime" description:"修改时间"`
	ModifyId      string `json:"modifyId" description:"修改人"`
	Corp          string `json:"corp" description:"企业code-分享专用"`
	CreateTimeNew string `json:"createTimeNew" description:"管理员生成的创建时间"` //为前端分享链接后 取不到封装的CorpCode字段使用
}

// 超级管理员、用户 - 新增推广新闻
type NewsInfoInUser struct {
	NewsName   string `json:"newsName" description:"新闻标题" `
	NewsText   string `json:"newsText" description:"新闻内容" `
	PlanId     string `json:"planId" description:"计划id"`
	PlanName   string `json:"planName" description:"计划Name"`
	CorpName   string `json:"corpName" description:"企业名称" `
	CorpCode   string `json:"corpCode" description:"企业code"`
	Creater    string `json:"creater" description:"创建人"`
	CreateTime string `json:"createTime" description:"创建时间"`
}

// 超级管理员、用户 - 新增视频
type VideoInfoInUser struct {
	NewsVideoName string `json:"newsVideoName" description:"视频媒体标题" `
	NewsVideoUrl  string `json:"newsVideoUrl" description:"视频媒体url地址" `
	PlanId        string `json:"planId" description:"计划id"`
	PlanName      string `json:"planName" description:"计划Name"`
	CorpName      string `json:"corpName" description:"企业名称" `
	CorpCode      string `json:"corpCode" description:"企业code"`
	Creater       string `json:"creater" description:"创建人"`
	CreateTime    string `json:"createTime" description:"创建时间"`
}

// 视频媒体列表  入参
type VideoList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
}

// 视频媒体列表-返回值
type VideoList_Result struct {
	entity.Paging             //分页信息
	VideoInfoList []VideoInfo `json:"videoInfoList" description:"视频媒体列表list" `
}

// 视频推广信息
type VideoInfo struct {
	Id            string `json:"id" description:"id,主键" `
	NewsVideoId   string `json:"newsVideoId" description:"视频媒体id" `
	NewsVideoName string `json:"newsVideoName" description:"视频媒体标题" `
	NewsVideoUrl  string `json:"newsVideoUrl" description:"视频媒体播放地址" `
	CorpName      string `json:"corpName" description:"企业名称" `
	CorpCode      string `json:"corpCode" description:"企业code"`
	Creater       string `json:"creater" description:"创建人"`
	CreateTime    string `json:"createTime" description:"创建时间"`
	IsDelete      string `json:"isDelete" description:"是否删除"` //0 存在  1删除
	ModifyTime    string `json:"modifyTime" description:"修改时间"`
	ModifyId      string `json:"modifyId" description:"修改人"`
	CreaterName   string `json:"createrName" description:"创建人Name"`
	CreateTimeNew string `json:"createTimeNew" description:"管理员生成的创建时间"`
}
