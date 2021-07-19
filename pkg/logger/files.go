package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var file *fileLog

type fileLog struct {
	log *log.Logger
}

func (f *fileLog) info(logs ...interface{}) {
	f.log.SetPrefix(f.log.Prefix() + "info ")
	f.log.Println(logs...)
}

func (f *fileLog) error(logs ...interface{}) {
	f.log.SetPrefix(f.log.Prefix() + "error ")
	f.log.Println(logs...)
}

func (f *fileLog) warning(logs ...interface{}) {
	f.log.SetPrefix(f.log.Prefix() + "warning ")
	f.log.Println(logs...)
}

func (f *fileLog) notice(logs ...interface{}) {
	f.log.SetPrefix(f.log.Prefix() + "notice ")
	f.log.Println(logs...)
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
	return &fileLog{log: log.New(f, "[api-debug] ", 3)}
}
