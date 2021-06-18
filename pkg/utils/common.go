package utils

import (
	"fmt"
	"strconv"
	"time"
)

func StrToTime(s string) time.Time {
	layout := "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if len(s) == 10{
		s = fmt.Sprintf("%s 00:00:00", s)
	}
	getTime, _ := time.ParseInLocation(layout, s, loc)
	return getTime
}

func StrToUint64(s string) uint64  {
	n, _ := strconv.ParseInt(s,10,64)
	return uint64(n)
}

func StrToInt64(s string) int64  {
	n, _ := strconv.ParseInt(s,10,64)
	return int64(n)
}

func StrToUint32(s string) uint32  {
	n, _ := strconv.Atoi(s)
	return uint32(n)
}

func StrToInt32(s string) int32  {
	n, _ := strconv.Atoi(s)
	return int32(n)
}
