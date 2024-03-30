package meeting

import "szyx_back/entity"

//会议
type Meeting struct {
	entity.Base
	MeetingId              string `json:"meetingId" description:"会议id" `
	MeetingTitle           string `json:"meetingTitle" description:"会议标题" `
	MeetingTime            string `json:"meetingTime" description:"会议时间  年-月-日  时：分：秒" `
	MeetingCity            string `json:"meetingCity" description:"会议城市" `
	MeetingAddress         string `json:"meetingAddress" description:"会议地址" `
	MeetingPeople          string `json:"meetingPeople" description:"参会人" `
	MeetingAudioFileUrl    string `json:"meetingAudioFileUrl" description:"音频文件地址'；'分隔" `
	MeetingMminutesFileUrl string `json:"meetingMminutesFileUrl" description:"会议纪要地址" `
	MeetingBrainMapFileUrl string `json:"meetingBrainMapFileUrl" description:"会议脑图地址" `
}

//获取会议列表-入参
type MeetingList_Param struct {
	entity.Paging        //分页信息
	SearchKey     string `json:"searchKey" description:"搜索关键字" `
	CorpCode      string `json:"corpCode" description:"企业ID" `
}

//会议列表-返回值
type MeetingList_Result struct {
	entity.Paging           //分页信息
	MeetingList   []Meeting //会议list
}
