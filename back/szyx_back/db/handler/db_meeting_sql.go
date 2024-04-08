package handler

const (
	InsertMeeting_sql = `INSERT INTO kdxf_meeting 
						( 
							MeetingId,
							MeetingTitle,
							MeetingTime,
							MeetingCity,
							MeetingAddress,
							MeetingPeople,
							MeetingAudioFileUrl,
							CorpName,
							CorpCode,
							CreateTime,
							Creater,
							Bz1,
							Bz2,
							Bz3
						) 
							VALUE
						(?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	ModifyMeetingAudioFileUrl_sql = `update kdxf_meeting set
									MeetingAudioFileUrl = ?
								where 
									MeetingId = ?
								`

	GetMeetingById_sql = `select * from  kdxf_meeting where MeetingId = ? `

	GetMeetingList_sql = `select * from  kdxf_meeting where corpCode = ? ORDER BY 
						createTime DESC  limit ?,? `

	GetMeetingListCount_sql = `select * from  kdxf_meeting where corpCode = ? ORDER BY createTime DESC `
)
