package common

import (
	"strings"
	time "time"
)

type DateStyle string

const (
	MM_DD                   = "MM-dd"
	MMDD                    = "MMddHHmm"
	YYYYMM                  = "yyyyMM"
	YYYY_MM                 = "yyyy-MM"
	YYYY_MM_DD              = "yyyy-MM-dd"
	YYYYMMDD                = "yyyyMMdd"
	YYYYMMDDHHMMSS          = "yyyyMMddHHmmss"
	YYYYMMDDHHMM            = "yyyyMMddHHmm"
	YYYYMMDDHH              = "yyyyMMddHH"
	YYMMDDHHMM              = "yyMMddHHmm"
	MM_DD_HH_MM             = "MM-dd HH:mm"
	MM_DD_HH_MM_SS          = "MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM        = "yyyy-MM-dd HH:mm"
	YYYY_MM_DD_HH_MM_SS     = "yyyy-MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_W   = "yyyy-MM-ddHH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS = "yyyy-MM-dd HH:mm:ss.SSS"

	MM_DD_EN                   = "MM/dd"
	YYYY_MM_EN                 = "yyyy/MM"
	YYYY_MM_DD_EN              = "yyyy/MM/dd"
	MM_DD_HH_MM_EN             = "MM/dd HH:mm"
	MM_DD_HH_MM_SS_EN          = "MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_EN        = "yyyy/MM/dd HH:mm"
	YYYY_MM_DD_HH_MM_SS_EN     = "yyyy/MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS_EN = "yyyy/MM/dd HH:mm:ss.SSS"

	MM_DD_CN               = "MM月dd日"
	YYYY_MM_CN             = "yyyy年MM月"
	YYYY_MM_DD_CN          = "yyyy年MM月dd日"
	MM_DD_HH_MM_CN         = "MM月dd日 HH:mm"
	MM_DD_HH_MM_SS_CN      = "MM月dd日 HH:mm:ss"
	YYYY_MM_DD_HH_MM_CN    = "yyyy年MM月dd日 HH:mm"
	YYYY_MM_DD_HH_MM_SS_CN = "yyyy年MM月dd日 HH:mm:ss"

	HH_MM       = "HH:mm"
	HH_MM_SS    = "HH:mm:ss"
	HH_MM_SS_MS = "HH:mm:ss.SSS"

	TIME_LAYOUT      = "2006-01-02 15:04:05"
	TIME_LAYOUT_DATE = "2006-01-02"
	TIME_LAYOUT_STR  = "20060102150405"
	DATE_LAYOUT_STR  = "20060102"
)

// 日期转字符串
func FormatDate(date time.Time, dateStyle DateStyle) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)

	return date.Format(layout)
}

func StringToTime(timeStr string) time.Time {
	result, err := time.ParseInLocation(TIME_LAYOUT, timeStr, time.Local)
	ErrorHandler(err, "时间转换发生错误！")
	return result
}

// yyyyMMddHHmmss转换成yyyy-MM-dd HH:mm:ss
func FormatDateTimeStr(timeStr string) string {
	str := ""
	if timeStr != "" {
		result, err := time.ParseInLocation(TIME_LAYOUT_STR, timeStr, time.Local)
		ErrorHandler(err, "时间转换发生错误！")
		str = result.Format(TIME_LAYOUT)
	}
	return str
}

// yyyyMMddHH转换成yyyy-MM-dd
func FormatDateStr(timeStr string) string {
	str := ""
	if timeStr != "" {
		result, err := time.ParseInLocation(DATE_LAYOUT_STR, timeStr, time.Local)
		ErrorHandler(err, "时间转换发生错误！")
		str = result.Format(TIME_LAYOUT_DATE)
	}
	return str
}

func GetDiffDate(start string, end string) int64 {
	startUnix, _ := time.ParseInLocation(TIME_LAYOUT_DATE, start, time.Local)
	endUnix, _ := time.ParseInLocation(TIME_LAYOUT_DATE, end, time.Local)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	// 求相差天数
	date := (endTime - startTime) / 86400
	return date
}

func GetHourDiffer(start string, end string) string {
	var hour int64
	t1, _ := time.ParseInLocation(TIME_LAYOUT, start, time.Local)
	t2, _ := time.ParseInLocation(TIME_LAYOUT, end, time.Local)
	diff := t2.Unix() - t1.Unix() //
	hour = diff / 3600
	Hour := Int64ToString(hour)
	stime := start[14 : len(start)-3]
	etime := end[14 : len(end)-3]
	time := Int64ToString(StringToInt64(etime) - StringToInt64(stime))
	runtime := Hour + "时" + time + "分"
	return runtime
}

// 获取指定日期的后一天
func GetNextDateStr(dateStr string) string {
	comeDateTime, _ := time.ParseInLocation(
		TIME_LAYOUT_DATE, dateStr, time.Local)
	nextDate := comeDateTime.Add(24 * time.Hour)
	return nextDate.Format(TIME_LAYOUT_DATE)
}

func CurrentTime() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	replace := strings.Replace(currentTime, "-", "", -1)
	s := strings.Replace(replace, " ", "", -1)
	i := strings.Replace(s, ":", "", -1)
	return i
}
