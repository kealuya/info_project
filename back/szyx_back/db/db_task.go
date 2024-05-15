package db

import (
	"errors"
	"fmt"
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity/meeting"
	"szyx_back/entity/task"
	"time"
)

//修改任务池中的任务
func ModifyTaskPool(info *task.Task) (res task.Task, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//修改任务池
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.TaskStatus)
	Param = append(Param, currentTime)
	Param = append(Param, info.TaskId)

	num, err := dbHandler.Update(db_handler.ModifyTaskPool_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "任务池数据库修改发生错误!")
	} else {
		//查询任务池，返回前台展示
		res1, err1 := dbHandler.SelectOne(db_handler.SelectTaskPoolById_sql, info.TaskId)
		if len(res1) > 0 && err1 == nil {
			decoder := ObtainDecoderConfig(&res)
			err := decoder.Decode(res1)
			common.ErrorHandler(err, "任务池数据库保存转换发生错误!")
		} else {
			err = err1
		}
	}

	return res, err
}

//查看任务池列表
func GetTaskPoolList(info *task.TaskList_Param) (res task.TaskList_Result, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//任务池信息
	var Param []interface{}
	Param = append(Param, info.UserId)
	Param = append(Param, info.CorpCode)
	Param = append(Param, info.Status)
	//根据用户是否参与过任务进行筛选
	sqlType := ""
	if info.UserJoinState != "" {
		Param = append(Param, info.UserJoinState)
		sqlType = sqlType + " WHERE userJoinState = ?  ORDER BY  taskpool.createTime DESC  limit ?,? "
	}else{
		sqlType = sqlType + " ORDER BY  taskpool.createTime DESC  limit ?,? "
	}

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)
	selRes, err := dbHandler.SelectList(db_handler.GetTaskPoolList_sql + sqlType, Param...)

	taskList := []task.Task{}
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&taskList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "任务池列表信息转换发生错误!")
	}
	taskListCount := []task.Task{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.UserId)
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.Status)
	//根据用户是否参与过任务进行筛选
	sqlCountType := ""
	if info.UserJoinState != "" {
		ParamCount = append(ParamCount, info.UserJoinState)
		sqlCountType = sqlCountType + " WHERE userJoinState = ?  ORDER BY  taskpool.createTime DESC "
	}else{
		sqlCountType = sqlCountType + " ORDER BY  taskpool.createTime DESC "
	}
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetTaskPoolList_sql + sqlCountType, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&taskListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "任务池列表分页信息转换发生错误!")
	}
	res.TaskList = taskList
	res.TotalCount = int64(len(taskListCount))
	//获取总页数，前端需要
	res.PageCount = res.TotalCount / info.PageSize
	if res.TotalCount%info.PageSize > 0 {
		res.PageCount++
	}
	return res, err
}

// 任务池 任务详情
func TaskPoolDetails(taskId string) (res task.Task, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//我的任务信息
	var Param []interface{}
	Param = append(Param, taskId)
	selRes, err := dbHandler.SelectOne(db_handler.TaskPoolDetails_sql, Param...)
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&res)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "查询单条我的任务转换发生错误!")
	}
	return res, err
}

//校验是否能参加任务，一个任务只能参加一次
func CheckIsJoinTask(myTask_Param *task.MyTask) (result string, msg string, err error) {
	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}
	Param = append(Param, myTask_Param.CorpCode)
	Param = append(Param, myTask_Param.UserId)
	Param = append(Param, myTask_Param.TaskId)
	selRes, err := dbHandler.SelectList(db_handler.CheckIsJoinTask_sql, Param...)
	//判断用户是否参与过任务
	if len(selRes) > 0 && err == nil {
		result = "yesJoinTask" //已参与任务
		msg = "您已参与任务，请勿重复参与。"
	} else {
		result = "noJoinTask" //未参与任务
		msg = "用户未参与任务"
	}
	return result,msg, err
}

//参加任务--添加我得任务
func ApplyJoinTask(info *task.MyTask) (msg error) {
	dbHandler := db_handler.NewDbHandler()
	//修改任务池
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.TaskId)
	Param = append(Param, info.CorpName)
	Param = append(Param, info.CorpCode)
	Param = append(Param, currentTime)
	Param = append(Param, info.Creater)
	Param = append(Param, info.Bz1)
	Param = append(Param, info.Bz2)
	Param = append(Param, info.Bz3)
	Param = append(Param, info.UserId)
	Param = append(Param, info.UserName)
	Param = append(Param, info.UserMobile)
	Param = append(Param, common.MY_TASK_FLAG_KEY_0) // 任务状态 0.待完成 1.已完成。首次调用申请任务 固定为待完成
	Param = append(Param, info.FinishTime)
	num, err := dbHandler.Insert(db_handler.ApplyJoinTask_sql, Param...)
	if num <= 0 || err != nil {
		msg = err
	}

	return msg
}

//====================================================我的任务========================

//查看我的任务列表
func GetTaskList(info *task.MyTaskList_Param) (res task.MyTaskList_Result, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//我的任务信息
	var Param []interface{}
	Param = append(Param, info.CorpCode)
	Param = append(Param, info.Flag)
	Param = append(Param, info.UserId)

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)
	var sqlOrderBy = ""
	//根据任务完成状态判断使用什么时间进行排序
	if info.Flag == common.MY_TASK_FLAG_KEY_1 {
	 	sqlOrderBy =  " ORDER BY mytask.finishTime DESC  limit ?,?"
	}else{
		sqlOrderBy =  " ORDER BY mytask.createTime DESC  limit ?,?"
	}
	selRes, err := dbHandler.SelectList(db_handler.GetTaskList_sql + sqlOrderBy, Param...)
	myTaskList := []task.MyTask{}

	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&myTaskList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "我的任务列表信息转换发生错误!")
	}
	//查询列表总条数
	myTaskListCount := []task.MyTask{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.Flag)
	ParamCount = append(ParamCount, info.UserId)
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetTaskListCount_sql, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&myTaskListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "我的任务列表分页信息转换发生错误!")
	}
	res.MyTaskList = myTaskList
	res.TotalCount = int64(len(myTaskListCount))
	//获取总页数，前端需要
	res.PageCount = res.TotalCount / info.PageSize
	if res.TotalCount%info.PageSize > 0 {
		res.PageCount++
	}
	return res, err
}

//完成任务--生成价值数据
func FinishMyTask(info *task.MyTask) (msg error) {
	dbHandler := db_handler.NewDbHandler()
	tx, _ := dbHandler.Db.Begin()
	defer common.RecoverHandler(func(err error) {
		_ = tx.Rollback()
		msg = errors.New(fmt.Sprintf("%s%s", "完成我的任务::", err.Error()))
	})
	//当前时间生成
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	//完成任务数据组装
	var finishParam []interface{}
	finishParam = append(finishParam, common.MY_TASK_FLAG_KEY_1) //修改状态为已完成
	finishParam = append(finishParam, currentTime)
	finishParam = append(finishParam, info.TaskId)
	finishParam = append(finishParam, info.UserId)
	finishParam = append(finishParam, info.CorpCode)
	num, err := tx.Exec(db_handler.ModifyMyTask_sql, finishParam...)
	rows, _ := num.RowsAffected()
	if rows < 0 || err != nil {
		errors.New("修改我的任务发生错误!")
		return
	}

	//遍历完成任务所选的会议列表，根据会议ID 更新会议表所关联的任务ID。一个任务对应多个会议
	for _, meetingId := range info.MeetingIdList {
		var ParamMeeting []interface{}
		ParamMeeting = append(ParamMeeting, info.TaskId)                  //任务id
		ParamMeeting = append(ParamMeeting, common.MY_MEETING_FLAG_KEY_1) //修改会议使用状态为 已使用:1
		ParamMeeting = append(ParamMeeting, meetingId)                    // 会议id
		ParamMeeting = append(ParamMeeting, info.UserId)
		ParamMeeting = append(ParamMeeting, info.CorpCode)
		numMeeting, err := tx.Exec(db_handler.UpDateMeetingInTackId_sql, ParamMeeting...)
		rows3, _ := numMeeting.RowsAffected()
		if rows3 < 0 || err != nil {
			errors.New("更新会议表所关联的任务ID发生错误!")
			return
		}
	}

	//完成任务后 价值数据生成组装
	var Param []interface{}
	btTime := time.Now().Format("2006-01-02")
	Param = append(Param, info.TaskId)                      //价值ID //TODO 目前先用任务id 作为价值id，目前业务是一个完成的任务对应一条价值
	Param = append(Param, "85")                             //价值评分 //TODO 评分怎样来的，还不确定，是否需要管理端设定，先固定值
	Param = append(Param, btTime+"-"+info.UserName+"-价值申请") //价值标题
	Param = append(Param, common.MY_WORTH_APPLY_FLAG_KEY_0) //价值状态 0：未申请   1：已申请
	Param = append(Param, "5000.00")                        //价值申请金额  //TODO 目前不知道价值申请金额怎么定义  先固定值
	Param = append(Param, info.UserId)
	Param = append(Param, info.UserName)
	Param = append(Param, info.UserMobile)
	Param = append(Param, info.CorpName)
	Param = append(Param, info.CorpCode)
	Param = append(Param, currentTime) //创建时间
	Param = append(Param, info.UserId)
	Param = append(Param, info.Bz1)
	Param = append(Param, info.Bz2)
	Param = append(Param, info.Bz3)
	insertNum, err := tx.Exec(db_handler.CreateWorth_sql, Param...)
	rows2, _ := insertNum.RowsAffected()
	if rows2 < 0 || err != nil {
		errors.New("添加价值数据发生错误!")
		return
	}
	//执行操作
	comRrr := tx.Commit()
	common.ErrorHandler(comRrr)
	return msg
}

//查看我的任务详情
func MyTaskDetails(info *task.MyTask) (res task.MyTask, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//我的任务信息
	var Param []interface{}
	Param = append(Param, info.TaskId)
	Param = append(Param, info.UserId)
	Param = append(Param, info.CorpCode)
	selRes, err := dbHandler.SelectOne(db_handler.MyTaskDetails_sql, Param...)
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&res)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "查询单条我的任务转换发生错误!")
	}
	//查询任务对应的会议信息
	myTaskMeeting := []meeting.Meeting{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.UserId)
	ParamCount = append(ParamCount, info.TaskId)
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetMeetingListByTaskId_sql, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&myTaskMeeting)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "我的任务完成会议信息转换发生错误!")
	}
	res.MeetingList = myTaskMeeting
	return res, err
}

//用户放弃未完成的任务
func GiveUpTask(myTask_Param *task.Task) (msg error) {
	defer common.RecoverHandler(func(rcErr error) {
		msg = rcErr
	})
	dbHandler := db_handler.NewDbHandler()
	var Param []interface{}
	Param = append(Param, myTask_Param.CorpCode)
	Param = append(Param, myTask_Param.Creater)
	Param = append(Param, myTask_Param.TaskId)
	_, err := dbHandler.Delete(db_handler.GiveUpTask_sql, Param...)
	common.ErrorHandler(err, "未完成的任务删除操作记录发生错误")
	var err1 error
	if err != nil {
		err1 = errors.New("未完成的任务删除失败")
	}
	return err1
}