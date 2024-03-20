package models

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/nfnt/resize"
	"image"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"strings"
	"szyx_back/common"
	"szyx_back/entity/system"
)

/**
获取网报token
*/
func GetTokenWb(yhid string, qyid string) (token string) {
	Key := "ZhCZ6GoP7GbOX2zY"
	respBody := common.DoHttpPost_WbToken(common.GetToken + "?userId=" + yhid + "&qyId=" + qyid + "&key=" + Key)
	bytes := []byte(respBody)
	res := new(system.GetTokenWbResp)
	common.Unmarshal(bytes, &res)
	return res.Data.Token
}

/**
获取网报发票夹列表
*/
func GetInvoiceWbList(orderList *system.InvoiceWbList) (list system.ReturnInvoiceList, err error) {
	var fpUrl = common.InvoiceListUrl + "?qyId=" + orderList.CorpCode + "&userId=" + orderList.EmpCode +
		"&row=" + orderList.Row + "&page=" + orderList.Page +
		"&zt=" + orderList.Zt + "&type=" + orderList.Type
	respBody := common.DoHttpGet(fpUrl, "")
	bytes := []byte(respBody)
	logs.Info("获取网报发票夹列表返回：" + string(bytes))
	common.Unmarshal(bytes, &list)
	return list, err
}

/**
修改网报发票夹内发票的状态
*/
func UpdateInvoiceStates(orderList *system.InvoiceStates) (list system.ReturnOrderList, err error) {
	param := common.Marshal(orderList)
	respBody := common.DoHttpPost(common.InvoiceStates, string(param))
	bytes := []byte(respBody)
	logs.Info("修改网报发票夹内发票的状态返回：" + string(bytes))
	common.Unmarshal(bytes, &list)
	return list, err
}

/**
根据唯一id 删除发票、释放发票的占有
*/
func DeleteInvoiceToWb(req *system.InvoiceWbList) (res string, err error) {
	//发票识别需要先获取token
	token := GetTokenWb(req.EmpCode, req.CorpCode)
	respBody := common.DoHttpPost_WbFp_Delete(common.DeleteInvoiceToWb, req.CorpCode, req.Dzfph, token)
	logs.Info("根据唯一id 删除发票返回:" + respBody)
	return respBody, err
}

/**
获取网报发票夹单条数据，根据唯一id
*/
func GetInvoiceWbById(orderList *system.InvoiceWbList) (list system.ReturnInvoiceByIdList, err error) {
	token := GetTokenWb(orderList.EmpCode, orderList.CorpCode)
	var fpUrl = common.InvoiceByIDUrl
	respBody := common.DoHttpGet_WbFpsb_(fpUrl, token, orderList.Dzfph, orderList.CorpCode)
	bytes := []byte(respBody)
	common.Unmarshal(bytes, &list)
	return list, err
}

/**
获取到网报pdf文件地址，转换成OBS图片地址
*/
func PDFTojpgObs(pdfTojpgObs *system.PDFTojpgObs) (pdfTojpgObsReq *system.PDFTojpgObsReq, err error) {
	param := common.Marshal(pdfTojpgObs)
	respBody := common.DoHttpPost(common.PDFTojpgObsUrl, string(param))
	bytes := []byte(respBody)
	common.Unmarshal(bytes, &pdfTojpgObsReq)
	fmt.Println("获取到网报pdf文件地址，转换成OBS图片地址。中间件的返回参数：" + string(respBody))
	return pdfTojpgObsReq, err
}

/**
网报上传图片（base64手动录入发票信息）
*/
func UploadBase64(req *system.UploadBase64Request) (res string, err error) {
	//网报上传图片 需要先获取token
	token := GetTokenWb(req.EmpCode, req.QyId)
	//读取文件
	path := fmt.Sprint(req.ImgUrl)
	resBody, err := http.Get(path)
	if err != nil {
		common.ErrorHandler(err, `获取图片文件失败::%+v`)
	}
	defer resBody.Body.Close()
	data, err := ioutil.ReadAll(resBody.Body)
	common.ErrorHandler(err, `读取图片文件失败::%+v`)
	//将图片转成base64
	imgBase64 := base64.StdEncoding.EncodeToString(data)
	index := strings.Index(path, "com/")
	filename := path[index+4:]
	//前缀
	var base64Encoding string
	//获取文件类型，拼接base64前缀
	mimeType := http.DetectContentType(data)
	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	case "image/jpg":
		base64Encoding = "data:image/jpg;base64,"
	}
	imgBase64_ := base64Encoding + imgBase64
	respBody := common.DoHttpPost_WbFpsb_UploadBase64(common.UploadBase64, filename, req.QyId, imgBase64_, token)
	logs.Info("网报上传图片（base64手动录入发票信息）返回:" + respBody)
	return respBody, err
}

/**
网报获取手动输入发票类型列表
*/
func GetFpInputTypeList(req *system.GetFpInputTypeListRequest) (res string, err error) {
	//网报获取手动输入发票类型列表 需要先获取token
	token := GetTokenWb(req.EmpCode, req.QyId)
	//拼接请求地址
	reqUrl := common.GetFpInputTypeList + "?qyId=" + req.QyId
	respBody := common.DoHttpGet_WbFpsb(reqUrl, token)
	logs.Info("网报获取手动输入发票类型列表信息）返回:" + respBody)
	return respBody, err
}

/**
网报发票上传保存
*/
func InvVerifyWb(req *system.InvVerifyRequest) (res system.InvVerifyResponse, err error) {
	//判断是否有obs图片路径  如果有就先拿到图片转成base64编码
	if req.FilePath != "" {
		//取得图片的的OBS地址，进行调用然后转成64编码
		res1, err1 := http.Get(req.FilePath)
		if err1 != nil {
			common.ErrorHandler(err, `解析responseBodyJson错误::%+v`)
		}
		defer res1.Body.Close()
		data, _ := ioutil.ReadAll(res1.Body)
		//获取文件类型，如果是pdf的不能进行压缩 ，因为pdf文件不支持压缩会报错
		mimeType := http.DetectContentType(data)
		if mimeType == "application/pdf" {
			str := base64.StdEncoding.EncodeToString(data)
			req.FileBase64 = str
		} else {
			//这个应该是按大小压缩，宽度跟高度设定
			var width uint = 1000
			var height uint = 800
			//文件压缩
			decodeBuf, layout, _ := image.Decode(bytes.NewReader(data))
			// 修改图片的大小
			set := resize.Resize(width, height, decodeBuf, resize.Lanczos3)
			NewBuf := bytes.Buffer{}
			switch layout {
			case "png":
				err = png.Encode(&NewBuf, set)
			case "jpeg", "jpg":
				err = jpeg.Encode(&NewBuf, set, &jpeg.Options{Quality: 80})
			}
			if NewBuf.Len() < len(data) {
				data = NewBuf.Bytes()
			}
			str := base64.StdEncoding.EncodeToString(data)
			req.FileBase64 = str
		}
	}
	param := common.Marshal(req)
	//发票识别需要先获取token
	token := GetTokenWb(req.EmpCode, req.QyId)
	respBody := common.DoHttpPost_WbFpsb(common.InvVerifyWb, string(param), token)
	logs.Info("网报发票上传保存返回:" + respBody)
	bytes := []byte(respBody)
	common.Unmarshal(bytes, &res)
	return res, err
}

/**
网报发票图片识别
*/
func GetInvOcrResultWb(req *system.InvOcrRequest) (res string, err error) {
	//取得图片的base64编码，进行调用然后转成64编码
	res1, err1 := http.Get(req.FilePath)
	if err1 != nil {
		common.ErrorHandler(err, `解析responseBodyJson错误::%+v`)
	}
	defer res1.Body.Close()
	data, _ := ioutil.ReadAll(res1.Body)

	//因为这个压缩会使识别图片时间变短，但是识别的正确率很低 所以不压缩 改为900会好点
	//这个应该是按大小压缩，宽度跟高度设定
	var width uint = 900
	var height uint = 900
	//文件压缩
	decodeBuf, layout, err := image.Decode(bytes.NewReader(data))
	// 修改图片的大小
	set := resize.Resize(width, height, decodeBuf, resize.Lanczos3)
	NewBuf := bytes.Buffer{}
	switch layout {
	case "png":
		err = png.Encode(&NewBuf, set)
	case "jpeg", "jpg":
		err = jpeg.Encode(&NewBuf, set, &jpeg.Options{Quality: 80})
	}
	if NewBuf.Len() < len(data) {
		data = NewBuf.Bytes()
	}

	//前缀
	var base64Encoding string
	//获取文件类型，拼接base64前缀
	mimeType := http.DetectContentType(data)
	switch mimeType {
	case "image/jpeg":
		base64Encoding = "data:image/jpeg;base64,"
	case "image/png":
		base64Encoding = "data:image/png;base64,"
	case "image/jpg":
		base64Encoding = "data:image/jpg;base64,"
	}
	//fmt.Print(base64Encoding)
	//组装后的base64编码
	imgBase64_ := base64Encoding + base64.StdEncoding.EncodeToString(data)
	//发票识别需要先获取token
	token := GetTokenWb(req.EmpCode, req.QyId)
	req_wb := new(system.InvOcrRequest_Wb)
	req_wb.OcrInvType = req.OcrInvType
	req_wb.IsBase64 = req.IsBase64
	req_wb.Base64 = imgBase64_
	req_wb.QyId = req.QyId
	param := common.Marshal(req_wb)
	respBody := common.DoHttpPost_WbFpsb(common.GetInvOcrResultWb, string(param), token)
	logs.Info("网报发票图片识别返回:" + respBody)
	return respBody, err
}

/**
网报发票pdf,ofd识别,手机端用【前端上传文件到obs】
*/
func GetInvOcrResultByPdfWb(req *system.InvOcrByPdfRequest) (res string, err error) {
	token := GetTokenWb(req.EmpCode, req.QyId)
	respBody := common.DoHttpPost_WbFpsb_Pdf(common.GetInvOcrResultByPdfWeb, req.QyId, req.PdfUrl, token)
	logs.Info("网报发票pdf文件解析识别返回：" + string(respBody))
	return respBody, err
}

/**
获取微信Access_tokenTicket   暂时不用，之前差旅管家调用微信卡包用的
*/
/*func GetAccess_tokenTicket(requestHeader string) (res system.Access_tokenTicket, err error) {
	req2 := httplib.Post("http://192.168.120.220:10003/v1/invoice/getAccess_tokenTicket")
	req2.Param("requestHeader", requestHeader)
	respBody, _ := req2.String()
	//处理json去除转义字符
	handleJsonString := strings.ReplaceAll(respBody, `\"`, `"`)
	handleJsonString = handleJsonString[1 : len(handleJsonString)-1]
	fmt.Println(handleJsonString)
	bytes := []byte(handleJsonString)
	common.Unmarshal(bytes, &res)
	logs.Info("获取微信Access_tokenTicket返回：" + respBody)
	return res, err
}*/
