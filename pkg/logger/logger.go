package logger

import (
	"fmt"
	"log"
)

// LogError 当存在错误时记录日志
func LogError(err error) {
	if err != nil {
		log.Println(err)
	}
}

func SystemError(err error)  {
	if err != nil{
		fmt.Println(err)
	}
}
