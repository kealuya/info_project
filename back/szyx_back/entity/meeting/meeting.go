package meeting

import "szyx_back/entity"

//会议
type Meeting struct {
	MeetingId              string        `json:"meetingId" description:"会议id" `
	MeetingType            string        `json:"meetingType" description:"会议类型" ` //audio：音频会议    document：文档会议
	MeetingTitle           string        `json:"meetingTitle" description:"会议标题" `
	MeetingTime            string        `json:"meetingTime" description:"会议时间  年-月-日  时：分：秒" `
	MeetingCity            string        `json:"meetingCity" description:"会议城市" `
	MeetingAddress         string        `json:"meetingAddress" description:"会议地址" `
	MeetingPeople          string        `json:"meetingPeople" description:"参会人" `
	MeetingFileUrl         string        `json:"meetingFileUrl" description:"会议文件地址" `         //完成任务是选择会议，添加会议文件。手动上传用
	TaskId                 string        `json:"taskId" description:"任务id，任务与会议为 1对多的关系，故而在此添加关联字段" `
	MeetingFlag            string        `json:"meetingFlag" description:"会议是否使用" ` //0：未使用   1：已使用   完成任务需要选择会议， 要区分会议是否被使用
	MeetingFile            []MeetingFile `json:"meetingFile" description:"会议文件list" `
	CorpName               string        `json:"corpName" description:"企业名称" `
	CorpCode               string        `json:"corpCode" description:"企业code"`
	CreateTime             string        `json:"createTime" description:"创建时间、结束时间"`
	Creater                string        `json:"creater" description:"创建人"`
	Bz1                    string        `json:"bz1" description:"备注1"`
	Bz2                    string        `json:"bz2" description:"备注2"`
	Bz3                    string        `json:"bz3" description:"备注3"`
	MeetingAudioText       string        `json:"meetingAudioText" description:"会议录音转译的文本内容"`    //通过关联关系获取
	MeetingSummary         string        `json:"meetingSummary" description:"会议摘要"`    //通过关联关系获取
	MeetingMinutes         string        `json:"meetingMinutes" description:"会议纪要"`    //通过关联关系获取
	MeetingBrainMap        string         `json:"meetingBrainMap" description:"会议脑图字符"` //通过关联关系获取
	TranslationState        string         `json:"translationState" description:"录音转译状态 0：进行中；1：完成；2：错误"` //通过关联关系获取
	MeetingSummaryState        string         `json:"meetingSummaryState" description:"摘要生成状态  0：进行中；1：完成；2：错误"` //通过关联关系获取

}

//会议文件
type MeetingFile struct {
	Id           string `json:"id" description:"id,主键" `
	MeetingId    string `json:"meetingId" description:"关联的会议id" `
	FileType     string `json:"fileType" description:"文件类型" ` //文件类型   音频：mp3   文档：word  脑图：xmind
	FileName     string `json:"fileName" description:"文件名称" `
	FileUrl      string `json:"fileUrl" description:"文件云地址" `
	AudioTime    string `json:"audioTime" description:"音频时长  只在fileType值为mp3 有作用" `
	CorpName     string `json:"corpName" description:"企业名称" `
	CorpCode     string `json:"corpCode" description:"企业code"`
	CreateTime   string `json:"createTime" description:"创建时间"`
	Creater      string `json:"creater" description:"创建人"`
}

//获取会议列表-入参
type MeetingList_Param struct {
	entity.Paging        //分页信息
	CorpCode      string `json:"corpCode" description:"企业ID" `
	UserId        string `json:"userId" description:"用户id" `
	MeetingType   string `json:"meetingType" description:"会议类型" `
	StartTime     string `json:"startTime" description:"会议时间范围开始时间" `
	EndTime       string `json:"endTime" description:"会议时间范围结束时间" `
	MeetingFlag   string `json:"meetingFlag" description:"会议是否使用 0：未使用   1：已使用" `
}

//会议列表-返回值
type MeetingList_Result struct {
	entity.Paging           //分页信息
	MeetingList   []Meeting //会议list
}

//文件上传返回值
type MeetingFile_Result struct {
	FileUrl string `json:"fileUrl" description:"上传后的地址" `
}
