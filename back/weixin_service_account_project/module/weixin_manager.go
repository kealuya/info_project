package module

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	"io"
	"net/http"
	"sort"
	"strings"
	"time"
)

type ManagerController struct {
	beego.Controller
}

var ACCESSTOKEN = ""

func init() {
	go RefreshToken()
}

func GetAccessToken() string {
	return ACCESSTOKEN
}
func GetAccessTokenNow() error {
	APPID, _ := config.String("APPID")
	APPSECRET, _ := config.String("APPSECRET")
	logs.Info(fmt.Sprintf("AppId: %s, AppSecret: %s", APPID, APPSECRET))
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
		logs.Info(fmt.Sprintf("Access Token: %s", ACCESSTOKEN))
		return nil
	}

	return fmt.Errorf("failed to get access token: %v", result)
}

// RefreshToken   定时器每 1.5 小时自动刷新一次 token
func RefreshToken() {
	logs.Info("RefreshToken goroutine start")
	err := GetAccessTokenNow()
	if err != nil {
		logs.Error("RefreshToken::", err)
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

// VerifyToken is used to verify the server address is valid
func (c *ManagerController) VerifyToken() {

	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	TOKEN, _ := config.String("TOKEN")

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
