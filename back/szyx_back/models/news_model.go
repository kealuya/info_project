package models

import (
	"szyx_back/db"
	"szyx_back/entity/newsInfo"
	"time"
)

// 推广新闻发布列表list
func GetNewsTemplateList(newsInfo *newsInfo.NewsList_Param) (newsInfoList newsInfo.NewsList_Result, err error) {
	newsInfoList, err = db.GetNewsTemplateList(newsInfo)
	return newsInfoList, err
}

// 视频新闻发布列表list
func GetVideoTemplateList(videoInfo *newsInfo.VideoList_Param) (videoInfoList newsInfo.VideoList_Result, err error) {
	videoInfoList, err = db.GetVideoTemplateList(videoInfo)
	return videoInfoList, err
}

// 获取推广新闻内容详情
func GetNewsDetail(newsInfoPar *newsInfo.NewsInfo) (newsInfoRes newsInfo.NewsInfo, err error) {
	newsInfoRes, err = db.GetNewsDetail(newsInfoPar)
	return newsInfoRes, err
}

// 新增视频新闻
func AddVideoTemplate(videoInfoPar *newsInfo.VideoInfo) (err error) {
	//判断是否存在生成的创建时间
	if videoInfoPar.CreateTimeNew == "" {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		videoInfoPar.CreateTime = currentTime
	} else {
		videoInfoPar.CreateTime = videoInfoPar.CreateTimeNew
	}
	err = db.AddVideoTemplate(videoInfoPar)
	return err
}

// 超级管理员、用户 - 新增推广视频
func AddVideoInUser(videoInfoPar *newsInfo.VideoInfoInUser) (err error) {
	//查询计划周期  生成周期范围内的创建时间
	createTime := db.GetPlanTime_RandomDateInRange(videoInfoPar.PlanId, videoInfoPar.CorpCode)
	videoInfoPar.CreateTime = createTime
	err = db.AddVideoInUser(videoInfoPar)
	return err
}

// 新增推广新闻
func AddNews(newsInfoInfoPar *newsInfo.NewsInfo) (err error) {
	//判断是否存在生成的创建时间
	if newsInfoInfoPar.CreateTimeNew == "" {
		currentTime := time.Now().Format("2006-01-02 15:04:05")
		newsInfoInfoPar.CreateTime = currentTime
	} else {
		newsInfoInfoPar.CreateTime = newsInfoInfoPar.CreateTimeNew
	}
	err = db.AddNews(newsInfoInfoPar)
	return err
}

// 超级管理员、用户 - 新增推广新闻
func AddNewsInUser(newsInfoInfoPar *newsInfo.NewsInfoInUser) (err error) {
	//查询计划周期  生成周期范围内的创建时间
	createTime := db.GetPlanTime_RandomDateInRange(newsInfoInfoPar.PlanId, newsInfoInfoPar.CorpCode)
	newsInfoInfoPar.CreateTime = createTime
	err = db.AddNewsInUser(newsInfoInfoPar)
	return err
}

// 修改推广新闻
func ModifyNews(newsInfoInfoPar *newsInfo.NewsInfo) (err error) {
	err = db.ModifyNews(newsInfoInfoPar)
	return err
}

// 删除视频媒体
func DeleteVideoNews(videoInfoInfoPar *newsInfo.VideoInfo) (err error) {
	err = db.DeleteVideoNews(videoInfoInfoPar)
	return err
}

// 删除推广信息
func DeleteNews(newsInfoPar *newsInfo.NewsInfo) (err error) {
	err = db.DeleteNews(newsInfoPar)
	return err
}
