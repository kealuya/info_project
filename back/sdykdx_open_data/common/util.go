package common

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"log"
	"runtime/debug"
)

// RecoverHandler 共通错误recover处理方法
func RecoverHandler[T any](f func(err T)) {
	if err := recover(); err != nil {
		logs.Error(err)
		logs.Error(string(debug.Stack()))
		switch e := err.(type) {
		case string:
			casted := errors.New(e).(T)
			f(casted)
		case error:
			casted := e.(T)
			f(casted)
		default:
			casted := fmt.Errorf("unknown error: %v", err).(T)
			f(casted)
		}
	}
}

// ErrorHandler 共通错误error处理方法
func ErrorHandler(err error, info ...any) {
	if err != nil {
		log.Panic(err, info)
	}
}
