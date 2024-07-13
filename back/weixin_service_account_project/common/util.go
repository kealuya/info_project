package common

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	jsoniter "github.com/json-iterator/go"
	"github.com/mozillazg/go-pinyin"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"reflect"
	"regexp"
	"runtime"
	"runtime/debug"
	"strconv"
	"strings"
	"time"
)

func init() {

}

// 获取当前路径
func GetCurrentPath() string {
	s, _ := exec.LookPath(os.Args[0])
	i := strings.LastIndex(s, "\\")
	path := string(s[0 : i+1])
	return path
}

// Float64转String
func Float64ToString(s1 float64, prec int) string {
	return strconv.FormatFloat(s1, 'f', prec, 64) //float64
}

// Float32转String
func Float32ToString(s1 float64, prec int) string {
	//s2 := strconv.FormatFloat(v, 'E', -1, 64)
	return strconv.FormatFloat(s1, 'f', prec, 32) //float32
}

// String转Int64
func StringToInt64(str string) int64 {
	n, _ := strconv.ParseInt(str, 10, 64)
	return n
}

// string转成int
func StringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

// String转Float64
func StringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// String转Float64
func StringToFloat64Jd(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

// Int64 转string
func Int64ToString(str int64) string {
	i := strconv.FormatInt(str, 10)
	return i
}

// 获取GoroutineID
func GetGoroutineID() string {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)

	return "goroutine:" + fmt.Sprintf(" %v", n)
}

// 共通错误recover处理方法
func RecoverHandler(f func(err error)) {
	if err := recover(); err != nil {
		logs.Error(string(debug.Stack()))
		if f != nil {
			f(err.(error))
		}
	}
}

// 共通错误error处理方法
func ErrorHandler(err error, info ...any) {
	if err != nil {
		logs.Error(err, info)
		//logs.Error(string(debug.Stack()))
		log.Panicln()
	}
}

// 字符串md5加密
func StringToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}

// 字符串md5加密 32位小写
func StringToMd5_X(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 字符转base64
func StringToBase64(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

// base64转字符
func Base64ToString(data string) string {
	ds, _ := base64.StdEncoding.DecodeString(data)
	return string(ds)
}

func Sha1s(s string) string {
	r := sha1.Sum([]byte(s))
	return hex.EncodeToString(r[:])
}

func Unmarshal(data []byte, v interface{}) {
	err := jsoniter.Unmarshal(data, v)
	if err != nil {
		ErrorHandler(err, "jsoniter.Unmarshal发生错误")
	}
}

func Marshal(v interface{}) []byte {
	jsonByte, err := jsoniter.Marshal(v)
	if err != nil {
		ErrorHandler(err, "jsoniter.Marshal发生错误")
		return []byte{}
	}
	return jsonByte
}

// struct深度拷贝
func DeepCopyStructToStruct(toStruct interface{}, fromStruct interface{}) error {
	toStruct_reflectType := reflect.TypeOf(toStruct)
	toStruct_reflectValue := reflect.ValueOf(toStruct)
	if toStruct_reflectType.Kind().String() != "ptr" {
		return fmt.Errorf("请传递指针对象进行赋值操作")
	}
	fromStruct_reflectType := reflect.TypeOf(fromStruct)
	fromStruct_reflectValue := reflect.ValueOf(fromStruct)
	for i := 0; i < fromStruct_reflectType.NumField(); i++ {
		plf := fromStruct_reflectType.Field(i)
		plfi := fromStruct_reflectValue.FieldByName(plf.Name).Interface()
		ppef := toStruct_reflectValue.Elem().FieldByName(plf.Name)
		ppef.Set(reflect.ValueOf(plfi))
	}
	return nil
}

// 去掉重复值
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return newArr
}

// 返回单位为：千米
func GetDistance(lat1, lat2, lng1, lng2 float64) float64 {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	result, _ := strconv.ParseFloat(fmt.Sprintf(
		"%.2f", dist*radius/1000), 64)
	return result
}

// 身份证号获取生日
func GetIDentityNO(IDentityNO string) string {
	reg, err := regexp.Compile("^(\\d{6})(\\d{8})(.*)")

	if err != nil {
		return ""
	}
	birthDate := ""
	if reg.MatchString(IDentityNO) == true {
		subMatch := reg.FindStringSubmatch(IDentityNO)
		dateStr := subMatch[2]
		birthDate = FormatDateStr(dateStr)
	}
	return birthDate
}

// 随机生成六位数
func CreateCaptcha() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}

// 随机生成四位数
func Createcode() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)) //这里面前面的04v是和后面的1000相对应的
}

// 短信生成验证码
func GenValidateCode(width int) string {
	numeric := [9]byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func TurnPinyin(str string) string {
	convert := pinyin.LazyConvert(str, nil)
	var pinyin string
	for _, s := range convert {
		pinyin += s
	}
	return pinyin
}

func ApproveState(t string) string {
	switch t {
	case "0":
		return "全部"
	case "1":
		return "编辑中"
	case "3":
		return "已撤回"
	case "4":
		return "已驳回"
	case "5":
		return "审批中"
	case "6":
		return "已通过"
	case "7":
		return "已关闭"
	case "8":
		return "已报废"
	default:
		return "其他"

	}
}

func ReimbursementState(t string) string {
	switch t {
	case "0":
		return "全部"
	case "1":
		return "编辑中"
	case "3":
		return "已撤回"
	case "4":
		return "已驳回"
	case "5":
		return "审批中"
	case "6":
		return "已通过"
	case "7":
		return "审核通过"
	case "8":
		return "审核驳回"
	case "9":
		return "待付款"
	case "10":
		return "开票驳回"
	case "11":
		return "付款中"
	case "12":
		return "已付款"
	default:
		return "其他"

	}
}

// 去掉转义字符
func DisableEscapeHtml(data interface{}) (string, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	if err := jsonEncoder.Encode(data); err != nil {
		return "", err
	}
	return bf.String(), nil
}

// 判断数组中是否存在相同的字符
func IsContain(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func HotelOrderStatus(t string) string {
	switch t {
	case "WAIT_CONFIRM":
		return "等待确认"
	case "WAIT_PAY":
		return "等待支付"
	case "NO_SHOW":
		return "未入住"
	case "CONFIRM":
		return "已确认"
	case "CHECK_OUT":
		return "已离店"
	case "CANCELED":
		return "已取消"
	case "CHECK_IN":
		return "已入住"
	case "INVALID":
		return "无效"
	default:
		return "其他"

	}
}

// mode层用的使用方法转换类型
func GetStringTypeData(param interface{}) (result string) {
	typeString := fmt.Sprintf("%T", param)
	value := ""
	switch typeString {
	case "float64":
		value = strconv.FormatFloat(param.(float64), 'g', -1, 64)
		break
	case "int64":
		value = strconv.FormatInt(param.(int64), 64)
		break
	case "string":
		value = param.(string)
		break
	}
	return value
}

// 航信开票的加密方法，请求java打包的类。
func HX_FpOrderEncode(fpString string) string {
	//如果调用本地加密  执行语句cmd —— java -jar fp_ctrl_0.1.jar
	/*req := httplib.Post("http://localhost:8001/encrypt")
	req.Param("code", fpString)
	enCode, err := req.String()
	ErrorHandler(err, "调动本机java发票加密失败::%+v")*/
	req2 := httplib.Post("http://124.70.111.57:10003/v1/invoice/encrypt")
	//req2 := httplib.Post("http://localhost:10003/v1/invoice/encrypt")
	req2.Param("code", fpString)
	enCode2, err2 := req2.String()
	fmt.Println(enCode2)
	ErrorHandler(err2, "调动cwxt的java发票加密方法失败::%+v")
	//去除springboot返回的字符串转义问题
	enCode2 = strings.Replace(enCode2, "\\n", "", -1)
	//本地启用加\\r  服务器上的只有\n
	//enCode2 = strings.Replace(enCode2, "\\r\\n", "", -1)
	enCode2 = strings.Replace(enCode2, "\"", "", -1)
	return enCode2
}
