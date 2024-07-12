package module

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/astaxie/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"weixin_service_account_project/common"
	"weixin_service_account_project/llm"
	. "weixin_service_account_project/util"
)

type WeChatController struct {
	beego.Controller
}

// HandleMessage is used to handle the incoming messages from WeChat
func (c *WeChatController) HandleMessage() {

	var msg request
	err := xml.Unmarshal(c.Ctx.Input.RequestBody, &msg)
	common.ErrorHandler(err)
	// 系统性错误处理
	common.RecoverHandler(func(err error) {
		// 回复多条客服消息
		err = sendCustomMessage(msg.FromUserName, "服务器暂时无法回复，请稍后再试...")
		if err != nil {
			logs.Error(err)
		}
	})

	// 收到消息后的自动回复
	reply := response{
		XMLName:      xml.Name{Local: "xml"},
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      "正在思考，请稍等...",
	}

	output, err := xml.Marshal(reply)
	common.ErrorHandler(err)

	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/xml")
	_, err = c.Ctx.ResponseWriter.Write(output)
	common.ErrorHandler(err)

	// 显示"正在输入"状态
	fmt.Println(GetAccessToken())
	err = sendTypingStatus(GetAccessToken(), msg.FromUserName, "Typing")
	common.ErrorHandler(err)

	llmQuery := ""
	if msg.MsgType == "voice" {
		voice, err := llm.DownloadVoice(GetAccessToken(), msg.MediaId16K)
		common.ErrorHandler(err)

		wav, err := llm.ConvertAMRToPCM(voice)
		common.ErrorHandler(err)
		llmQuery = llm.CallVoiceToText(wav)
	} else {
		llmQuery = msg.Content
	}

	// 记录
	logs.Info("接受用户消息::", msg.MsgID, msg.FromUserName, msg.MsgType, llmQuery)

	type llmAnswer struct {
		answer string
		err    error
	}

	llmChannel := make(chan llmAnswer, 1)
	go func() {
		answer, err := llm.CallLlm(llmQuery)
		llmChannel <- llmAnswer{answer: answer, err: err}
		return
	}()

	select {

	case answer := <-llmChannel:
		if answer.err != nil {
			common.ErrorHandler(answer.err)
			return
		}
		// 回复客服消息
		err = sendCustomMessage(msg.FromUserName, answer.answer)
		common.ErrorHandler(err)

	case <-time.After(20 * time.Second): // 超时
		// 回复多条客服消息
		err = sendCustomMessage(msg.FromUserName, "服务器请求超时，请稍后再试...")
		if err != nil {
			logs.Error(err)
		}
	}

	// 取消"正在输入"状态
	err = sendTypingStatus(GetAccessToken(), msg.FromUserName, "CancelTyping")
	common.ErrorHandler(err)
}

// 标准交互方法，无需考虑错误
func sendCustomMessage(openID, text string) error {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/custom/send?access_token=%s", ACCESSTOKEN)
	message := map[string]interface{}{
		"touser":  openID,
		"msgtype": "text",
		"text": map[string]string{
			"content": text,
		},
	}

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", strings.NewReader(string(body)))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if val, ok := result["errcode"]; ok && val.(float64) != 0 {
		return fmt.Errorf("send custom message error: %v", result)
	}

	return nil
}

func sendTypingStatus(accessToken, openID, command string) error {
	typingURL := "https://api.weixin.qq.com/cgi-bin/message/custom/typing?access_token=%s"
	url := fmt.Sprintf(typingURL, accessToken)

	request := TypingRequest{
		ToUser:  openID,
		Command: command,
	}

	jsonData, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshalling request: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("error unmarshalling response: %v", err)
	}

	if errcode, ok := result["errcode"].(float64); ok && errcode != 0 {
		return fmt.Errorf("API error: %v", result["errmsg"])
	}

	return nil
}

type request struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	MediaId      string   `xml:"MediaId"`
	Format       string   `xml:"Format"`
	Content      string   `xml:"Content"`
	MsgID        int64    `xml:"MsgId"`
	MediaId16K   string   `xml:"MediaId16K"`
}

type response struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
}

type TypingRequest struct {
	ToUser  string `json:"touser"`
	Command string `json:"command"`
}
