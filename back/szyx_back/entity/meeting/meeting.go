package meeting

import "szyx_back/entity"

//会议
type Meeting struct {
	//entity.Base
	MeetingId              string `json:"meetingId" description:"会议id" `
	MeetingTitle           string `json:"meetingTitle" description:"会议标题" `
	MeetingTime            string `json:"meetingTime" description:"会议时间  年-月-日  时：分：秒" `
	MeetingCity            string `json:"meetingCity" description:"会议城市" `
	MeetingAddress         string `json:"meetingAddress" description:"会议地址" `
	MeetingPeople          string `json:"meetingPeople" description:"参会人" `
	MeetingAudioFileUrl    string `json:"meetingAudioFileUrl" description:"音频文件地址'；'分隔" ` //完成任务是选择文件，不是选择会议，分表处理，存到speech中
	MeetingMminutesFileUrl string `json:"meetingMminutesFileUrl" description:"会议纪要地址" `   //完成任务是选择文件，不是选择会议，分表处理，存到speech中
	MeetingBrainMapFileUrl string `json:"meetingBrainMapFileUrl" description:"会议脑图地址" `   //完成任务是选择文件，不是选择会议，分表处理，存到file中
	TaskId                 string `json:"taskId" description:"任务id，任务与会议为 1对多的关系，故而在此添加关联字段" `
	//FIXME base 字段，上面的写法，在查询db后，不能赋值
	CorpName   string `json:"corpName" description:"企业名称" `
	CorpCode   string `json:"corpCode" description:"企业code"`
	CreateTime string `json:"createTime" description:"创建时间"`
	Creater    string `json:"creater" description:"创建人"`
	Bz1        string `json:"bz1" description:"备注1"`
	Bz2        string `json:"bz2" description:"备注2"`
	Bz3        string `json:"bz3" description:"备注3"`
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
