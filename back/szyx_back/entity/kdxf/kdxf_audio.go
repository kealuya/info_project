package kdxf

//语音识别返回值
type Kdxf_audio_result struct {
	Msg      string `json:"msg" description:"msg" `
	FileName string `json:"fileName" description:"文件名" `
	OrderId  string `json:"orderId" description:"ID" `
	Success  string `json:"success" description:"" `
}
