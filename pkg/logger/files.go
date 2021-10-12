package logger

import (
	"log"
	"os"
	"path/filepath"
	"time"
)

var file *fileLog

type fileLog struct {
	path string
	loggerBase
}

func getFileLog() *fileLog {
	return logFile(fullDir())
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
	if file != nil && file.path == filePath {
		return file
	}
	_, err := os.Stat(filePath)
	var f *os.File
	if err != nil {
		f, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
	} else {
		f, _ = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0777)
	}
	file = &fileLog{path: filePath}
	file.log = log.New(f, "", 3)
	return file
}
