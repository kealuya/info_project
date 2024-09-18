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

// Organization 组织架构信息
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

// GraduateStudent Organization 研究生基本信息
func (c *BusinessController) GraduateStudent() {
	defer func() {
		_ = c.ServeJSON()
	}()
	xh := c.GetString("xh")

	if xh == "" {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失",
			Data:    nil,
		}
		return
	}

	graduateStudent, err := models.ObtainGraduateStudent(xh)

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
		Data:    graduateStudent,
	}
	return
}

// ClinicalTeacher 临床教师信息
func (c *BusinessController) ClinicalTeacher() {
	defer func() {
		_ = c.ServeJSON()
	}()
	zgh := c.GetString("zgh")

	if zgh == "" {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失",
			Data:    nil,
		}
		return
	}

	clinicalTeacher, err := models.ObtainClinicalTeacher(zgh)

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
		Data:    clinicalTeacher,
	}
	return
}

// Faculty 教职工信息
func (c *BusinessController) Faculty() {
	defer func() {
		_ = c.ServeJSON()
	}()
	zgh := c.GetString("zgh")

	if zgh == "" {
		c.Data["json"] = Response{
			Success: false,
			Msg:     "参数缺失",
			Data:    nil,
		}
		return
	}

	faculty, err := models.ObtainFaculty(zgh)

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
		Data:    faculty,
	}
	return
}
