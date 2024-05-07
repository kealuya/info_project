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

	GetMeetingById_sql = `select meeting.*, 
									speech.summary as meetingAudioSummary,speech.meetingMinutes as meetingAudioMinutes,  
									knowledge.summary as meetingDocumentSummary,knowledge.meeting_minutes as meetingDocumentMinutes 
									from  kdxf_meeting meeting  
									left join kdxf_speech speech on meeting.meetingId = speech.meetingId  
									left join kdxf_knowledge  knowledge  on meeting.meetingId = knowledge.meetingId 
								where meeting.MeetingId = ? `

	GetMeetingList_sql = `select * from  kdxf_meeting where corpCode = ? and creater = ? `

	GetMeetingListByTaskId_sql = `select * from  kdxf_meeting where corpCode = ? and creater = ? and taskId = ? ORDER BY createTime DESC `

	GetMeetingFileList_sql = `select * from  kdxf_meetingFile where meetingId = ? `

	AddMeetingFileInfo_sql = `INSERT INTO kdxf_meetingFile ( 
							  meetingId,meetingTitle,fileType,
							  fileName,fileUrl,creater,corpName,
							  corpCode,createTime) VALUE  (?,?,?,?,?,?,?,?,?)`
)
