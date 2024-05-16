package common

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"io"
	"runtime/debug"
	"strings"
)

//华为obs基础配置信息
const (
	AccessBaseUrl = "https://jbkj.obs.cn-east-3.myhuaweicloud.com/"
	Endpoint      = "obs.cn-east-3.myhuaweicloud.com"
	Ak            = "9O6QEHVJRXZDXBQ0IKTZ"
	Sk            = "HhUkyO89XLiXwt3tyJlm0uCiP36zzqn4EvuG6GRU"
	BucketName    = "jbkj"
)

//返回响应参数组装类
type ObsPutFileResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	ImgURL  string `json:"imgURL"`
}

/**
华为obs上传文件
*/
func ObsPutFile(corpCode string, fileName string, file io.Reader,folder string,) (obsPutFileRep ObsPutFileResponse, resultErr error) {

	defer func() {
		if err := recover(); err != nil {
			logs.Error("★★★★错误recover::", err)
			logs.Error(string(debug.Stack()))
			obsPutFileRep.Success = false
			obsPutFileRep.Msg = "对象化存储接口发生错误!"
			resultErr = fmt.Errorf("对象化存储接口发生错误：%+v", err)
		}
	}()
	// 创建ObsClient结构体
	var obsClient, err_obsNew = obs.New(Ak, Sk, Endpoint)
	ErrorHandler(err_obsNew, "获取obsClient发生错误:%+v")
	// 关闭obsClient
	defer obsClient.Close()
	//取得上传文件的文件类型后缀
	fileType := fileName[strings.LastIndexAny(fileName, ".")+1:]
	if fileType != "pdf" && fileType != "jpeg" && fileType != "xls" && fileType != "xlsx" && fileType != "doc" && fileType != "docx" && fileType != "png" && fileType != "jpg" && fileType != "mp3" {
		ErrorHandler(errors.New("错误文件类型"), "通过url获取对象类型发生错误:%+v")
		obsPutFileRep.Success = false
		obsPutFileRep.Msg = "对象化存储接口发生错误!- 错误文件类型"
		return obsPutFileRep, nil
	} else {
		input_object := &obs.PutObjectInput{}
		// 指定存储桶名称
		input_object.Bucket = BucketName
		//tc := t.New(time.Now().UnixNano())
		// 对象化存储文件生成
		on := "data/"+ folder +"/" + corpCode + "_" + fileName
		input_object.Key = on
		input_object.Body = file
		output, err_obsClientPutObject := obsClient.PutObject(input_object)
		ErrorHandler(err_obsClientPutObject, "进行对象化存储发生错误:%+v")
		if output.StatusCode != 200 {
			obsPutFileRep.Success = false
			obsPutFileRep.Msg = "对象化存储接口发生错误!"
			if obsError, ok := err_obsClientPutObject.(obs.ObsError); ok {
				ErrorHandler(obsError)
			} else {
				ErrorHandler(err_obsClientPutObject, "进行对象化存储发生错误:%+v")
			}
		} else {
			obsPutFileRep.Success = true
			obsPutFileRep.ImgURL = AccessBaseUrl + on
			obsPutFileRep.Msg = "obs对象化存储上传成功！"
		}
		return obsPutFileRep, resultErr
	}
}
