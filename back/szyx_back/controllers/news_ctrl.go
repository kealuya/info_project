package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"path"
	"szyx_back/common"
	"szyx_back/entity/newsInfo"
	"szyx_back/models"
	"time"
)

type NewsCtrl struct {
	beego.Controller
}

// @Title 推广新闻发布列表list
// @Tags GetNewsTemplateList
// @Summary 推广新闻发布列表list
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.NewsInfo true "NewsInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getNewsTemplateList [post]
func (NewsCtrl *NewsCtrl) GetNewsTemplateList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	newsList_Param := new(newsInfo.NewsList_Param)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("推广新闻列表list入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &newsList_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetNewsTemplateList(newsList_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "推广新闻列表获取成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("查询推广新闻列表失败:", err)
	}
}

// @Title 获取推广新闻内容详情
// @Tags GetNewsDetail
// @Summary 获取推广新闻内容详情
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.NewsInfo true "Question struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getNewsDetail [post]
func (NewsCtrl *NewsCtrl) GetNewsDetail() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	question_Param := new(newsInfo.NewsInfo)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("推广新闻内容详情详情的入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &question_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetNewsDetail(question_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "推广新闻内容详情获取成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("推广新闻内容详情失败:", err)
	}
}

// @Title 新增推广新闻
// @Tags AddNews
// @Summary 新增推广新闻
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.NewsInfo true "NewsInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /addNews [post]
func (NewsCtrl *NewsCtrl) AddNews() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	newsInfo_Param := new(newsInfo.NewsInfo)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("新增推广新闻入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &newsInfo_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.AddNews(newsInfo_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "新增推广新闻成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("新增推广新闻失败:", err)
	}
}

// @Title 修改推广新闻
// @Tags ModifyNews
// @Summary 修改推广新闻
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.NewsInfo true "NewsInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /modifyNews [post]
func (NewsCtrl *NewsCtrl) ModifyNews() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	newsInfo_Param := new(newsInfo.NewsInfo)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("修改推广新闻入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &newsInfo_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.ModifyNews(newsInfo_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "修改推广新闻成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("修改推广新闻失败:", err)
	}
}

// @Title 删除推广新闻
// @Tags DeleteNews
// @Summary  删除推广新闻
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.NewsInfo true "NewsInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @router /deleteNews [post]
func (NewsCtrl *NewsCtrl) DeleteNews() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	newsInfo_Param := new(newsInfo.NewsInfo)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("删除新闻入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &newsInfo_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.DeleteNews(newsInfo_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "删除新闻成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("删除新闻失败:", err)
	}
}

// @Title 视频新闻发布列表list
// @Tags GetVideoTemplateList
// @Summary 视频新闻发布列表list
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.VideoInfo true "VideoInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @router /getVideoTemplateList [post]
func (NewsCtrl *NewsCtrl) GetVideoTemplateList() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	newsList_Param := new(newsInfo.VideoList_Param)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("视频新闻发布列表list入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &newsList_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	res, err := models.GetVideoTemplateList(newsList_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "视频新闻发布列表获取成功"
		resJson.Data = res
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("查询视频新闻发布列表失败:", err)
	}
}

// @Title 新增视频新闻
// @Tags AddVideoTemplate
// @Summary 新增视频新闻
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.VideoInfo true "VideoInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /addVideoTemplate [post]
/*func (NewsCtrl *NewsCtrl) AddVideoTemplate() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	videoInfo_Param := new(newsInfo.VideoInfo)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("新增视频新闻的入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &videoInfo_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.AddVideoTemplate(videoInfo_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "新增视频新闻成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("新增视频新闻失败:", err)
	}
}
*/
// @Title [上传]新增视频新闻
// @Tags AddVideoTemplateUpload
// @Summary 新增视频新闻
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.VideoInfo true "VideoInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /addVideoTemplateUpload [post]
func (NewsCtrl *NewsCtrl) AddVideoTemplateUpload() {
	//设置允许跨域请求
	NewsCtrl.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	//入参
	file, h, _ := NewsCtrl.GetFile("file")               //获取上传的文件
	newsVideoName := NewsCtrl.GetString("newsVideoName") //视频名称
	creater := NewsCtrl.GetString("creater")             //创建人
	corpCode := NewsCtrl.GetString("corpCode")           //企业code.
	corpName := NewsCtrl.GetString("corpName")           //企业name
	createTimeNew := NewsCtrl.GetString("createTimeNew") //管理员生成的创建时间
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求-mp4
	AllowExtMap := map[string]bool{
		".mp4": true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		resJson.Success = false
		resJson.Msg = "文件类型错误，请上传mp4格式的文件"
		return
	}
	//调用华为云OBS上传文件 //视频传入videoFile文件夹
	//保证唯一
	numNo := time.Now().Format("150405")
	obsPutFileRep, err := common.ObsPutFile(corpCode, numNo+"_"+h.Filename, file, "videoFile")
	if obsPutFileRep.Success {
		videoInfo_Param := new(newsInfo.VideoInfo)
		videoInfo_Param.NewsVideoName = newsVideoName
		videoInfo_Param.NewsVideoUrl = obsPutFileRep.ImgURL
		videoInfo_Param.Creater = creater
		videoInfo_Param.CorpCode = corpCode
		videoInfo_Param.CorpName = corpName
		videoInfo_Param.CreateTimeNew = createTimeNew
		//业务处理
		err := models.AddVideoTemplate(videoInfo_Param)
		if err == nil {
			resJson.Success = true
			resJson.Msg = "视频新增成功"
		} else {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("新增视频新闻失败:", err)
		}
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("视频文件上传OBS失败:%s", err)
	}
}

// @Title 删除视频媒体
// @Tags DeleteVideoNews
// @Summary 删除视频媒体
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.VideoInfo true "VideoInfo struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /deleteVideoNews [post]
func (NewsCtrl *NewsCtrl) DeleteVideoNews() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	videoInfo_Param := new(newsInfo.VideoInfo)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("删除视频媒体入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &videoInfo_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.DeleteVideoNews(videoInfo_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "删除视频媒体成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("删除视频媒体失败:", err)
	}
}

// @Title 超级管理员、用户 - 新增推广新闻
// @Tags AddNewsInUser
// @Summary 超级管理员、用户 - 新增推广新闻
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.NewsInfoInUser true "NewsInfoInUser struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /addNewsInUser [post]
func (NewsCtrl *NewsCtrl) AddNewsInUser() {
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	newsInfo_Param := new(newsInfo.NewsInfoInUser)
	var jsonByte = NewsCtrl.Ctx.Input.RequestBody
	logs.Info("超级管理员、用户 - 新增推广新闻入参：" + string(jsonByte))
	paramErr := jsoniter.Unmarshal(jsonByte, &newsInfo_Param)
	if paramErr != nil {
		resJson.Success = false
		resJson.Msg = "入参有误"
		return
	}
	//业务处理
	err := models.AddNewsInUser(newsInfo_Param)
	if err == nil {
		resJson.Success = true
		resJson.Msg = "新增推广新闻成功"
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("新增推广新闻失败:", err)
	}
}

// @Title 超级管理员、用户 - 新增推广视频
// @Tags AddVideoInUser
// @Summary 超级管理员、用户 - 新增推广视频
// @accept application/json
// @Produce application/json
// @Param data body newsInfo.VideoInfoInUser true "VideoInfoInUser struct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"提交成功"}"
// @router /addVideoInUser [post]
func (NewsCtrl *NewsCtrl) AddVideoInUser() {
	//设置允许跨域请求
	NewsCtrl.Ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	resJson := NewJsonStruct(nil)
	defer func() {
		NewsCtrl.Data["json"] = resJson
		NewsCtrl.ServeJSON()
	}()
	//入参
	file, h, _ := NewsCtrl.GetFile("file")               //获取上传的文件
	newsVideoName := NewsCtrl.GetString("newsVideoName") //视频名称
	creater := NewsCtrl.GetString("creater")             //创建人
	planId := NewsCtrl.GetString("planId")               //计划id.
	planName := NewsCtrl.GetString("planName")           //计划name
	corpCode := NewsCtrl.GetString("corpCode")           //企业code.
	corpName := NewsCtrl.GetString("corpName")           //企业name
	ext := path.Ext(h.Filename)
	//验证后缀名是否符合要求-mp4
	AllowExtMap := map[string]bool{
		".mp4": true,
	}
	if _, ok := AllowExtMap[ext]; !ok {
		resJson.Success = false
		resJson.Msg = "文件类型错误，请上传mp4格式的文件"
		return
	}
	//调用华为云OBS上传文件 //视频传入videoFile文件夹
	//保证唯一
	numNo := time.Now().Format("150405")
	obsPutFileRep, err := common.ObsPutFile(corpCode, numNo+"_"+h.Filename, file, "videoFile")
	if obsPutFileRep.Success {
		videoInfo_Param := new(newsInfo.VideoInfoInUser)
		videoInfo_Param.NewsVideoName = newsVideoName
		videoInfo_Param.NewsVideoUrl = obsPutFileRep.ImgURL
		videoInfo_Param.PlanId = planId
		videoInfo_Param.PlanName = planName
		videoInfo_Param.Creater = creater
		videoInfo_Param.CorpCode = corpCode
		videoInfo_Param.CorpName = corpName
		//业务处理
		err := models.AddVideoInUser(videoInfo_Param)
		if err == nil {
			resJson.Success = true
			resJson.Msg = "视频新增成功"
		} else {
			resJson.Success = false
			resJson.Msg = fmt.Sprintf("新增视频新闻失败:", err)
		}
	} else {
		resJson.Success = false
		resJson.Msg = fmt.Sprintf("视频文件上传OBS失败:%s", err)
	}
}
