package utils

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"text/template"
)

type FileGenConfig struct {
	Dir          string
	Filename     string
	TemplateFile string
	Data         interface{}
}

func GenFile(c FileGenConfig) error {
	fp, created, err := MaybeCreteFile(c.Dir, c.Filename)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	defer fp.Close()
	text := c.TemplateFile
	var t = template.Must(template.New("name").Parse(text))
	buffer := new(bytes.Buffer)
	err = t.Execute(buffer, c.Data)
	if err != nil {
		return err
	}
	code := FormatCode(buffer.String())
	_, err = fp.WriteString(code)
	return err
}

func FormatCode(code string) string {
	ret, err := format.Source([]byte(code))
	if err != nil {
		return code
	}
	return string(ret)
}

func MaybeCreteFile(dir string, fileName string) (fp *os.File, created bool, err error) {
	if _, err = os.Stat(dir); err != nil {
		os.MkdirAll(dir, 0777)
	}
	fPath := filepath.Join(dir, fileName)
	_, err = os.Stat(fPath)
	if err == nil {
		return nil, false, nil
	}
	fp, err = os.Create(fPath)
	created = err == nil
	return
}

func GetDir(method string, name string) string {
	baseDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(baseDir, "app", method, name)
}
