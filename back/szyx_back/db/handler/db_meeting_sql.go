package handler

const (
	InsertMeeting_sql = `INSERT INTO kdxf_meeting 
						( 
							meetingId,
							meetingTitle,
							meetingType,
							meetingTime,
							meetingCity,
							meetingAddress,
							meetingPeople,
							meetingFlag,
							corpName,
							corpCode,
							createTime,
							creater,
							bz1,
							bz2,
							bz3
						) 
							VALUE
						(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	ModifyMeetingAudioFileUrl_sql = `update kdxf_meeting set
									MeetingAudioFileUrl = ?
								where 
									MeetingId = ?
								`

	GetMeetingById_sql =   `select meeting.*,
							speech.content as meetingAudioText,
							knowledge.summary as meetingSummary,
							knowledge.meeting_minutes as meetingMinutes,
							knowledge.brain  as meetingBrainMap
							from  kdxf_meeting meeting  
							left join kdxf_speech speech on meeting.meetingId = speech.meetingId  
							left join kdxf_knowledge knowledge on meeting.meetingId = knowledge.meetingId 
							where meeting.MeetingId = ? and creater = ? `

	GetMeetingList_sql = `select * from  kdxf_meeting where corpCode = ? and creater = ? `

	GetMeetingListByTaskId_sql = `select * from  kdxf_meeting where corpCode = ? and creater = ? and taskId = ? ORDER BY createTime DESC `

	GetMeetingFileList_sql = `select * from  kdxf_meetingFile where meetingId = ? and creater = ? `

	AddMeetingFileInfo_sql = `INSERT INTO kdxf_meetingFile ( 
							  meetingId,fileType,audioTime,
							  fileName,fileUrl,creater,corpName,
							  corpCode,createTime) VALUE  (?,?,?,?,?,?,?,?,?)`

	DeleteAudioMeetingFile_sql = ` delete from kdxf_meetingFile where meetingId = ? and fileName = ? 
                                 and creater = ? and corpCode = ? `
)
