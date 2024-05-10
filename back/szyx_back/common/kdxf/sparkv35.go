package kdxf

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"szyx_back/secret"
	"time"
)

// 调用样例
func Test() {

	// 追加systemcontent描述信息，以及传入的问题需要符合Message规范，为了传入历史上下文
	msgs := []Message{
		{
			Role:    User,
			Content: "你是谁",
		},
		{
			Role:    Assistant,
			Content: "我是星火大模型",
		},
		{
			Role:    User,
			Content: "你会做什么",
		},
	}

	c, err := SparkV35("你是一个医生，只回答医疗相关问题，如果没关系的问题，请礼貌的拒绝", msgs)
	if err != nil {
		fmt.Println(err)
	}

	for m := range c {
		fmt.Println(m)
	}
}

var (
	hostUrlSpark   = secret.KDXF_SPARK_HOST_URL
	appidSpark     = secret.KDXF_SPARK_APP_ID
	apiSecretSpark = secret.KDXF_SPARK_API_SECRET
	apiKeySpark    = secret.KDXF_SPARK_API_KEY
)

// const SYSTEM_CONTENT = "你是一个专业的文职秘书，可以把描述的事情变成标准的会议纪要，其中可能包括参会人员、会议总结、各个角色的会话要点及后续计划等。"
//const SYSTEM_CONTENT = "你是一个医生，只回答医疗相关问题，如果没关系的问题，请礼貌的拒绝"

func SparkV35(systemContent string, questions []Message) (<-chan map[string]interface{}, error) {

	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	logs.Info("星火大模型问答开始")
	//握手并建立websocket 连接
	conn, resp, err := d.Dial(assembleAuthUrl1(hostUrlSpark, apiKeySpark, apiSecretSpark), nil)
	if err != nil {
		logs.Error(readResp(resp))
		return nil, err

	} else if resp.StatusCode != 101 {
		logs.Error(readResp(resp))
		return nil, err
	}

	myChannel := make(chan map[string]interface{})

	go func() {

		answer := ""
		data := genParams1(appidSpark, systemContent, questions)
		_ = conn.WriteJSON(data)
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				logs.Error("read message error:", err)
				return
			}

			var data map[string]interface{}
			err1 := json.Unmarshal(msg, &data)
			if err1 != nil {
				logs.Error("Error parsing JSON:", err)
				return
			}
			//fmt.Println(string(msg))
			//解析数据
			payload := data["payload"].(map[string]interface{})
			choices := payload["choices"].(map[string]interface{})
			header := data["header"].(map[string]interface{})
			code := header["code"].(float64)

			if code != 0 {
				logs.Error("code is not ok:", err)
				return
			}
			status := choices["status"].(float64)
			//fmt.Println(status)
			text := choices["text"].([]interface{})
			content := text[0].(map[string]interface{})["content"].(string)

			m := make(map[string]interface{})
			if status != 2 {
				answer += content
				m["content"] = content
				m["status"] = status

				myChannel <- m

			} else {
				answer += content
				usage := payload["usage"].(map[string]interface{})
				temp := usage["text"].(map[string]interface{})
				totalTokens := temp["total_tokens"].(float64)
				fmt.Println("total_tokens:", totalTokens)
				conn.Close()
				// todo 考虑存入db
				logs.Info("问答记录::", fmt.Sprintf("question->%s , answer->%s", questions, answer))
				logs.Async()
				m["content"] = content
				m["status"] = status

				myChannel <- m
				close(myChannel)
				break
			}
		}
	}()

	return myChannel, nil
}

type Answer struct {
	Header struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Sid     string `json:"sid"`
		Status  int    `json:"status"`
	} `json:"header"`
	Payload struct {
		Choices struct {
			Status int `json:"status"`
			Seq    int `json:"seq"`
			Text   []struct {
				Content string `json:"content"`
				Role    string `json:"role"`
				Index   int    `json:"index"`
			} `json:"text"`
		} `json:"choices"`
	} `json:"payload"`
}

// 生成参数
func genParams1(appid, systemContent string, questions []Message) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名

	messages := []Message{

		{Role: "system", Content: systemContent},
		//{Role: "user", Content: question},
	}
	messages = append(messages, questions...)

	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      "generalv3.5", // 根据实际情况修改返回的数据结构和字段名
				"temperature": float64(0.8),  // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(6),      // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(2048),   // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",     // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}
	return data // 根据实际情况修改返回的数据结构和字段名
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		fmt.Println(err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sgin := strings.Join(signString, "\n")
	// fmt.Println(sgin)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}

type SparkRole string

const (
	User      SparkRole = "user"
	Assistant           = "assistant"
)

type Message struct {
	Role    SparkRole `json:"role"`
	Content string    `json:"content"`
}
