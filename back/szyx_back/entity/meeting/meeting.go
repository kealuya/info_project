package meeting

import "szyx_back/entity"

//会议
type Meeting struct {
	MeetingId              string `json:"meetingId" description:"会议id" `
	MeetingType            string `json:"meetingType" description:"会议类型" ` //audio：音频会议    document：文档会议
	MeetingTitle           string `json:"meetingTitle" description:"会议标题" `
	MeetingTime            string `json:"meetingTime" description:"会议时间  年-月-日  时：分：秒" `
	MeetingCity            string `json:"meetingCity" description:"会议城市" `
	MeetingAddress         string `json:"meetingAddress" description:"会议地址" `
	MeetingPeople          string `json:"meetingPeople" description:"参会人" `
	MeetingAudioFileUrl    string `json:"meetingAudioFileUrl" description:"音频文件地址'；'分隔" ` //完成任务是选择文件，不是选择会议，分表处理，存到speech中
	MeetingMminutesFileUrl string `json:"meetingMminutesFileUrl" description:"会议纪要地址" `   //完成任务是选择文件，不是选择会议，分表处理，存到speech中
	MeetingBrainMapFileUrl string `json:"meetingBrainMapFileUrl" description:"会议脑图地址" `   //完成任务是选择文件，不是选择会议，分表处理，存到file中
	TaskId                 string `json:"taskId" description:"任务id，任务与会议为 1对多的关系，故而在此添加关联字段" `
	MeetingFlag            string `json:"meetingFlag" description:"会议是否使用" ` //0：未使用   1：已使用   完成任务需要选择会议， 要区分会议是否被使用
	MeetingAudioName       string `json:"meetingAudioName" description:"音频文件总名称" `   //创建会议 生成的文件总名称  用于选择会议前端展示该字段
	MeetingMminutesName    string `json:"meetingMminutesName" description:"文档文件总名称" `//创建会议 生成的文件总名称  用于选择会议前端展示该字段
	MeetingBrainMapName    string `json:"meetingBrainMapName" description:"脑图文件总名称" `//创建会议 生成的文件总名称  用于选择会议前端展示该字段
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
	CorpCode      string `json:"corpCode" description:"企业ID" `
	UserId        string `json:"userId" description:"用户id" `
	MeetingType   string `json:"meetingType" description:"会议类型" `
	StartTime     string `json:"startTime" description:"会议时间范围开始时间" `
	EndTime		  string `json:"endTime" description:"会议时间范围结束时间" `
	MeetingFlag   string `json:"meetingFlag" description:"会议是否使用 0：未使用   1：已使用" `
}

//会议列表-返回值
type MeetingList_Result struct {
	entity.Paging           //分页信息
	MeetingList   []Meeting //会议list
}
