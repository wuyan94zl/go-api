package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var file *fileLog

type fileLog struct {
	file *os.File
}

func (f *fileLog) info(logs ...interface{}) {
	log.New(f.file, "", 0).Println(logPrefix(infoLog, logs...))
}

func (f *fileLog) error(logs ...interface{}) {
	log.New(f.file, "", 0).Println(logPrefix(errorLog, logs...))
}

func (f *fileLog) warning(logs ...interface{}) {
	log.New(f.file, "", 0).Println(logPrefix(warningLog, logs...))
}

func (f *fileLog) notice(logs ...interface{}) {
	log.New(f.file, "", 0).Println(logPrefix(noticeLog, logs...))
}

func getFileLog() *fileLog {
	if file != nil {
		return file
	} else {
		return logFile(fullDir())
	}
}

func fullDir() string {
	dir, _ := os.Getwd()
	fPath := filepath.Join(dir, "logs", time.Now().Format("200601"))
	if _, err := os.Stat(fPath); err != nil {
		os.MkdirAll(fPath, 0777)
	}
	return fPath
}

func logFile(dir string) *fileLog {
	filePath := filepath.Join(dir, time.Now().Format("20060102")+".log")
	_, err := os.Stat(filePath)
	var f *os.File
	if err != nil {
		f, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	} else {
		f, _ = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0777)
	}
	return &fileLog{file: f}
}
