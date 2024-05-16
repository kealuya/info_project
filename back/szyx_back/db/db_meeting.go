package db

import (
	"errors"
	"fmt"
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity/meeting"
	"time"
)

//创建会议
func CreateMeeting(info *meeting.Meeting) (msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.MeetingId)
	Param = append(Param, info.MeetingTitle)
	Param = append(Param, info.MeetingType)
	Param = append(Param, info.MeetingTime)
	Param = append(Param, info.MeetingCity)
	Param = append(Param, info.MeetingAddress)
	Param = append(Param, info.MeetingPeople)
	Param = append(Param, "0") //会议是否使用 0：未使用   1：已使用
	Param = append(Param, info.CorpName)
	Param = append(Param, info.CorpCode)
	Param = append(Param, currentTime)
	Param = append(Param, info.Creater)
	Param = append(Param, info.Bz1)
	Param = append(Param, info.Bz2)
	Param = append(Param, info.Bz3)

	num, err := dbHandler.Insert(db_handler.InsertMeeting_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "创建会议发生错误!")
	}
	return err
}

//查看会议列表
func GetMeetingList(info *meeting.MeetingList_Param) (res meeting.MeetingList_Result, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	var Param []interface{}
	Param = append(Param, info.CorpCode)
	Param = append(Param, info.UserId)
	//根据类型筛选
	sqlType := ""
	if info.MeetingType != "" {
		Param = append(Param, info.MeetingType)
		sqlType = sqlType + " and meetingType = ? "
	}
	//根据会议是否使用筛选
	sqlMeetingFlag := ""
	if info.MeetingFlag != "" {
		Param = append(Param, info.MeetingFlag)
		sqlMeetingFlag = sqlMeetingFlag + " and meetingFlag = ? "
	}
	//根据时间范围筛选
	sqlTime := ""
	if info.StartTime != "" && info.EndTime != "" {
		Param = append(Param, info.StartTime)
		Param = append(Param, info.EndTime)
		sqlTime = sqlTime + " and DATE(createTime) BETWEEN ? AND ? "
	}

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)
	sqlOrderby := "ORDER BY createTime DESC  limit ?,?"
	selRes, err := dbHandler.SelectList(db_handler.GetMeetingList_sql+sqlType+sqlMeetingFlag+sqlTime+sqlOrderby, Param...)
	fmt.Print(db_handler.GetMeetingList_sql + sqlType + sqlTime + sqlOrderby)
	meetingList := []meeting.Meeting{}

	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&meetingList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "会议列表信息转换发生错误!")
	}
	meetingListCount := []meeting.Meeting{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.UserId)
	//根据类型筛选
	sqlType_Count := ""
	if info.MeetingType != "" {
		ParamCount = append(ParamCount, info.MeetingType)
		sqlType_Count = sqlType_Count + " AND meetingType = ? "
	}
	//根据会议是否使用筛选
	sqlMeetingFlag_Count := ""
	if info.MeetingFlag != "" {
		ParamCount = append(ParamCount, info.MeetingFlag)
		sqlMeetingFlag_Count = sqlMeetingFlag_Count + " and meetingFlag = ? "
	}
	//根据时间范围筛选
	sqlTime_Count := ""
	if info.StartTime != "" && info.EndTime != "" {
		ParamCount = append(ParamCount, info.StartTime)
		ParamCount = append(ParamCount, info.EndTime)
		sqlTime_Count = sqlTime_Count + " AND DATE(createTime) BETWEEN ? AND ? "
	}

	selCountRes, err2 := dbHandler.SelectList(db_handler.GetMeetingList_sql+sqlType_Count+sqlMeetingFlag_Count+sqlTime_Count, ParamCount...)
	//fmt.Print(db_handler.GetMeetingListCount_sql + sqlType_Count + sqlTime_Count + sqlOrderby_Count)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&meetingListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "会议列表分页信息转换发生错误!")
	}
	res.MeetingList = meetingList
	res.TotalCount = int64(len(meetingListCount))
	//获取总页数，前端需要
	res.PageCount = res.TotalCount / info.PageSize
	if res.TotalCount%info.PageSize > 0 {
		res.PageCount++
	}
	return res, err
}

//根据会议id 查询会议文件list
func GetMeetingFileList(meetingId string, userId string) (res []meeting.MeetingFile, msg error) {
	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}
	Param = append(Param, meetingId)
	Param = append(Param, userId)
	selRes, err := dbHandler.SelectList(db_handler.GetMeetingFileList_sql, Param...)
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&res)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "会议文件list列表信息转换发生错误!")
	}
	return res, err
}

//修改会议
func ModifyMeeting(info *meeting.Meeting) (res meeting.Meeting, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	//currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, "")
	Param = append(Param, info.MeetingId)

	num, err := dbHandler.Update(db_handler.ModifyMeetingAudioFileUrl_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "修改会议录音地址发生错误!")
	} else {
		//查询会议，返回前台展示
		res1, err1 := dbHandler.SelectOne(db_handler.GetMeetingById_sql, info.MeetingId)
		if len(res1) > 0 && err1 == nil {
			decoder := ObtainDecoderConfig(&res)
			err := decoder.Decode(res1)
			common.ErrorHandler(err, "会议展示信息转换发生错误!")
		} else {
			err = err1
		}
	}

	return res, err
}

//查看会议详情
func GetMeetingDetails(info *meeting.Meeting) (res meeting.Meeting, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//会议信息
	var Param []interface{}
	Param = append(Param, info.MeetingId)
	Param = append(Param, info.Creater)

	//查询会议主表
	num, err := dbHandler.SelectOne(db_handler.GetMeetingById_sql, Param...)
	if err != nil {
		common.ErrorHandler(err, "会议详情查询发生错误!")
	} else {
		//查询会议，返回前台展示
		decoder := ObtainDecoderConfig(&res)
		err := decoder.Decode(num)
		common.ErrorHandler(err, "会议详情转换发生错误!")
		meetingFileList := []meeting.MeetingFile{}
		num2, err2 := dbHandler.SelectList(db_handler.GetMeetingFileList_sql, Param...)
		if err2 != nil {
			common.ErrorHandler(err2, "会议文件详情查询发生错误!")
		} else {
			//查询会议，返回前台展示
			decoder1 := ObtainDecoderConfig(&meetingFileList)
			err3 := decoder1.Decode(num2)
			common.ErrorHandler(err3, "会议文件详情转换发生错误!")
			res.MeetingFile = meetingFileList
		}
	}

	return res, err
}

//上传文件，文件基础信息存表。用于选择业务内容关联展示
func AddMeetingFileInfo(meetingFile *meeting.MeetingFile) (msg error) {
	dbHandler := db_handler.NewDbHandler()
	//创建时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, meetingFile.MeetingId)
	//Param = append(Param, meetingFile.MeetingTitle)
	Param = append(Param, meetingFile.FileType)
	Param = append(Param, meetingFile.AudioTime)
	Param = append(Param, meetingFile.FileName)
	Param = append(Param, meetingFile.FileUrl)
	Param = append(Param, meetingFile.Creater)
	Param = append(Param, meetingFile.CorpName)
	Param = append(Param, meetingFile.CorpCode)
	Param = append(Param, currentTime)
	num, err := dbHandler.Insert(db_handler.AddMeetingFileInfo_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "保存文件基础信息存表发生错误!")
	}
	return err
}

//删除音频会议上传的 音频文件
func DeleteAudioMeetingFile(meetingFile *meeting.MeetingFile) (msg error) {
	defer common.RecoverHandler(func(rcErr error) {
		msg = rcErr
	})
	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}
	Param = append(Param, meetingFile.MeetingId)
	Param = append(Param, meetingFile.FileName)
	Param = append(Param, meetingFile.Creater)
	Param = append(Param, meetingFile.CorpCode)
	_, err := dbHandler.Delete(db_handler.DeleteAudioMeetingFile_sql, Param...)
	common.ErrorHandler(err, "删除音频文件操作记录发生错误")
	var err1 error
	if err != nil {
		err1 = errors.New("删除音频文件删除失败")
	}
	return err1
}

//获取3天内，待转译的会议
func GetMeetingListTranslation() (res []meeting.Meeting, msg error) {
	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}

	// 获取当前时间
	now := time.Now()
	// 计算三天前的时间
	threeDaysAgo := now.AddDate(0, 0, -3)

	startDate := now.Format("2006-01-02 15:04:05")
	endDate := threeDaysAgo.Format("2006-01-02 15:04:05")

	Param = append(Param, "1") //1 代表转译中
	Param = append(Param, startDate)
	Param = append(Param, endDate)
	selRes, err := dbHandler.SelectList(db_handler.GetMeetingListTranslation, Param...)
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&res)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "会议list信息转换发生错误!")
	}
	return res, err
}
