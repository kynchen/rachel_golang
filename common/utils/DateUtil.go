package utils

import (
	"github.com/kynchen/rachel_golang/common/constants"
	"time"
)



// TimestampToDate 时间戳转日期格式
func TimestampToDate(timestamp int) string {
	timeUnix := time.Unix(int64(timestamp), 0)
	return timeUnix.Format(constants.DefaultTime)
}

// DateToTimestamp 日期转时间戳
func DateToTimestamp(date string) (int64,error)  {
	parse, err := time.Parse(constants.DefaultTime, date)
	return parse.Unix(),err
}

// CurrentUnix 获取当前时间戳
func CurrentUnix() int64 {
	return time.Now().Unix()
}

// CurrentDate 获取当前日期
func CurrentDate() string {
	return time.Now().Format(constants.DefaultTime)
}