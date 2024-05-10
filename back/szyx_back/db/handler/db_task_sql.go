package handler

const (
	//下架任务  修改任务池数据
	ModifyTaskPool_sql = `update kdxf_taskpool set
									status = ?,applyTime = ?
								where 
									taskId = ?
								`

	SelectTaskPoolById_sql = `select * from  kdxf_taskpool where taskId = ?`

	GetTaskPoolList_sql = `select * from  kdxf_taskpool where corpCode = ? and taskStatus = ? ORDER BY 
						createTime DESC  limit ?,? `

	GetTaskPoolListCount_sql = `select * from  kdxf_taskpool where corpCode = ?  and taskStatus = ? ORDER BY createTime DESC `

	CheckIsJoinTask_sql = `select * from kdxf_mytask where corpCode = ? and userId  = ?  and taskId = ?`
	//ApplyJoinTask_sql = `insert into kdxf_mytask(
	//												TaskId,
	//												TaskTitle,
	//												TaskTarget,
	//												TaskContent,
	//												TaskType,
	//												TaskStatus,
	//												TaskImg,
	//												CorpName,
	//												CorpCode,
	//												CreateTime,
	//												Creater,
	//												Bz1,
	//												Bz2,
	//												Bz3,
	//												UserId,
	//												UserName,
	//												UserMobile,
	//												Flag,
	//												FinishTime,
	//												TaskData
	//											)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	ApplyJoinTask_sql = `insert into kdxf_mytask(
													TaskId, 
													CorpName,
													CorpCode,
													CreateTime,
													Creater,
													Bz1,
													Bz2,
													Bz3,
													UserId,
													UserName,
													UserMobile,
													Flag,
													FinishTime
												)values(?,?,?,?,?,?,?,?,?,?,?,?,?)`

	GetTaskList_sql = `select mytask.*,kt.taskTitle ,kt.taskData ,kt.taskContent,kt.taskType,
						kt.taskImg ,kt.taskStatus  from kdxf_mytask mytask 
						left join kdxf_taskpool kt  ON mytask.taskId  = kt.taskId 
                        where mytask.corpCode = ? and mytask.flag = ? and mytask.userId = ?  `

	GetTaskListCount_sql = `select mytask.*,kt.taskTitle ,kt.taskData ,kt.taskContent,kt.taskType,
						kt.taskImg ,kt.taskStatus  from kdxf_mytask mytask 
						left join kdxf_taskpool kt  ON mytask.taskId = kt.taskId
						where mytask.corpCode = ? and mytask.flag = ? and mytask.userId = ? 
                        ORDER BY mytask.createTime DESC`

	ModifyMyTask_sql = `update kdxf_mytask set  flag = ?,finishTime = ?
						 where taskId = ? and userId = ? and corpCode = ? `

	MyTaskDetails_sql = `select mytask.*,kt.taskTitle ,kt.taskData ,kt.taskContent,kt.taskType,
						kt.taskImg ,kt.taskStatus  from kdxf_mytask mytask 
						left join kdxf_taskpool kt  ON mytask.taskId  = kt.taskId  
						where  mytask.taskId = ? and mytask.userId = ?  and mytask.corpCode = ? `

	//更新会议表所关联的任务ID，一个任务对应多个会议
	UpDateMeetingInTackId_sql = `update kdxf_meeting  set  taskId  = ?,meetingFlag  = ? 
						where meetingId  = ? and creater  = ? and corpCode = ? `

	GiveUpTask_sql = ` delete from kdxf_mytask where corpCode = ? and userId = ? and taskId = ? `
)
