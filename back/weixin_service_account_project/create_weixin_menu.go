package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	. "weixin_service_account_project/util"
)

type MenuButton struct {
	Type      string       `json:"type,omitempty"`
	Name      string       `json:"name"`
	Key       string       `json:"key,omitempty"`
	URL       string       `json:"url,omitempty"`
	SubButton []MenuButton `json:"sub_button,omitempty"`
}

type Menu struct {
	Button []MenuButton `json:"button"`
}

func createMenu(accessToken string) error {
	url_weixin := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", accessToken)
	fmt.Println(accessToken)
	fmt.Println(url.QueryEscape(`https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx5276498b4494a34a&redirect_uri=https://zc.qcykj.com.cn/xx_asset_h5/#/judgePage&response_type=code&scope=snsapi_base#wechat_redirect`))

	menu := Menu{
		Button: []MenuButton{

			{
				Name: "业财服务",
				SubButton: []MenuButton{
					{
						Type: "view",
						Name: "浩天物流",
						URL:  "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx5276498b4494a34a&redirect_uri=https%3A%2F%2Ftech.szhtkj.com.cn%2Fwuliu%2Fstatic%2Findex.html%23%2FssoLogin&response_type=code&scope=snsapi_base#wechat_redirect",
					},
					{
						Type: "view",
						Name: "智慧资产",
						URL:  "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx5276498b4494a34a&redirect_uri=https%3A%2F%2Fzc.qcykj.com.cn%2Fxx_asset_h5%2F%23%2Flogin%3Fclient%3Dfwwx&response_type=code&scope=snsapi_base&state=STATE#wechat_redirect",
					},
				},
			},
			{
				Name: "账户",
				SubButton: []MenuButton{
					{
						Type: "view",
						Name: "用户绑定",
						URL:  "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx5276498b4494a34a&redirect_uri=https%3A%2F%2Fzc.qcykj.com.cn%2Fxx_asset_h5%2F%23%2FjudgePage&response_type=code&scope=snsapi_base#wechat_redirect",
					},
				},
			},
		},
	}

	/*
		在 JSON 编码中，符号 & 通常会被转义为 Unicode 表示形式 \u0026。这可能会导致你的 API 请求出错。为了强制使用原始的 & 字符，你可以手动替换转义后的字符。
		使用 json.Marshal 编码 JSON 数据时，不可避免的会遇到字符转义问题。默认情况下，Go 的 json.Marshal 会将某些字符（比如 &等）转换为对应的 Unicode 转义代码。
	*/
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(menu)
	if err != nil {
		return fmt.Errorf("error encoding menu data: %v", err)
	}

	//fmt.Println(buf.String()) // 调试输出修正后的 JSON 数据

	resp, err := http.Post(url_weixin, "application/json", strings.NewReader(buf.String()))
	if err != nil {
		return fmt.Errorf("error creating menu: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return fmt.Errorf("error unmarshaling response: %v", err)
	}

	if result["errcode"].(float64) != 0 {
		return fmt.Errorf("error creating menu: %v", result["errmsg"])
	}

	return nil
}

func main() {
	var err error
	err = GetAccessTokenNow()

	if err != nil {
		fmt.Printf("Error getting access token: %v\n", err)
		return
	}

	err = createMenu(ACCESSTOKEN)
	if err != nil {
		fmt.Printf("Error creating menu: %v\n", err)
		return
	}

	fmt.Println("Menu created successfully")
}
