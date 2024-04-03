package common

import (
	"github.com/astaxie/beego/httplib"
	"strconv"
	"strings"
)

//获取网报发票夹getURL接口
var InvoiceListUrl = "http://XXXXXXXXX/wsyy/invoices"

//网报发票夹识别获取token
var GetToken = "http://XXXXXXXXX/wsyy/authLogin"

//网报发票夹删除发票
var DeleteInvoiceToWb = "http://XXXXXXXXX/invManage/deleteInv"

//获取网报发票图片地址接口
var InvoiceImgUrl = "http://XXXXXXXXX/invocr/loadImage"

//获取单条网报发票getURL接口
var InvoiceByIDUrl = "http://XXXXXXXXX/invManage/invDetails"

//PDF文件地址转OBS中间件地址
var PDFTojpgObsUrl = "http://XXXXXXXXX/v1/openApi/pdfTojpgObs"

//网报上传图片base64
var UploadBase64 = "http://XXXXXXXXX/file/uploadBase64"

//网报获取手动输入发票类型列表
var GetFpInputTypeList = "http://XXXXXXXXX/invManage/getFpInputTypeList"

//网报发票保存
var InvVerifyWb = "http://XXXXXXXXX/invManage/invVerify"

//更新网报发票使用状态接口
var InvoiceStates = "http://XXXXXXXXX/wsyy/updateDzfpgl"

//网报发票夹图片识别
//var GetInvOcrResultWb = "http://XXXXXXXXX/invocr/getInvOcrResult"
var GetInvOcrResultWb = "http://XXXXXXXXX/invUnifiedOcr/getInvUnifiedOcrResult"

//网报发票夹pdf网络文件识别
var GetInvOcrResultByPdfWeb = "http://XXXXXXXXX/invUnifiedOcr/getInvUnifiedOcrResultByWebPdf"

//http Post 请求
func DoHttpPost(url string, param string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Post 请求 网报发票识别
func DoHttpPost_WbFpsb(url string, param string, token string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Header("token", token)
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Post 请求 网报发票识别
func DoHttpPost_WbFpsb_Pdf(pdf_url string, qyid string, pdfUrl string, token string, options ...map[string]string) (respBody string) {
	req := httplib.Post(pdf_url)
	req.Header("Content-Type", "form-data")
	req.Header("token", token)

	req.Param("qyId", qyid)
	req.Param("isBase64", "false") //固定值
	req.Param("isPdf", "true")     //固定值
	req.Param("ocrInvType", "")    //固定值
	req.Param("pdfUrl", pdfUrl)
	req.Debug(true)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Get 请求网报
func DoHttpGet_WbFpsb(url string, token string, options ...map[string]string) (respBody string) {
	req := httplib.Get(url)
	req.Header("Content-Type", "application/json")
	req.Header("token", token)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Post 请求 网报上传图片
func DoHttpPost_WbFpsb_UploadBase64(url string, fileName string, qyid string, imgBase64 string, token string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	req.Header("Content-Type", "form-data")
	req.Header("token", token)

	req.Param("qyId", qyid)
	req.Param("file", imgBase64)
	req.Param("fileName", fileName)
	req.Param("fileType", "fptype")  //固定值
	req.Param("businessType", "113") //固定值
	req.PostFile("file", "")
	req.Debug(true)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Get 请求网报
func DoHttpGet_WbFpsb_(url1 string, token string, que string, que2 string, options ...map[string]string) (respBody string) {
	req := httplib.Get(url1)
	req.Header("Content-Type", "application/json;charset=UTF-8")
	req.Header("token", token)
	req.Param("invCodes", que)
	req.Param("qyId", que2)
	req.Debug(true)
	respBody, err := req.String()

	ErrorHandler(err)
	return respBody
}

//http Post 请求 网报发票识别 删除接口
func DoHttpPost_WbFp_Delete(pdf_url string, qyid string, invCode string, token string, options ...map[string]string) (respBody string) {
	req := httplib.Post(pdf_url)
	req.Header("Content-Type", "form-data")
	req.Header("token", token)
	req.Param("qyId", qyid)
	req.Param("invCode", invCode)
	req.Debug(true)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Post 请求 网报获取token
func DoHttpPost_WbToken(url string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//unicode转中文
func UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(
		strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

//http Get 请求
func DoHttpGet(url string, param string, options ...map[string]string) (respBody string) {
	req := httplib.Get(url)
	req.Header("Content-Type", "application/json")
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

//http Post 请求 差旅管家授权专用
func DoHttpPost_CLGJ(url string, sq string, param string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	req.Header("Content-Type", "application/json")
	req.Header("Authorization", sq)
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}

func DoHttpPost_CLYH(url string, param string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)

	return Base64ToString(respBody)
}

//http Post 请求
func DoHttpPost_Content_text(url string, param string, options ...map[string]string) (respBody string) {
	req := httplib.Post(url)
	req.Header("Content-Type", "text/plain")
	req.Body(param)
	respBody, err := req.String()
	ErrorHandler(err)
	return respBody
}
