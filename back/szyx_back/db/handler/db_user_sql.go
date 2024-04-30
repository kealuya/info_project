package handler

const (
	Select_LoginUser = `SELECT * from kdxf_user `

	//查看用户会议计数
	Select_userMeetingCount = `SELECT Count(id) as meetingCount from kdxf_meeting where creater = ? and corpCode = ?  `

	//查看用户会议计数
	Select_userTaskCount = `SELECT Count(id) as taskCount from kdxf_mytask where userId = ? and corpCode = ?  `

	//查看价值钱数
	Select_userWorthCount = `SELECT Sum(money) as worthCount from kdxf_worth where userId = ? and corpCode = ? and status = ? `
)
