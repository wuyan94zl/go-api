package logger

var store = "file"

func Info(message ...interface{}) {
	newLogger(store).info(message...)
}

func Error(message ...interface{}) {
	newLogger(store).error(message...)
}

func Notice(message ...interface{}) {
	newLogger(store).notice(message...)
}

func Warning(message ...interface{}) {
	newLogger(store).warning(message...)
}

func InfoToStore(store string, message ...interface{}) {
	newLogger(store).info(message...)
}

func ErrorToStore(store string, message ...interface{}) {
	newLogger(store).error(message...)
}

func NoticeToStore(store string, message ...interface{}) {
	newLogger(store).notice(message...)
}

func WarningToStore(store string, message ...interface{}) {
	newLogger(store).warning(message...)
}
