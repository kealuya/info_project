package db

import (
	"errors"
	"fmt"
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
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
	Param = append(Param, info.CorpCode)
	Param = append(Param, info.Status)

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)

	selRes, err := dbHandler.SelectList(db_handler.GetTaskPoolList_sql, Param...)

	taskList := []task.Task{}

	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&taskList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "任务池列表信息转换发生错误!")
	}

	taskListCount := []task.Task{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.Status)
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetTaskPoolListCount_sql, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&taskListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "任务池列表分页信息转换发生错误!")
	}

	res.TaskList = taskList
	res.TotalCount = int64(len(taskListCount))

	return res, err
}

//参加任务--添加我得任务
func ApplyJoinTask(info *task.MyTask) (msg error) {
	dbHandler := db_handler.NewDbHandler()
	//修改任务池
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.TaskId)
	Param = append(Param, info.TaskTitle)
	//Param = append(Param, info.TaskTarget)
	//Param = append(Param, info.TaskContent)
	//Param = append(Param, info.TaskType)
	//Param = append(Param, info.TaskContent)
	//Param = append(Param, info.TaskStatus)
	//Param = append(Param, info.TaskImg)
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
	Param = append(Param, info.Flag)
	Param = append(Param, info.FinishTime)
	Param = append(Param, info.TaskData)
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

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)

	selRes, err := dbHandler.SelectList(db_handler.GetTaskList_sql, Param...)

	myTaskList := []task.MyTask{}

	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&myTaskList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "我的任务列表信息转换发生错误!")
	}

	myTaskListCount := []task.MyTask{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.Flag)
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetTaskListCount_sql, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&myTaskListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "我的任务列表分页信息转换发生错误!")
	}

	res.MyTaskList = myTaskList
	res.TotalCount = int64(len(myTaskListCount))

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

	//价值数据组装
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, info.TaskId)                      //价值ID
	Param = append(Param, info.TaskTitle)                   //价值标题
	Param = append(Param, "12000")                          //价值评分 //TODO 评分怎样来的，还不确定，是否需要管理端设定，先固定值
	Param = append(Param, common.MY_WORTH_APPLY_FLAG_KEY_0) //价值状态
	Param = append(Param, "12000")
	Param = append(Param, info.UserId)
	Param = append(Param, info.UserName)
	Param = append(Param, info.UserMobile)
	Param = append(Param, info.CorpName)
	Param = append(Param, info.CorpCode)
	Param = append(Param, currentTime) //创建时间
	Param = append(Param, "完成创建")
	Param = append(Param, info.Bz1)
	Param = append(Param, info.Bz2)
	Param = append(Param, info.Bz3)

	//完成任务数据组装
	var finishParam []interface{}
	finishParam = append(finishParam, common.MY_TASK_FLAG_KEY_1)
	finishParam = append(finishParam, currentTime)
	finishParam = append(finishParam, info.TaskId)
	finishParam = append(finishParam, info.UserId)
	finishParam = append(finishParam, info.CorpCode)

	num, err := tx.Exec(db_handler.ModifyMyTask_sql, finishParam...)
	rows, _ := num.RowsAffected()
	if (rows <= 0 || err != nil) {
		errors.New("修改我的任务状态发生错误!")
		return
	}
	insertNum, err := tx.Exec(db_handler.CreateWorth_sql, Param...)
	rows2, _ := insertNum.RowsAffected()
	if (rows2 <= 0 || err != nil) {
		errors.New("添加价值数据发生错误!")
		return
	}
	comRrr := tx.Commit()
	common.ErrorHandler(comRrr)

	return msg
}
