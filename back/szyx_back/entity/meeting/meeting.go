package meeting

//base struct
type Base struct {
	CorpName   string `json:"corpName" description:"企业名称" `
	CorpCode   string `json:"corpCode" description:"企业code"`
	CreateTime string `json:"createTime" description:"创建时间"`
	Creater    string `json:"creater" description:"创建人"`
	Bz1        string `json:"bz1" description:"备注1"`
	Bz2        string `json:"bz2" description:"备注2"`
	Bz3        string `json:"bz3" description:"备注3"`
}

//会议
type Meeting struct {
	Base
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
