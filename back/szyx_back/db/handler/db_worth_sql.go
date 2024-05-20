package handler

const (
	//申请价值 修改价值数据
	ModifyWorthByApplyWorth_sql = `update kdxf_worth set
									status = ?,applyTime = ?
								where 
									userId = ? and corpCode = ? and worthId = ?
								`
	//查询价值
	SelectWorthById_sql = `select 
						   worth.*,
						   km.taskTitle,km.taskData,
						   km.taskContent,meeting.meetingId,
						   meeting.meetingTitle  from  kdxf_worth worth  
						   left join kdxf_taskpool km on  worth.taskId = km.taskId 
						   left join kdxf_meeting meeting on  worth.taskId = meeting.taskId 
						   where worth.corpCode = ? and worth.userId = ? and worth.worthId = ? 
`

	GetWorthList_sql = `select * from  kdxf_worth where corpCode = ? and userId = ? and status = ? ORDER BY 
						createTime DESC  limit ?,? `

	GetWorthListCount_sql = `select * from  kdxf_worth where corpCode = ? and userId = ? and status = ? ORDER BY createTime DESC `

	CreateWorth_sql = `insert into kdxf_worth(worthId,taskId,worthScore,worthTitle,
											  status,money,userId,userName,
											  userMobile,corpName,corpCode,
											  createTime,creater,bz1,bz2,bz3)
						values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
)
