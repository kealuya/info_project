package models

import (
	"szyx_back/db"
	"szyx_back/entity/worth"
)

/**
导出申请过的价值excel
*/
func ExportWorthExcel(worth *worth.Worth) (res []worth.Worth, err error) {

	return res, err
}

/**
价值申请
*/
func ApplyWorth(worth *worth.Worth) (err error) {
	//价值申请
	_, err = db.ModifyWorthApply(worth)
	return err
}

/**
价值列表
*/
func GetWorthList(worthList_Param *worth.WorthList_Param) (res worth.WorthList_Result, err error) {
	res, err = db.GetWorthList(worthList_Param)
	return res, err
}
