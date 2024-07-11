package llm

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"context"
)

func ConvertAMRToPCM(amrData []byte) ([]byte, error) {
	inputReader := bytes.NewReader(amrData)
	outputBuffer := bytes.NewBuffer(nil)

	// 添加更多的 FFmpeg 参数
	err := ffmpeg.Input("pipe:0").
		Output("pipe:1", ffmpeg.KwArgs{
			"f":      "wav",
			"acodec": "pcm_s16le",
			"ar":     "8000", // AMR 通常是 8kHz
			"ac":     "1",    // 单声道
		}).
		WithInput(inputReader).
		WithOutput(outputBuffer).
		Run()

	if err != nil {
		return nil, fmt.Errorf("FFmpeg conversion failed: %v", err)
	}

	if outputBuffer.Len() == 0 {
		return nil, fmt.Errorf("output buffer is empty")
	}

	return outputBuffer.Bytes(), nil
}

const (
	// 微信获取临时素材的 API URL
	mediaURL = "https://api.weixin.qq.com/cgi-bin/media/get?access_token=%s&media_id=%s"
)

func DownloadVoice(accessToken, mediaID string) ([]byte, error) {
	url := fmt.Sprintf(mediaURL, accessToken, mediaID)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error downloading voice: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return ioutil.ReadAll(resp.Body)
}

/**
 * 语音听写流式 WebAPI 接口调用示例 接口文档（必看）：https://doc.xfyun.cn/rest_api/语音听写（流式版）.html
 * webapi 听写服务参考帖子（必看）：http://bbs.xfyun.cn/forum.php?mod=viewthread&tid=38947&extra=
 * 语音听写流式WebAPI 服务，热词使用方式：登陆开放平台https://www.xfyun.cn/后，找到控制台--我的应用---语音听写---服务管理--上传热词
 * 注意：热词只能在识别的时候会增加热词的识别权重，需要注意的是增加相应词条的识别率，但并不是绝对的，具体效果以您测试为准。
 * 错误码链接：https://www.xfyun.cn/document/error-code （code返回错误码时必看）
 * @author iflytek
 */
var (
	hostUrl   = "wss://iat-api.xfyun.cn/v2/iat"
	appid     = "286be1ca"
	apiSecret = "ZGI3Y2QyNTJjZTMwYTZhMDE2MzFhN2Mz"
	apiKey    = "0dbfb37ce8900a9d53153f62ef81686a"
)

const (
	STATUS_FIRST_FRAME    = 0
	STATUS_CONTINUE_FRAME = 1
	STATUS_LAST_FRAME     = 2
)

func CallVoiceToText(voice []byte) string {
	st := time.Now()
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}

	// 握手并建立websocket连接
	conn, resp, err := d.Dial(assembleAuthUrl(hostUrl, apiKey, apiSecret), nil)
	if err != nil {
		panic(readResp(resp) + err.Error())
	} else if resp.StatusCode != 101 {
		panic(readResp(resp) + err.Error())
	}
	defer conn.Close()
	var intervel = 40 * time.Millisecond
	var frameSize = 1280 // 每一帧的音频大小
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 开启协程，发送数据
	go func() {
		status := STATUS_FIRST_FRAME
		for offset := 0; offset < len(voice); offset += frameSize {
			select {
			case <-ctx.Done():
				fmt.Println("session end ---")
				return
			default:
			}

			var buffer []byte
			if offset+frameSize > len(voice) {
				buffer = voice[offset:]
				status = STATUS_LAST_FRAME
			} else {
				buffer = voice[offset : offset+frameSize]
			}

			switch status {
			case STATUS_FIRST_FRAME:
				frameData := map[string]interface{}{
					"common": map[string]interface{}{
						"app_id": appid,
					},
					"business": map[string]interface{}{
						"language": "zh_cn",
						"domain":   "iat",
						"accent":   "mandarin",
					},
					"data": map[string]interface{}{
						"status":   STATUS_FIRST_FRAME,
						"format":   "audio/L16;rate=8000",
						"audio":    base64.StdEncoding.EncodeToString(buffer),
						"encoding": "raw",
					},
				}
				conn.WriteJSON(frameData)
				status = STATUS_CONTINUE_FRAME
			case STATUS_CONTINUE_FRAME:
				frameData := map[string]interface{}{
					"data": map[string]interface{}{
						"status":   STATUS_CONTINUE_FRAME,
						"format":   "audio/L16;rate=8000",
						"audio":    base64.StdEncoding.EncodeToString(buffer),
						"encoding": "raw",
					},
				}
				conn.WriteJSON(frameData)
			case STATUS_LAST_FRAME:
				frameData := map[string]interface{}{
					"data": map[string]interface{}{
						"status":   STATUS_LAST_FRAME,
						"format":   "audio/L16;rate=8000",
						"audio":    base64.StdEncoding.EncodeToString(buffer),
						"encoding": "raw",
					},
				}
				conn.WriteJSON(frameData)
				return
			}
			//模拟音频采样间隔
			time.Sleep(intervel)
		}
	}()

	var resultText strings.Builder
	//var decoder Decoder
	// 获取返回的数据
	for {
		var resp = RespData{}
		_, msg, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("read message error:", err)
			break
		}
		json.Unmarshal(msg, &resp)
		if resp.Code != 0 {
			fmt.Println(resp.Code, resp.Message, time.Since(st))

		}

		resultText.WriteString(resp.Data.Result.String())
		if resp.Data.Status == 2 {
			fmt.Println(resp.Code, resp.Message, time.Since(st))
			break
		}
	}
	return resultText.String()
}

type RespData struct {
	Sid     string `json:"sid"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    Data   `json:"data"`
}

type Data struct {
	Result Result `json:"result"`
	Status int    `json:"status"`
}

// 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl(hosturl string, apiKey, apiSecret string) string {
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
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sgin, apiSecret)
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

// 解析返回数据，仅供demo参考，实际场景可能与此不同。
type Decoder struct {
	results []*Result
}

func (d *Decoder) Decode(result *Result) {
	if len(d.results) <= result.Sn {
		d.results = append(d.results, make([]*Result, result.Sn-len(d.results)+1)...)
	}
	if result.Pgs == "rpl" {
		for i := result.Rg[0]; i <= result.Rg[1]; i++ {
			d.results[i] = nil
		}
	}
	d.results[result.Sn] = result
}

func (d *Decoder) String() string {
	var r string
	for _, v := range d.results {
		if v == nil {
			continue
		}
		r += v.String()
	}
	return r
}

type Result struct {
	Ls  bool   `json:"ls"`
	Rg  []int  `json:"rg"`
	Sn  int    `json:"sn"`
	Pgs string `json:"pgs"`
	Ws  []Ws   `json:"ws"`
}

func (t *Result) String() string {
	var wss string
	for _, v := range t.Ws {
		wss += v.String()
	}
	return wss
}

type Ws struct {
	Bg int  `json:"bg"`
	Cw []Cw `json:"cw"`
}

func (w *Ws) String() string {
	var wss string
	for _, v := range w.Cw {
		wss += v.W
	}
	return wss
}

type Cw struct {
	Sc int    `json:"sc"`
	W  string `json:"w"`
}
