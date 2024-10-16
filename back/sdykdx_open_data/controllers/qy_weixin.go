package controllers

import (
	"encoding/json"
	"github.com/beego/beego/v2/server/web"
	"io"
	"sdykdx_open_data/internal/qy_weixin"
	models_qy_weixin "sdykdx_open_data/models/qy_weixin"
)

type QyWeixinController struct {
	web.Controller
}

// GetAppId  获取企业AppId
func (c *QyWeixinController) GetAppId() {

	m := make(map[string]any)
	m["appId"] = qy_weixin.AppId
	c.Data["json"] = Response{
		Success: true,
		Data:    m,
	}
	_ = c.ServeJSON()
	return
}

// GetUserByCode 通过Code获取企业微信用户信息
func (c *QyWeixinController) GetUserByCode() {
	defer func() {
		_ = c.ServeJSON()
	}()

	// 读取请求体
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "无法读取请求体",
			Data:    nil,
		}
		return
	}

	// 解析 JSON 数据
	var payload map[string]any
	if err := json.Unmarshal(body, &payload); err != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "请求体解析错误",
			Data:    nil,
		}
		return
	}

	code, ok := payload["code"]
	if !ok || code == "" {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失，获取用户Code未提供",
			Data:    nil,
		}
		return
	}

	userInfo, errObtainUserInfoByCode := models_qy_weixin.ObtainUserInfoByCode(code.(string))

	if errObtainUserInfoByCode != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     errObtainUserInfoByCode.Error(),
			Data:    nil,
		}
		return
	}
	c.Data["json"] = Response{
		Success: true,
		Data:    userInfo,
	}
	return
}

type RequestPayload struct {
	Description string   `json:"description"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
	Numbers     []string `json:"numbers"`
}

// SendQyWeixinMessage  发送企业微信消息
func (c *QyWeixinController) SendQyWeixinMessage() {
	defer func() {
		_ = c.ServeJSON()
	}()
	// 读取请求体
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "无法读取请求体",
			Data:    nil,
		}
		return
	}

	// 解析 JSON 数据
	var payload RequestPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "请求体解析错误",
			Data:    nil,
		}
		return
	}

	// 检查参数是否完整
	if payload.Description == "" || payload.Title == "" || payload.URL == "" || len(payload.Numbers) == 0 {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失",
			Data:    nil,
		}
		return
	}

	errSendQyWeixinNewsMessage := models_qy_weixin.SendQyWeixinNewsMessage(payload.Numbers, payload.Title, payload.Description, payload.URL)

	if errSendQyWeixinNewsMessage != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     errSendQyWeixinNewsMessage.Error(),
			Data:    nil,
		}
		return
	}
	c.Data["json"] = Response{
		Success: true,
		Data:    nil,
	}
	return
}
