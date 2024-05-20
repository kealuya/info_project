package db

import (
	"szyx_back/common"
	db_handler "szyx_back/db/handler"
	"szyx_back/entity/worth"
	"time"
)

//提交申请价值
func ModifyWorthApply(info *worth.Worth) (res worth.Worth, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//申请价值时间
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	var Param []interface{}
	Param = append(Param, common.MY_WORTH_APPLY_FLAG_KEY_1)
	Param = append(Param, currentTime)
	Param = append(Param, info.UserId)
	Param = append(Param, info.CorpCode)
	Param = append(Param, info.WorthId)

	num, err := dbHandler.Update(db_handler.ModifyWorthByApplyWorth_sql, Param...)
	if num <= 0 {
		common.ErrorHandler(err, "价值申请数据库保存发生错误!")
	} else {
		//查询价值，返回前台展示
		res1, err1 := dbHandler.SelectOne(db_handler.SelectWorthById_sql, info.WorthId)
		if len(res1) > 0 && err1 == nil {
			decoder := ObtainDecoderConfig(&res)
			err := decoder.Decode(res1)
			common.ErrorHandler(err, "价值申请数据库查询转换发生错误!")
		} else {
			err = err1
		}
	}

	return res, err
}

//查看价值列表
func GetWorthList(info *worth.WorthList_Param) (res worth.WorthList_Result, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//价值信息
	var Param []interface{}
	Param = append(Param, info.CorpCode)
	Param = append(Param, info.UserId)
	Param = append(Param, info.Status)

	//计算limit起始值
	startNum := (info.CurrentPage - 1) * info.PageSize
	Param = append(Param, startNum)
	Param = append(Param, info.PageSize)

	selRes, err := dbHandler.SelectList(db_handler.GetWorthList_sql, Param...)

	worthList := []worth.Worth{}

	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&worthList)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "价值列表信息转换发生错误!")
	}

	worthListCount := []worth.Worth{}
	var ParamCount []interface{}
	ParamCount = append(ParamCount, info.CorpCode)
	ParamCount = append(ParamCount, info.UserId)
	ParamCount = append(ParamCount, info.Status)
	selCountRes, err2 := dbHandler.SelectList(db_handler.GetWorthListCount_sql, ParamCount...)
	if err2 == nil {
		decoder := ObtainDecoderConfig(&worthListCount)
		err1 := decoder.Decode(selCountRes)
		common.ErrorHandler(err1, "价值列表分页信息转换发生错误!")
	}

	res.WorthList = worthList
	res.TotalCount = int64(len(worthListCount))
	//获取总页数，前端需要
	res.PageCount = res.TotalCount / info.PageSize
	if res.TotalCount%info.PageSize > 0 {
		res.PageCount++
	}
	return res, err
}

//查看价值详情
func GetWorthDetails(info *worth.Worth) (res worth.Worth, msg error) {
	dbHandler := db_handler.NewDbHandler()
	//价值详情信息
	var Param []interface{}
 	Param = append(Param, info.CorpCode)
	Param = append(Param, info.UserId)
	Param = append(Param, info.WorthId)


	selRes, err := dbHandler.SelectOne(db_handler.SelectWorthById_sql, Param...)
	if len(selRes) > 0 && err == nil {
		decoder := ObtainDecoderConfig(&res)
		err1 := decoder.Decode(selRes)
		common.ErrorHandler(err1, "价值信息转换发生错误!")
	}

	return res, err
}
