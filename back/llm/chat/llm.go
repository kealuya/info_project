package chat

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const API_KEY = "app-mVIfxhcLyJ3KBhUhTYMTbjuP"
const API_URL = "http://122.9.41.45/v1/chat-messages"

func NewInstance(conversationID string) Chat {
	return Chat{
		ResponseMode:   blocking,
		ConversationID: conversationID,
	}
}

type ResponseMode string

const (
	streaming ResponseMode = "streaming"
	blocking  ResponseMode = "blocking"
)

type Chat struct {
	ResponseMode   ResponseMode
	ConversationID string
}

func (receiver *Chat) Conversation(query string) BlockingResponse {

	// 将请求体序列化为JSON字符串
	chatStr := ChatStruct{
		Inputs:         nil,
		Query:          query,
		ResponseMode:   string(receiver.ResponseMode),
		ConversationID: receiver.ConversationID,
		User:           "llm",
		Files:          nil,
	}
	jsonData, err := json.Marshal(chatStr)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	// 创建新的HTTP请求
	req, err := http.NewRequest("POST", API_URL, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}

	// 设置请求头
	req.Header.Set("Authorization", "Bearer "+API_KEY)
	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// 打印响应状态和数据
	if resp.StatusCode != 200 {
		fmt.Println("Error server response:", resp.Status)
		os.Exit(1)
	}
	br := BlockingResponse{}
	b, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(b, &br)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		os.Exit(1)
	}

	return br
}

type ChatStruct struct {
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	ResponseMode   string                 `json:"response_mode"`
	ConversationID string                 `json:"conversation_id"`
	User           string                 `json:"user"`
	Files          []struct {
		Type           string `json:"type"`
		TransferMethod string `json:"transfer_method"`
		URL            string `json:"url"`
	} `json:"files"`
}
type BlockingResponse struct {
	MessageId      string `json:"message_id"`
	ConversationId string `json:"conversation_id"`
	Mode           string `json:"mode"`
	Answer         string `json:"answer"`
	CreatedAt      int    `json:"created_at"`
}
