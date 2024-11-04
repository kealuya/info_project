package web

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	configs "my_pilot/config"
	"time"
)

// 自定义ResponseWriter
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 重写Write方法
func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginLogger(c *gin.Context) {
	// 开始时间
	start := time.Now()
	// 匹配唯一编码
	no := uuid.NewString()[:6]
	req, _ := io.ReadAll(c.Request.Body)
	logs.Info(fmt.Sprintf("[请求日志:%s] %s | %s | %s | reuqest_body:%s", no,
		c.ClientIP(), c.Request.Method, c.Request.Host+c.Request.RequestURI, string(req)))

	// 使用自定义ResponseWriter
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw

	// ----请求前----
	c.Next()
	// ----请求后----
	// 结束时间
	end := time.Now()
	// 执行时间
	latency := end.Sub(start)

	// 获取状态码
	statusCode := c.Writer.Status()

	// 获取错误信息
	errors := c.Errors.String()

	// 获取处理器名称（如果需要）
	handlerName := c.HandlerName()

	// 请求路径
	path := c.Request.Host + c.Request.RequestURI

	// 获取响应体
	config := configs.GetConfig()
	logMessage := ""
	if config["server"]["mode"] == "debug" {
		responseBody := blw.body.String()
		// 构造日志信息
		logMessage = fmt.Sprintf("[响应日志:%s] %3d | %13v | %s | %s | response_body:%s",
			no,
			statusCode,
			latency,
			path,
			handlerName,
			responseBody,
		)
	} else {
		// 构造日志信息
		logMessage = fmt.Sprintf("[响应日志:%s] %3d | %13v | %s | %s ",
			no,
			statusCode,
			latency,
			path,
			handlerName,
		)
	}

	// 根据状态码选择日志级别
	switch {
	case statusCode >= 500:
		logs.Error(logMessage)
	case statusCode >= 400:
		logs.Warn(logMessage)
	default:
		logs.Info(logMessage)
	}

	// 如果有错误，额外记录错误信息
	if errors != "" {
		logs.Error(fmt.Printf("[错误记录:%s] Errors: %s", no, errors))
	}
}
