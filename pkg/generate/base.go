package generate

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type mapValue struct {
	validateInfo     string
	searchInfo       string
	relationshipInfo string
}

// 获取文件位置
func getDir(name string) string {
	baseDir, _ := os.Getwd()
	dir := fmt.Sprintf("%s%s%s", baseDir, "\\app\\controllers\\", name)
	setDir(dir)
	filePath := fmt.Sprintf("%s%s%s%s%s", baseDir, "\\app\\controllers\\", name, "\\", "curd.go")
	setFile(filePath)
	return filePath
}

// 创建文件夹
func setDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(dir, 0777)
			fmt.Println("创建文件夹")
		}
	}
}

// 创建文件
func setFile(file string) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("创建文件", file)
			_, e := os.Create(file)
			fmt.Println("创建文件", e)
		}
	}
}

// 获取字段信息
func getField(rlt []map[string]mapValue, KindType reflect.Type) []map[string]mapValue {
	for i := 0; i < KindType.NumField(); i++ {
		if !KindType.Field(i).Anonymous {
			item := make(map[string]mapValue)
			item[KindType.Field(i).Name] = mapValue{validateInfo: KindType.Field(i).Tag.Get("validate"), searchInfo: KindType.Field(i).Tag.Get("search"), relationshipInfo: KindType.Field(i).Tag.Get("relationship")}
			rlt = append(rlt, item)
		}
	}
	return rlt
}

// 设置字段验证规则信息
func setValidate(data []map[string]mapValue) string {
	str := ""
	for _, mapVal := range data {
		for k, v := range mapVal {
			if v.validateInfo != "" {
				info := strings.Split(v.validateInfo, ",")
				vail := ""
				for _, i := range info {
					vail = fmt.Sprintf("%s,\"%s\"", vail, i)
				}
				str = fmt.Sprintf("%s\tdata[\"%s\"] = []string{%s} \n", str, k, vail[1:])
			}
		}
	}
	return str
}

// 获取Model创建数据
func getModelData(KindType reflect.Type) string {
	name := KindType.Name()
	str := fmt.Sprintf("\t%s := %s{", name, KindType)
	pwd := ""
	param := ""
	for i := 0; i < KindType.NumField(); i++ {
		if !KindType.Field(i).Anonymous {
			if KindType.Field(i).Tag.Get("validate") != "" {
				if KindType.Field(i).Tag.Get("pwd") != "" {
					pwd = `pwd, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("%s")), bcrypt.DefaultCost)`
					pwd = fmt.Sprintf(pwd, KindType.Field(i).Name)
					param = fmt.Sprintf("%s,%s: string(%s)", param, KindType.Field(i).Name, "pwd")
				} else {
					param = fmt.Sprintf("%s,%s: c.PostForm(\"%s\")", param, KindType.Field(i).Name, KindType.Field(i).Name)
				}
			}
		}
	}
	if pwd != "" {
		str = fmt.Sprintf("%s\n%s", pwd, str)
	}
	str = fmt.Sprintf("%s%s}", str, param[1:])
	return str
}

func getInfo(KindType reflect.Type) string {
	name := KindType.Name()
	str := "Id, _ := strconv.Atoi(c.Query(\"Id\"))\n"

	str = fmt.Sprintf("%s\tvar %sInfo %s\n", str, name, KindType)
	str = fmt.Sprintf("%s\tmodel.First(&%sInfo,Id)\n", str, name)

	return str
}
