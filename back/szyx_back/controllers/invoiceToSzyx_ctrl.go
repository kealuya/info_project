package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
	"szyx_back/common"
	"szyx_back/entity/system"
	"szyx_back/models"
)

type InvoiceToSzyxApi struct {
	beego.Controller
}

// @Title 获取网报发票夹list
// @Tags GetInvoiceWbList
// @Summary 获取网报发票夹getURL
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvoiceWbList true "企业编码 ,申请人Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @router /getInvoiceWbList [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) GetInvoiceWbList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invoiceWbList := new(system.InvoiceWbList)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &invoiceWbList)
	logs.Info("获取网报发票夹列表入参：" + string(jsonByte))
	//业务处理
	invoiceWbLists, err := models.GetInvoiceWbList(invoiceWbList)
	if err == nil {
		resJson.Success = true
		resJson.Data = invoiceWbLists
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("获取网报发票夹list失败::%s", err)
	}
}

// @Title 根据唯一id 删除发票、释放发票的占有
// @Tags DeleteInvoiceToWb
// @Summary 根据唯一id 删除发票、释放发票的占有
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvoiceWbList true "企业编码 ,申请人Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @router /deleteInvoiceToWb [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) DeleteInvoiceToWb() {
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invoiceWbList := new(system.InvoiceWbList)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &invoiceWbList)
	logs.Info("根据唯一id 删除发票入参：" + string(jsonByte))
	//业务处理
	respBody, err := models.DeleteInvoiceToWb(invoiceWbList)
	if err == nil {
		resJson.Success = true
		resJson.Data = respBody
	} else {
		resJson.Success = false
		resJson.Msg = "取网报发票夹单条数据，根据唯一id失败."
	}
}

// @Title 更新网报发票使用状态
// @Tags UpdateInvoiceStates
// @Summary 更新网报发票使用状态
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvoiceStates true "企业编码 ,申请人Id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @router /updateInvoiceStates [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) UpdateInvoiceStates() {
	//这个接口目前用不到 没有调用
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invoiceStates := new(system.InvoiceStates)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &invoiceStates)
	logs.Info("更新网报发票使用状态入参：" + string(jsonByte))
	//业务处理
	updateStates, err := models.UpdateInvoiceStates(invoiceStates)
	if err == nil {
		resJson.Success = true
		resJson.Data = updateStates
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("更新网报发票使用状态失败::%s", err)
	}
}

// @Title 获取网报发票图片地址
// @Tags GetInvoiceImgUrl
// @Summary 获取网报发票图片地址
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvoiceWbList true "企业编码 ,申请人Id."
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @router /getInvoiceImgUrl [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) GetInvoiceImgUrl() {
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invoiceList := new(system.InvoiceWbList)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	logs.Info("获取网报发票图片地址入参：" + string(jsonByte))
	common.Unmarshal(jsonByte, &invoiceList)
	//业务处理 取的图片url
	var imgUrl = ""
	imgUrl = common.InvoiceImgUrl + "?uuid=" + invoiceList.Uuid + "&qyId=" + invoiceList.CorpCode
	//业务处理 取得pdf图片url
	invoiceWbLists, err := models.GetInvoiceWbById(invoiceList)
	//判断接口报错机制
	if err == nil && invoiceWbLists.Data != nil {
		var pdfUrl = invoiceWbLists.Data[0].Filepath
		//判断文件地址是否为PDFURL地址
		if strings.HasSuffix(pdfUrl, "pdf") {
			//获取到网报pdf文件地址，转换成OBS图片地址
			pdfTojpgObs := new(system.PDFTojpgObs)
			pdfTojpgObs.PDFUrl = pdfUrl
			pdfTojpgObsReq, err := models.PDFTojpgObs(pdfTojpgObs)
			if err == nil {
				if pdfTojpgObsReq.Code == "200" {
					imgUrl = pdfTojpgObsReq.Data
				} else {
					imgUrl = "端口10009中间件报错：" + pdfTojpgObsReq.Msg + pdfTojpgObsReq.Code
				}
			} else {
				imgUrl = "端口10009中间件pdfTojpgObs转换文件错误~请联系管理员"
			}
		}
	}
	resJson.Data = imgUrl
}

// @Title 网报上传图片（手动录入发票） 【手机端使用】
// @Tags UploadBase64
// @Summary 网报上传图片（手动录入发票
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.UploadBase64Request true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"图片上传成功"}"
// @router /uploadBase64 [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) UploadBase64() {
	//网报上传图片（手动录入发票）
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	uploadBase64Request := new(system.UploadBase64Request)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	logs.Info("网报上传图片（手动录入发票）入参：" + string(jsonByte))
	common.Unmarshal(jsonByte, &uploadBase64Request)
	//业务处理
	uploadBase64Res, err := models.UploadBase64(uploadBase64Request)
	if err == nil {
		resJson.Success = true
		resJson.Data = uploadBase64Res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("网报手动上传图片错误::%s", err)
	}
}

// @Title 网报获取手动输入发票类型列表
// @Tags GetFpInputTypeList
// @Summary 网报获取手动输入发票类型列表
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.GetFpInputTypeListRequest true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取手动录入发票类型列表成功"}"
// @router /getFpInputTypeList [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) GetFpInputTypeList() {
	//获取手动输入发票类型列表
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	getFpInputTypeListRequest := new(system.GetFpInputTypeListRequest)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	logs.Info("网报获取手动输入发票类型列表入参：" + string(jsonByte))
	common.Unmarshal(jsonByte, &getFpInputTypeListRequest)
	//业务处理
	getFpInputTypeListRes, err := models.GetFpInputTypeList(getFpInputTypeListRequest)
	if err == nil {
		resJson.Success = true
		resJson.Data = getFpInputTypeListRes
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("网报获取手动输入发票类型列表::%s", err)
	}
}

// @Title 网报发票上传保存
// @Tags InvVerifyWb
// @Summary 网报发票上传保存
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvVerifyRequest true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"发票保存成功"}"
// @router /invVerifyWb [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) InvVerifyWb() {
	//保存发票到发票夹
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invVerifyRequest := new(system.InvVerifyRequest)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &invVerifyRequest)
	logs.Info("网报发票上传保存入参：" + string(jsonByte))
	//业务处理
	invVerifyRes, err := models.InvVerifyWb(invVerifyRequest)
	if err == nil {
		resJson.Success = true
		resJson.Data = invVerifyRes
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("上传发票失败::%s", err)
	}
}

// @Title 网报发票图片解析识别
// @Tags GetInvOcrResultWb
// @Summary 网报发票图片解析
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvOcrRequest true "图片base64 ,是否base64固定值传1,固定值300,企业编号"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"图片识别成功"}"
// @router /getInvOcrResultWb [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) GetInvOcrResultWb() {
	//发票识别接口
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invOcrRequest := new(system.InvOcrRequest)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &invOcrRequest)
	invOcrRequest.IsBase64 = "1"     //将参数置为固定值
	invOcrRequest.OcrInvType = "300" //固定值300

	logs.Info("网报发票图片解析识别入参：" + string(jsonByte))
	//业务处理
	invOcrResultWb, err := models.GetInvOcrResultWb(invOcrRequest)
	if err == nil {
		resJson.Success = true
		//无报错信息，直接返回网报的返回值
		resJson.Data = invOcrResultWb
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("发票识别接口调用失败::%s", err)
	}
}

// @Title 网报发票pdf文件解析识别
// @Tags GetInvOcrResultByPdfWb
// @Summary 网报发票pdf与ofd解析
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvOcrByPdfRequest true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"文件识别成功"}"
// @router /getInvOcrResultByPdfWb [post]
func (invoiceToSzyxCtrl *InvoiceToSzyxApi) GetInvOcrResultByPdfWb() {
	//发票识别接口
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	invOcrByPdfRequest := new(system.InvOcrByPdfRequest)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &invOcrByPdfRequest)
	invOcrByPdfRequest.IsBase64 = false //将参数置为固定值
	logs.Info("网报发票pdf文件解析识别入参：" + string(jsonByte))
	//业务处理
	invOcrResultWb, err := models.GetInvOcrResultByPdfWb(invOcrByPdfRequest)
	if err == nil {
		resJson.Success = true
		//无报错信息，直接返回网报的返回值
		resJson.Data = invOcrResultWb
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("发票识别接口调用失败::%s", err)
	}
}

// @Title 获取微信Access_tokenTicket
// @Tags GetAccess_tokenTicket
// @Summary 获取微信Access_tokenTicket
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body system.InvOcrByPdfRequest true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取微信Access_tokenTicket成功"}"
// @router /getAccess_tokenTicket [post]
/*func (invoiceToSzyxCtrl *InvoiceToSzyxApi) GetAccess_tokenTicket() {
	//发票识别接口
	resJson := NewJsonStruct(nil)
	defer func() {
		invoiceToSzyxCtrl.Data["json"] = string(common.Marshal(resJson))
		invoiceToSzyxCtrl.ServeJSON()
	}()
	access_tokenTicket := new(system.Access_tokenTicket)
	var jsonByte = invoiceToSzyxCtrl.Ctx.Input.RequestBody
	common.Unmarshal(jsonByte, &access_tokenTicket)
	fmt.Println(access_tokenTicket.TokenTicketUrl)
	if(access_tokenTicket.TokenTicketUrl != ""){
		invOcrResultWb, err := models.GetAccess_tokenTicket(access_tokenTicket.TokenTicketUrl)
		//invOcrResultWb, err := models.GetAccess_tokenTicket("43g748b172.imdo.co/clgjInvoice")
		if err == nil {
			resJson.Success = true
			//无报错信息，直接返回返回值
			resJson.Data = invOcrResultWb
		} else {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("获取微信Access_tokenTicket失败::%s", err)
		}
	}else{
		resJson.Success = false
		resJson.Msg = "获取Access_tokenTicket的url参数为空，请检查请求参数！"
	}
}*/
