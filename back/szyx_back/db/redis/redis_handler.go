package redis

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"time"
)

var RedisCache cache.Cache

func init() {
	logs.Info("========redis连接初始化========")
	redisHost := beego.AppConfig.String("RedisHost")
	dataBase := beego.AppConfig.String("RedisDbNum")
	password := beego.AppConfig.String("RedisPass")
	redisKey := beego.AppConfig.String("RedisKey")
	config := fmt.Sprintf(`{"key":"%s","conn":"%s","dbNum":"%s","password":"%s"}`, redisKey, redisHost, dataBase, password)
	var err error
	//config := fmt.Sprintf(`{"conn":"%s","dbNum":"%s"}`, redisHost, dataBase)
	RedisCache, err = cache.NewCache("redis", config)
	if err != nil || RedisCache == nil {
		errMsg := "redis初始化失败"
		logs.Error(errMsg)
	}
}

/**
缓存一个值
*/
func SetStr(key, value string, time time.Duration) (err error) {
	err = RedisCache.Put(key, value, time)
	if err != nil {
		logs.Error("缓存失败")
	}
	return
}

/**
取出缓存中的一个值
*/
func GetStr(key string) (value string) {
	v := RedisCache.Get(key)
	if v != nil {
		value = string(v.([]byte)) //这里的转换很重要，Get返回的是interface
		return
	} else {
		value = ""
		return
	}

}

/**
删除缓存中的一个值
*/
func DelKey(key string) (err error) {
	err = RedisCache.Delete(key)
	return
}

/**
取出缓存中的一个值
*/
func GetStrs(key string) (value string) {
	v := RedisCache.Get(key)
	if v == nil {
		return ""
	}
	value = string(v.([]byte)) //这里的转换很重要，Get返回的是interface
	return
}
