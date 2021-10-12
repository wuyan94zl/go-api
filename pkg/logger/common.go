package logger

import "log"

type logger interface {
	info(logs ...interface{})
	error(logs ...interface{})
	warning(logs ...interface{})
	notice(logs ...interface{})
}

const (
	infoLog    = "info"
	errorLog   = "error"
	warningLog = "warning"
	noticeLog  = "notice"
)

func newLogger(store string) logger {
	switch store {
	case "file":
		return getFileLog()
	default:
		return getConsoleLog()
	}
}

type loggerBase struct {
	log *log.Logger
}

func (f *loggerBase) info(logs ...interface{}) {
	f.log.SetPrefix("[go-api] [Info]    ")
	f.log.Println(logs...)
}

func (f *loggerBase) error(logs ...interface{}) {
	f.log.SetPrefix("[go-api] [Error]   ")
	f.log.Println(logs...)
}

func (f *loggerBase) warning(logs ...interface{}) {
	f.log.SetPrefix("[go-api] [Warning] ")
	f.log.Println(logs...)
}

func (f *loggerBase) notice(logs ...interface{}) {
	f.log.SetPrefix("[go-api] [Notice]  ")
	f.log.Println(logs...)
}
