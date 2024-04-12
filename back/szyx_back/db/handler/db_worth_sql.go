package handler

const (
	//申请价值 修改价值数据
	ModifyWorthByApplyWorth_sql = `update kdxf_worth set
									status = ?,applyTime = ?
								where 
									userId = ?
								`
	GetWorthList_sql = `select * from  kdxf_worth where corpCode = ? and userId = ? and status = ? ORDER BY 
						createTime DESC  limit ?,? `

	GetWorthListCount_sql = `select * from  kdxf_meeting where corpCode = ? and userId = ? and status = ? ORDER BY createTime DESC `
)
