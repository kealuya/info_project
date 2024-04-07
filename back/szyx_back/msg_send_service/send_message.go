package msg_send_service

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"szyx_back/common"
	"szyx_back/entity/system"
)

//短信服务地址
var Message_Service_Url = "XXXXXX"

func Send_AlibabaMessage(phone string, code string) (messageReturn system.MessageAlibabaReturn) {
	messageAlibaba := new(system.MessageAlibaba)
	messageAlibaba.Code = code
	messageAlibaba.Phone = phone
	messageJson := common.Marshal(messageAlibaba)
	logs.Debug("手机号登录调用阿里云短信获取验证码入参：" + string(messageJson))
	respBodyMessage := common.DoHttpPost(Message_Service_Url, string(messageJson))
	logs.Debug("手机号登录调用阿里云短信获取验证码返回：" + respBodyMessage)
	//接口返回的josn转对象
	str := []byte(respBodyMessage)
	responseData := new(system.MessageAlibabaReturn)
	json.Unmarshal(str, &responseData)
	if responseData.Success == false {
		messageReturn.Msg = responseData.Msg
		return messageReturn
	} else {
		messageReturn.Success = true
		messageReturn.Msg = responseData.Msg
	}
	return messageReturn
}
