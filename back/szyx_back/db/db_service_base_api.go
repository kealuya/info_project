package db

import (
	"github.com/mitchellh/mapstructure"
	"reflect"
	"szyx_back/common"
	"time"
)

//mapstructure.DecoderConfig 提供 转换Hook方法
var mapstructure_hook_func = func(t1 reflect.Type, t2 reflect.Type, v interface{}) (interface{}, error) {
	temp := v
	//解决 struct 从Time转换到String的问题
	if t1.Name() == "Time" && t2.Name() == "string" {
		temp = common.FormatDate(v.(time.Time), common.YYYY_MM_DD_HH_MM_SS)
	}
	return temp, nil
}

//追加DecoderConfig，返回Decoder，传入result结果参数指针
func ObtainDecoderConfig(result interface{}) *mapstructure.Decoder {
	config := &mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		Result:           result,
		DecodeHook:       mapstructure_hook_func,
	}
	decoder, _ := mapstructure.NewDecoder(config)
	return decoder
}
