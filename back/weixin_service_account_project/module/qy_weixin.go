package module

import (
	"encoding/json"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"io/ioutil"
	"net/http"
)

type QyWeixinController struct {
	beego.Controller
}

func (c *QyWeixinController) Test() {
	//fmt.Println(c.Ctx.Request)

	//code := c.GetString("code")
	//
	//accessToken, err := getAccessToken()
	//if err != nil {
	//	log.Fatalf("Failed to get access token: %v", err)
	//}
	//
	//userID, err := getUserInfo(accessToken, code)
	//if err != nil {
	//	log.Fatalf("Failed to get user info: %v", err)
	//}
	//
	//userDetails, err := getUserDetails(accessToken, userID)
	//if err != nil {
	//	log.Fatalf("Failed to get user details: %v", err)
	//}

	//fmt.Printf("User Details: %+v\n", userDetails)
	c.Ctx.WriteString("it's ok")
}

const (
	corpID     = "wxa059996e5d72516b"
	corpSecret = "Vpb4xpCQoT6Zq4KwaBgMYeflWPHlyIoyEw3pMOaeqyI"
)

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type UserInfoResponse struct {
	UserId   string `json:"UserId"`
	DeviceId string `json:"DeviceId"`
}

type UserDetailResponse struct {
	UserId   string `json:"userid"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
	Position string `json:"position"`
	// 添加更多你需要的字段
}

func getAccessToken() (string, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpID, corpSecret)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var accessTokenResponse AccessTokenResponse
	if err := json.Unmarshal(body, &accessTokenResponse); err != nil {
		return "", err
	}

	return accessTokenResponse.AccessToken, nil
}

func getUserInfo(accessToken, code string) (string, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s", accessToken, code)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var userInfoResponse UserInfoResponse
	if err := json.Unmarshal(body, &userInfoResponse); err != nil {
		return "", err
	}

	return userInfoResponse.UserId, nil
}

func getUserDetails(accessToken, userID string) (*UserDetailResponse, error) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s", accessToken, userID)

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var userDetailResponse UserDetailResponse
	if err := json.Unmarshal(body, &userDetailResponse); err != nil {
		return nil, err
	}

	return &userDetailResponse, nil
}
