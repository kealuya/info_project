package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
	"weixin_service_account_project/llm"
	. "weixin_service_account_project/util"
)

type WeChatController struct {
	beego.Controller
}

// VerifyToken is used to verify the server address is valid
func (c *WeChatController) VerifyToken() {
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")

	if validateWeChatRequest(TOKEN, signature, timestamp, nonce) {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("Signature verification failed")
	}
}

// validateWeChatRequest is used to validate the signature in the request
func validateWeChatRequest(token, signature, timestamp, nonce string) bool {
	strs := []string{token, timestamp, nonce}
	sort.Strings(strs) // Sort alphabetically
	str := strings.Join(strs, "")

	h := sha1.New()
	io.WriteString(h, str)
	sha1Hash := hex.EncodeToString(h.Sum(nil))

	return sha1Hash == signature
}

// HandleMessage is used to handle the incoming messages from WeChat
func (c *WeChatController) HandleMessage() {
	var msg request
	err := xml.Unmarshal(c.Ctx.Input.RequestBody, &msg)
	fmt.Println(string(c.Ctx.Input.RequestBody))
	if err != nil {
		c.Ctx.WriteString("Invalid request")
		return
	}
	// 收到消息后的自动回复
	func() {
		reply := response{
			XMLName:      xml.Name{Local: "xml"},
			ToUserName:   msg.FromUserName,
			FromUserName: msg.ToUserName,
			CreateTime:   time.Now().Unix(),
			MsgType:      "text",
			Content:      "正在思考，请稍等...",
		}

		output, err := xml.Marshal(reply)
		if err != nil {
			c.Ctx.WriteString("Failed to generate response")
			return
		}

		c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/xml")
		_, err = c.Ctx.ResponseWriter.Write(output)
		if err != nil {
			fmt.Println(err)
			c.Ctx.WriteString("Failed to generate response")
			return
		}
	}()

	go func() {
		// 显示"正在输入"状态
		err := sendTypingStatus(GetAccessToken(), msg.FromUserName, "Typing")
		if err != nil {
			fmt.Printf("Error sending typing status: %v\n", err)
			return
		}
		fmt.Println("Sent 'Typing' status")

		llmQuery := ""

		if msg.MsgType == "voice" {
			voice, err := llm.DownloadVoice(GetAccessToken(), msg.MediaId16K)
			if err != nil {
				fmt.Println(err)
			}

			wav, err := llm.ConvertAMRToPCM(voice)
			if err != nil {
				fmt.Println(err)
			}

			llmQuery = llm.CallVoiceToText(wav)
		} else {
			llmQuery = msg.Content
		}

		answer := llm.CallLlm(llmQuery)

		// 回复多条客服消息
		err = sendCustomMessage(msg.FromUserName, answer)
		if err != nil {
			fmt.Println("Failed to send first custom message:", err)
			return
		}

		// 取消"正在输入"状态
		err = sendTypingStatus(GetAccessToken(), msg.FromUserName, "CancelTyping")
		if err != nil {
			fmt.Printf("Error cancelling typing status: %v\n", err)
			return
		}
		fmt.Println("Sent 'CancelTyping' status")

	}()

}

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

func main() {
	// 用于注册微信返回地址
	beego.Router("/wechat", &WeChatController{}, "get:VerifyToken")
	// 用于接受用户提问，通过大模型回答
	beego.Router("/wechat", &WeChatController{}, "post:HandleMessage")

	// Start beego server
	beego.Run()
}
