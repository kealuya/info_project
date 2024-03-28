package handler

const (
	InsertMeeting = `INSERT INTO szyx_meeting 
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
)
