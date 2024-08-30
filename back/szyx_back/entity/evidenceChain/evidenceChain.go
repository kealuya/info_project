package evidenceChain

// 获取计划下生成证据的 -入参
type EvidenceChain_Param struct {
	PlanId   string `json:"planId" description:"计划id" `
	CorpCode string `json:"corpCode" description:"企业code"`
}

// 会议证据 返回参数
type MeetingEvidenceChain_Result struct {
	MeetingIdList []MeetingIdList `json:"meetingIdList" description:"会议id-List" `
}
type MeetingIdList struct {
	MeetingId string `json:"meetingId" description:"会议id" `
}

// 新闻推广证据 返回参数
type NewsEvidenceChain_Result struct {
	NewsIdList []NewsIdList `json:"newsIdList" description:"新闻id-List" `
}
type NewsIdList struct {
	NewsId string `json:"newsId" description:"新闻id" `
}

// 获取智能问答证据 -入参
type EvidenceChainWD_Param struct {
	StartDte string `json:"startDate" description:"开始时间" `
	EndDate  string `json:"endDate" description:"结束时间" `
	ImgNum   int64  `json:"imgNum" description:"需要截图的张数" `
	UserId   string `json:"userId" description:"提问人ID"`
	CorpCode string `json:"corpCode" description:"企业code"`
}

// 智能问答证据 返回参数
type EvidenceChainWD_Result struct {
	ParentIdList []ParentIdList `json:"parentIdList" description:"问答id-List" `
}
type ParentIdList struct {
	ParentId string `json:"parentId" description:"问答id" `
}

// 问卷回答截图证据 返回参数
type QuestionNaireEvidenceChain_Result struct {
	QuestionNaireList []QuestionNaireList `json:"questionNaireList" description:"问卷回答id-List" `
}
type QuestionNaireList struct {
	QuestionNaireId string `json:"questionNaireId" description:"问卷回答id" `
}
