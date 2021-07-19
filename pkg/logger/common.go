package logger

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
