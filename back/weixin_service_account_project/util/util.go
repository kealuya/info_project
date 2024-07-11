package util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var TOKEN = "renhao666"                                    // Your token
var APPID = "wx5276498b4494a34a"                           // Your AppID
var AESKEY = "WgM6RPo5WvVjSP3BFeEvuoMntlq6qYgIy2rroFACl12" // Your EncodingAESKey
var APPSECRET = "e5946c74933fee21389bc22164eb4600"
var ACCESSTOKEN = ""

func init() {
	go RefreshToken()
}

func GetAccessToken() string {
	return ACCESSTOKEN
}
func GetAccessTokenNow() error {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", APPID, APPSECRET)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return err
	}

	if val, ok := result["access_token"]; ok {
		ACCESSTOKEN = val.(string)
		return nil
	}

	return fmt.Errorf("failed to get access token: %v", result)
}

// RefreshToken   定时器每 1.5 小时自动刷新一次 token
func RefreshToken() {
	err := GetAccessTokenNow()
	if err != nil {
		return
	}
	ticker := time.NewTicker(90 * time.Minute) // 每 1.5 小时
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			err := GetAccessTokenNow()
			if err != nil {
				return
			}
			if err != nil {
				fmt.Printf("Error getting access token: %v\n", err)
				continue
			}

			fmt.Printf("Access token refreshed: %s\n", ACCESSTOKEN)
		}
	}
}