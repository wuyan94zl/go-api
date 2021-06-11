package file_tools

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"time"
)

func fullDir(name string) string {
	dir, _ := os.Getwd()
	fPath := filepath.Join(dir, "upload", "export", name, time.Now().Format("200601"))
	if _, err := os.Stat(fPath); err != nil {
		os.MkdirAll(fPath, 0777)
	}
	return fPath
}

func ExportCsv(fileName string, data [][]string, isEnd bool) string {
	fullFir := fullDir("csv")
	filePath := filepath.Join(fullFir, fileName)
	_, err := os.Stat(filePath)
	var file *os.File
	var w *csv.Writer
	if err != nil {
		file, _ = os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0777)
		file.WriteString("\xEF\xBB\xBF")
	} else {
		file, _ = os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0777)
	}
	defer file.Close()
	w = csv.NewWriter(file)
	for _, v := range data {
		w.Write(v)
	}
	w.Flush()
	if isEnd{

	}
	return filePath
}
