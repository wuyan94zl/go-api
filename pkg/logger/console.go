package logger

import (
	"log"
)

var console *consoleLog

type consoleLog struct {
	log *log.Logger
}

func (f *consoleLog) info(logs ...interface{}) {
	f.log.SetPrefix("[api-debug] info ")
	f.log.Println(logs...)
}

func (f *consoleLog) error(logs ...interface{}) {
	f.log.SetPrefix("[api-debug] error ")
	f.log.Println(logs...)
}

func (f *consoleLog) warning(logs ...interface{}) {
	f.log.SetPrefix("[api-debug] warning ")
	f.log.Println(logs...)
}

func (f *consoleLog) notice(logs ...interface{}) {
	f.log.SetPrefix("[api-debug] notice ")
	f.log.Println(logs...)
}

func getConsoleLog() *consoleLog {
	if console != nil {
		return console
	} else {
		return &consoleLog{log: log.Default()}
	}
}
