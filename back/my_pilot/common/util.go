package common

import (
	"context"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"golang.org/x/time/rate"
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

// RateLimiter 限流器
type RateLimiter struct {
	rateLimiter     *rate.Limiter
	concurrentLimit chan struct{}
	ctx             context.Context
}

func NewRateLimiter(rateLimit float64, burst int, concurrent int, ctx context.Context) *RateLimiter {
	return &RateLimiter{
		rateLimiter:     rate.NewLimiter(rate.Limit(rateLimit), burst),
		concurrentLimit: make(chan struct{}, concurrent),
		ctx:             ctx,
	}
}

func (l *RateLimiter) Acquire() error {
	if err := l.rateLimiter.Wait(l.ctx); err != nil {
		return fmt.Errorf("rate limit: %w", err)
	}

	select {
	case l.concurrentLimit <- struct{}{}:
		return nil
	case <-l.ctx.Done():
		return fmt.Errorf("concurrent limit: %w", l.ctx.Err())
	}
}

func (l *RateLimiter) Release() {
	<-l.concurrentLimit
}
