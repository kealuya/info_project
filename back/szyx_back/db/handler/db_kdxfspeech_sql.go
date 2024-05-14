package handler

const (
	SelectSpeechById_sql = `select 
							order_id as orderId,
							meetingId,
							no,
							file_name as fileName
							from  kdxf_speech where meetingId = ?  `
)
