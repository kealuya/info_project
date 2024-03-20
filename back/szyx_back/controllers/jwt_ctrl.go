package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"strings"
	"szyx_back/common"
	"szyx_back/db/redis"
	"time"
)

/**
JWt 续签
*/
type JwtController struct {
	beego.Controller
}

/**
jwt续签请求入参
*/
type RefreshTokenRequestInfo struct {
	EmpCode       string `description:"用户编号" json:"empCode"`
	CorpCode      string `description:"企业编号" json:"corpCode"`
	EquipmentType string `description:"设备类型" json:"equipmentType"` //登录设备类型 PC,APP
	Token         string `json:"token"`                            //刷新的token
}

// @Title jwt续签
// @router /refreshToken [post]
// @Tags RefreshToken
// @Summary jwt续签
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body RefreshTokenRequestInfo true "token"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"回写成功"}"
func (jwtCtrl *JwtController) RefreshToken() {
	//返回值
	resultMap := make(map[string]interface{})
	logs.Debug("================token续签接口=============")

	var input_pull = new(RefreshTokenRequestInfo)
	var jsonByte = jwtCtrl.Ctx.Input.RequestBody
	var reqString = string(jsonByte)
	logs.Info("接收到的参数：" + reqString)

	//处理接收到的信息
	reqString = strings.ReplaceAll(reqString, "\\\\n", "")
	reqString = strings.ReplaceAll(reqString, "\\n", "")
	reqString = strings.ReplaceAll(reqString, "\"{", "{")
	reqString = strings.ReplaceAll(reqString, "}\"", "}")
	reqString = strings.ReplaceAll(reqString, "\\\"", "\"")

	common.Unmarshal([]byte(reqString), &input_pull)
	//判断token的合法性（此时前端传参是refashToken）
	claims, err_ParseToken := common.ParseToken(input_pull.Token)
	if err_ParseToken == nil {
		//判断token与入参，是否为同一人所有
		if claims.EmpCode == input_pull.EmpCode && claims.CorpCode == input_pull.CorpCode {
			//判断，redis中是否有当前登录信息  empcode_corpcode_equipmentType作为key值
			key := input_pull.EmpCode + "_" + input_pull.CorpCode + "_" + input_pull.EquipmentType
			redisCacheToken := redis.GetStr(key)
			//判断token未过期的情况下且redis存储和前端一致
			if redisCacheToken != "" && redisCacheToken == input_pull.Token {
				//签发新的token,更新redis中token的过期时间
				jwtStr, err1 := common.GenerateToken(input_pull.EmpCode, "", input_pull.CorpCode)
				common.ErrorHandler(err1, "签发access_token失败")
				//签发refreshToken
				refreshToken, err2 := common.GenerateRefreshToken(input_pull.EmpCode, "", input_pull.CorpCode)
				common.ErrorHandler(err2, "签发refresh_token失败")
				err3 := redis.SetStr(key, refreshToken, 35*time.Minute)
				common.ErrorHandler(err3, "redis,cachaToken失败")

				resultMap["code"] = 200
				resultMap["msg"] = "token刷新成功"
				resultMap["success"] = true
				resultMap["jwt"] = jwtStr
				resultMap["refreshToken"] = refreshToken

				_, _ = jwtCtrl.Ctx.ResponseWriter.Write(common.Marshal(resultMap))

			} else if redisCacheToken == "" { //token在redis中过期情况
				resultMap["code"] = 500
				resultMap["msg"] = "token已过期，请重新登录"
				resultMap["success"] = false
				_, _ = jwtCtrl.Ctx.ResponseWriter.Write(common.Marshal(resultMap))
			} else { //用户在其他电脑登录
				resultMap["code"] = 500
				resultMap["msg"] = "当前用户已下线或在其它地方登录"
				resultMap["success"] = false
				_, _ = jwtCtrl.Ctx.ResponseWriter.Write(common.Marshal(resultMap))
			}
		} else {
			resultMap["code"] = 500
			resultMap["msg"] = "请求用户不匹配"
			resultMap["success"] = false
			_, _ = jwtCtrl.Ctx.ResponseWriter.Write(common.Marshal(resultMap))
		}
	} else {
		resultMap["code"] = 500
		resultMap["msg"] = "token已过期，请重新登录"
		resultMap["success"] = false
		_, _ = jwtCtrl.Ctx.ResponseWriter.Write(common.Marshal(resultMap))
	}

}

/**
jwt获取最新的token入参
*/
type GetTokenRequestInfo struct {
	EmpCode  string `json:"EmpCode"`  //用户code
	CorpCode string `json:"corpCode"` //用户id
}

/**
jwt获取最新的token出参
*/
type GetTokenResultInfo struct {
	Token string `json:"token"` //获取到的token
}

// @Title 获取token
// @router /getToken [post]
// @Tags GetToken
// @Summary jwt获取token
// @Security token
// @accept application/json
// @Produce application/json
// @Param data body GetTokenRequestInfo true "token"
// @Success 200 {object} GetTokenResultInfo
func (jwtCtrl *JwtController) GetToken() {
	var resJson = NewJsonStruct(nil)
	dataJson := new(GetTokenResultInfo)

	defer func() {
		jwtCtrl.Data["json"] = string(common.Marshal(resJson))
		jwtCtrl.ServeJSON()
	}()

	logs.Debug("================token获取接口=============")
	//TODO ip稳定后需要将下面的代码放开
	flag := true

	//flag := false
	//
	//
	////取当前请求的ip
	//ip := jwtCtrl.Ctx.Request.RemoteAddr
	//
	//ipList :=[]string{"192.168.120.242","192.168.120.222"}
	//
	//for _,v := range ipList {
	//	if (strings.Contains(ip,v)){
	//		flag = true
	//	}
	//}

	if flag {
		//请求入参
		var input_pull = new(GetTokenRequestInfo)
		var jsonByte = jwtCtrl.Ctx.Input.RequestBody
		var reqString = string(jsonByte)
		logs.Info("接收到的参数：" + reqString)

		//处理接收到的信息
		reqString = strings.ReplaceAll(reqString, "\\\\n", "")
		reqString = strings.ReplaceAll(reqString, "\\n", "")
		reqString = strings.ReplaceAll(reqString, "\"{", "{")
		reqString = strings.ReplaceAll(reqString, "}\"", "}")
		reqString = strings.ReplaceAll(reqString, "\\\"", "\"")

		common.Unmarshal([]byte(reqString), &input_pull)

		token, err := common.GenerateToken(input_pull.EmpCode, "", input_pull.CorpCode)
		if err == nil {
			dataJson.Token = token
			resJson.Success = true
			resJson.Msg = "token获取成功"
			resJson.Data = dataJson
		} else {
			resJson.Success = false
			resJson.Msg = "token获取失败"
			resJson.Data = dataJson
		}

	} else {
		resJson.Success = false
		resJson.Msg = "访问限制"
	}

}
