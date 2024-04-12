package controllers

type JsonStruct struct {
	Success bool        `json:"success"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	//UserInfos []ht.LoginUserInfo `json:"userInfos"`
}

func NewJsonStruct(data interface{}) *JsonStruct {
	return &JsonStruct{
		Success: false,
		Msg:     "服务器错误",
		Data:    data,
	}
}
