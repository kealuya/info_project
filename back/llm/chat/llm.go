package chat

import (
	_ "embed"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

const API_KEY = "app-NgrhFkARRZpbzMUudqt8m2RC"
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

	br := BlockingResponse{}
	client := resty.New()
	resp, err := client.R().SetBody(chatStr).EnableTrace().
		SetHeader("Authorization", "Bearer "+API_KEY).
		SetHeader("Content-Type", "application/json").
		SetResult(&br).
		Post(API_URL)
	if err != nil {
		fmt.Println("Error request:", err)
		os.Exit(1)
	}

	// 打印响应状态和数据
	if resp.StatusCode() != 200 {
		fmt.Println("Error server response:", resp.Error())
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
	Metadata       struct {
		Usage struct {
			PromptTokens        int     `json:"prompt_tokens"`
			PromptUnitPrice     string  `json:"prompt_unit_price"`
			PromptPriceUnit     string  `json:"prompt_price_unit"`
			PromptPrice         string  `json:"prompt_price"`
			CompletionTokens    int     `json:"completion_tokens"`
			CompletionUnitPrice string  `json:"completion_unit_price"`
			CompletionPriceUnit string  `json:"completion_price_unit"`
			CompletionPrice     string  `json:"completion_price"`
			TotalTokens         int     `json:"total_tokens"`
			TotalPrice          string  `json:"total_price"`
			Currency            string  `json:"currency"`
			Latency             float64 `json:"latency"`
		} `json:"usage"`
		CreatedAt int `json:"created_at"`
	}
}
