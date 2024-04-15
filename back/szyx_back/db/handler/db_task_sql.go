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

	GetTaskList_sql = `select * from  kdxf_mytask where corpCode = ? and flag = ? ORDER BY 
						createTime DESC  limit ?,? `

	GetTaskListCount_sql = `select * from  kdxf_mytask where corpCode = ? and flag = ? ORDER BY 
						createTime DESC`

	ModifyMyTask_sql = `update kdxf_mytask set
									flag = ?,finishTime = ?
								where 
									taskId = ? and userId = ? and corpCode = ?
								`
)
