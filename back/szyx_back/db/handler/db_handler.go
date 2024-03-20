package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
	"strings"
	"szyx_back/common"
	"szyx_back/configs"
)

type DbHandler struct {
	Db *sql.DB
}
type DbResultMapping map[string]interface{}

var dbHandler_self *DbHandler = nil

func NewDbHandler() *DbHandler {

	if dbHandler_self == nil {
		logs.Info("========DB连接初始化========")
		db, err1 := sql.Open("mysql", configs.DB_dataSource)
		common.ErrorHandler(err1, "DB连接初始化失败")

		db.SetMaxOpenConns(2000)
		db.SetMaxIdleConns(1000)

		handler := new(DbHandler)
		handler.Db = db
		dbHandler_self = handler
		return dbHandler_self
	} else {
		pingErr := dbHandler_self.Db.Ping()
		if pingErr != nil {
			dbHandler_self = nil
			dbHandler_self = NewDbHandler()
		}
		return dbHandler_self
	}

}

// 插入insert
func (dbHandler_self *DbHandler) Insert(sql_str string, args ...interface{}) (res int64, resultErr error) {
	//错误处理
	defer common.RecoverHandler(func(rcErr error) {
		res = 0
		resultErr = errors.New(fmt.Sprintf("插入insert记录发生错误::%s", rcErr))
	})
	printLog("Insert", sql_str, args)
	stmt, err := dbHandler_self.Db.Prepare(sql_str)
	defer stmt.Close()
	common.ErrorHandler(err, "数据库操作异常_Insert_1")

	rowsAffected, err := stmt.Exec(args...)
	common.ErrorHandler(err, "数据库操作异常_Insert_2")

	return rowsAffected.RowsAffected()
}

// 更新update
func (dbHandler_self *DbHandler) Update(sql_str string, args ...interface{}) (res int64, resultErr error) {
	//错误处理
	defer common.RecoverHandler(func(rcErr error) {
		res = 0
		resultErr = errors.New(fmt.Sprintf("更新update记录发生错误::%s", rcErr))
	})
	printLog("Update", sql_str, args)
	stmt, err := dbHandler_self.Db.Prepare(sql_str)
	defer stmt.Close()
	common.ErrorHandler(err, "数据库操作异常_Update_1")

	rowsAffected, err := stmt.Exec(args...)
	common.ErrorHandler(err, "数据库操作异常_Update_2")

	return rowsAffected.RowsAffected()
}

// 删除delete
func (dbHandler_self *DbHandler) Delete(sql_str string, args ...interface{}) (res int64, resultErr error) {
	//错误处理
	defer common.RecoverHandler(func(rcErr error) {
		res = 0
		resultErr = errors.New(fmt.Sprintf("删除delete记录发生错误::%s", rcErr))
	})
	printLog("Delete", sql_str, args)
	stmt, err := dbHandler_self.Db.Prepare(sql_str)
	defer stmt.Close()
	common.ErrorHandler(err, "数据库操作异常_Delete_1")

	rowsAffected, err := stmt.Exec(args...)
	common.ErrorHandler(err, "数据库操作异常_Delete_2")

	return rowsAffected.RowsAffected()
}

// 查询单条数据
func (dbHandler_self *DbHandler) SelectOne(sql_str string, args ...interface{}) (result DbResultMapping, resultErr error) {
	//错误处理
	defer common.RecoverHandler(func(rcErr error) {
		result = nil
		resultErr = errors.New(fmt.Sprintf("查询单条数据发生错误::%s", rcErr))
	})
	printLog("Query", sql_str, args)
	stmt, err := dbHandler_self.Db.Prepare(sql_str)
	defer stmt.Close()
	common.ErrorHandler(err, "数据库操作异常_SelectOne_1")

	rows, err := stmt.Query(args...)
	defer rows.Close()
	common.ErrorHandler(err, "数据库操作异常_SelectOne_2")

	return obtainDbResultMapping(rows), nil
}

// 查询多条数据
func (dbHandler_self *DbHandler) SelectList(sql_str string, args ...interface{}) (result []DbResultMapping, resultErr error) {
	//错误处理
	defer common.RecoverHandler(func(rcErr error) {
		result = nil
		resultErr = errors.New(fmt.Sprintf("查询多条数据发生错误::%s", rcErr))
	})
	printLog("Query", sql_str, args)
	stmt, err := dbHandler_self.Db.Prepare(sql_str)
	defer stmt.Close()
	common.ErrorHandler(err, "数据库操作异常_SelectList_1")

	rows, err := stmt.Query(args...)
	common.ErrorHandler(err, "数据库操作异常_SelectList_2")

	return obtainDbResultMappingArray(rows), nil
}

/**
打印SQL
*/
func printLog(sqlType string, sql_str string, args ...interface{}) {
	if configs.LOG_SQL_FLAG {
		logs.Info(sqlType + "::" + sql_str)
		logs.Info(args)
	}

}

/**
in条件拼接
*/
func (dbHandler_self *DbHandler) Placeholders(n int) string {
	ps := make([]string, n)
	for i := 0; i < n; i++ {
		ps[i] = ":param" + strconv.Itoa(i+1)
	}
	in_str := "(" + strings.Join(ps, ",") + ")"
	fmt.Println(in_str)
	return in_str
}

func obtainDbResultMapping(rows *sql.Rows) DbResultMapping {

	column, err := rows.Columns() //读出查询出的列字段名
	common.ErrorHandler(err, "数据库操作异常_obtainDbResultMapping_1")

	//values是每个列的值，这里获取到byte里
	// renhao:不能是byte，也不能是string
	//1，为了通用，用interface，2，有些数据string化后，无法转回来
	values := make([]interface{}, len(column))
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度

	for i := range values {
		scans[i] = &values[i]
	}

	var results DbResultMapping //最后得到的map
	for rows.Next() {
		err := rows.Scan(scans...)
		common.ErrorHandler(err, "数据库操作异常_obtainDbResultMapping_2")

		row := make(DbResultMapping) //每行数据
		for k, v := range values {
			//每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = v
		}
		results = row
		break //只取第一条
	}
	return results
}
func obtainDbResultMappingArray(rows *sql.Rows) []DbResultMapping {

	column, err := rows.Columns() //读出查询出的列字段名
	common.ErrorHandler(err, "数据库操作异常_obtainDbResultMappingArray_1")

	//values是每个列的值，这里获取到byte里
	// renhao:不能是byte，也不能是string
	//1，为了通用，用interface，2，有些数据string化后，无法转回来
	values := make([]interface{}, len(column))
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度

	for i := range values {
		scans[i] = &values[i]
	}

	var results []DbResultMapping //最后得到的map
	for rows.Next() {
		err := rows.Scan(scans...)
		common.ErrorHandler(err, "数据库操作异常_obtainDbResultMappingArray_2")

		row := make(DbResultMapping) //每行数据
		for k, v := range values {
			//每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = v
		}
		results = append(results, row)
	}
	rows.Close()
	return results
}
