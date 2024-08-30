package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"sdykdx_open_data/models"
)

type BusinessController struct {
	web.Controller
}
type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Data    any    `json:"data"`
}

func (c *BusinessController) Organization() {
	defer func() {
		_ = c.ServeJSON()
	}()
	dwdm := c.GetString("dwdm")

	if dwdm == "" {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失",
			Data:    nil,
		}
		return
	}

	organization, err := models.ObtainOrganization(dwdm)

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
		Data:    organization,
	}
	return
}
