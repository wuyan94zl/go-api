package logger

import (
	"log"
)

var console *consoleLog

type consoleLog struct {
	loggerBase
}

func getConsoleLog() *consoleLog {
	if console != nil {
		return console
	} else {
		console = &consoleLog{}
		console.log = log.Default()
		return console
	}
}
