package models

import (
	"szyx_back/db"
	"szyx_back/entity/evidenceChain"
)

// 生成证据链模块 - 获取计划下的所有会议，生成会议主题
func CreateMeetingEvidenceChain(meetingevidencechain_param *evidenceChain.EvidenceChain_Param) (meetingIdList evidenceChain.MeetingEvidenceChain_Result, err error) {
	meetingIdList, err = db.CreateMeetingEvidenceChain(meetingevidencechain_param)
	return meetingIdList, err
}

// 生成证据链模块 - 获取计划下的新闻 生成截图
func CreateNewsEvidenceChain(newsevidencechain_param *evidenceChain.EvidenceChain_Param) (newsIdList evidenceChain.NewsEvidenceChain_Result, err error) {
	newsIdList, err = db.CreateNewsEvidenceChain(newsevidencechain_param)
	return newsIdList, err
}

// 生成证据链模块 - 根据时间范围、需要截图的张数 获取智能问答id
func CreateWDEvidenceChain(info *evidenceChain.EvidenceChainWD_Param) (resList evidenceChain.EvidenceChainWD_Result, err error) {
	resList, err = db.CreateWDEvidenceChain(info)
	return resList, err
}

// 生成证据链模块 - 获取计划下的问卷回答id,生成问卷回答截图
func CreateQuestionAnswerEvidenceChain(evidencechain_param *evidenceChain.EvidenceChain_Param) (resList evidenceChain.QuestionNaireEvidenceChain_Result, err error) {
	resList, err = db.CreateQuestionAnswerEvidenceChain(evidencechain_param)
	return resList, err
}
