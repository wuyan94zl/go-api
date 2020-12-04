package utils

import (
	"fmt"
	"time"
)

func SetStrToTime(t string) time.Time {
	layout := "2006-01-02 15:04:05" //时间常量
	loc, _ := time.LoadLocation("Asia/Shanghai")
	if len(t) == 10{
		t = fmt.Sprintf("%s 00:00:00",t)
	}
	getTime, _ := time.ParseInLocation(layout, t, loc)
	return getTime
}
