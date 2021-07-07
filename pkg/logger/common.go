package logger

import "time"

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
	}
	return nil
}

func logPrefix(logType string, data ...interface{}) []interface{} {
	var logs []interface{}
	logs = append(logs, "[go-api logger]", logType, time.Now().Format("2006-01-02 15:04:05"))
	logs = append(logs, data...)
	return logs
}
