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
													TaskTitle,
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
													FinishTime,
													TaskData
												)values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	GetTaskList_sql = `select mytask.*,
								pool.taskTarget as taskTarget,
								pool.taskTarget as taskContent,
								pool.taskImg as taskImf
						from  kdxf_mytask mytask 
						left join kdxf_taskpool pool 
						on mytask.taskId = pool.taskId 
						where mytask.corpCode = ? and mytask.flag = ? and mytask.userId = ? ORDER BY mytask.createTime DESC  limit ?,? 
						`

	GetTaskListCount_sql = `select * from  kdxf_mytask where corpCode = ? and flag = ? and userId = ? ORDER BY 
						createTime DESC`

	ModifyMyTask_sql = `update kdxf_mytask set
									flag = ?,finishTime = ?,meetingId = ?
								where 
									taskId = ? and userId = ? and corpCode = ?
								`
	MyTaskDetails_sql = `select * from kdxf_mytask where corpCode = ? and taskId = ?
								`
)
