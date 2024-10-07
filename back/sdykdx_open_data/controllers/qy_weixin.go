package controllers

import (
	"github.com/beego/beego/v2/server/web"
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
	code := c.GetString("code")

	if code == "" {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失，获取用户Code未提供",
			Data:    nil,
		}
		return
	}

	userInfo, err := models_qy_weixin.ObtainUserInfoByCode(code)

	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     err.Error(),
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

// SendQyWeixinMessage  发送企业微信消息
func (c *QyWeixinController) SendQyWeixinMessage() {
	defer func() {
		_ = c.ServeJSON()
	}()
	numbers := c.GetStrings("numbers")
	description := c.GetString("description")
	title := c.GetString("title")
	url := c.GetString("url")

	if description == "" || title == "" || url == "" || numbers == nil || len(numbers) == 0 {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失",
			Data:    nil,
		}
		return
	}

	err := models_qy_weixin.SendQyWeixinNewsMessage(numbers, title, description, url)

	if err != nil {
		c.Data["json"] = Response{
			Success: false,
			Msg:     err.Error(),
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
