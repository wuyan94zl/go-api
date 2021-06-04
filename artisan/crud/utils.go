package crud

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func getDir(name string) string {
	baseDir, err := os.Getwd()
	if err != nil {
		return ""
	}
	return filepath.Join(baseDir, "app", "http", name)
}

func jsonData(structData *jsonStruct, name string) error {
	open, err := os.Open(filepath.Join(getDir(name), "model.json"))
	if err != nil {
		return err
	}
	defer open.Close()
	byteValue, _ := ioutil.ReadAll(open)
	err = json.Unmarshal(byteValue, structData)
	if err != nil {
		return err
	}
	return nil
}

type jsonStruct struct {
	PackageName string `json:"package_name"`
	StructName  string `json:"struct_name"`
	Fields      []struct {
		Field    string            `json:"field"`
		TypeName string            `json:"type_name"`
		Tags     map[string]string `json:"tags"`
	} `json:"fields"`
}

func (data jsonStruct) getStructFields() (string, string) {
	structFields := ""
	validateData := ""
	for _, v := range data.Fields {
		var tags []string
		for k, tag := range v.Tags {
			tags = append(tags, fmt.Sprintf("%s:\"%s\"", k, tag))
		}
		structFields = fmt.Sprintf("%s%s %s`%s`\n", structFields, v.Field, v.TypeName, strings.Join(tags, " "))
		if _, ok := v.Tags["validate"]; ok {
			validateData = fmt.Sprintf("%sst.%s = c.DefaultPostForm(\"%s\",\"\")\n", validateData, v.Field, v.Tags["json"])
		}
	}
	return structFields, validateData
}
