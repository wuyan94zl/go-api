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
	typeInfo         string
}

// 获取文件位置
func getDir(name string, uri string) string {
	baseDir, _ := os.Getwd()
	if uri != "" {
		uri = fmt.Sprintf("%s\\", uri)
	}
	dir := fmt.Sprintf("%s%s%s%s", baseDir, "\\app\\controllers\\", uri, name)
	fmt.Println(dir, uri)
	setDir(dir)
	filePath := fmt.Sprintf("%s%s%s", dir, "\\", "curd.go")
	setFile(filePath)
	return filePath
}

// 创建文件夹
func setDir(dir string) {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			_ = os.MkdirAll(dir, 0777)
		}
	}
}

// 创建文件
func setFile(file string) {
	_, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("创建文件", file)
			_, _ = os.Create(file)
		}
	}
}

// 获取字段信息
func getField(rlt []map[string]mapValue, KindType reflect.Type) []map[string]mapValue {
	for i := 0; i < KindType.NumField(); i++ {
		if !KindType.Field(i).Anonymous {
			item := make(map[string]mapValue)
			var v mapValue
			v.searchInfo = KindType.Field(i).Tag.Get("search")
			v.validateInfo = KindType.Field(i).Tag.Get("validate")
			v.relationshipInfo = KindType.Field(i).Tag.Get("relationship")
			if KindType.Field(i).Tag.Get("pwd") != "" {
				v.typeInfo = "pwd"
			} else {
				v.typeInfo = KindType.Field(i).Type.Name()
			}
			item[KindType.Field(i).Name] = v
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
				str = fmt.Sprintf("%s\tdata[\"%s\"] = []string{%s} \n", str, setToLower(k), vail[1:])
			}
		}
	}
	return str
}

// 获取Model创建数据
func getModelData(KindType reflect.Type, data []map[string]mapValue) string {
	name := KindType.Name()
	//str := fmt.Sprintf("%s := %s{", name, KindType)
	str := ""
	pwd := ""
	param := ""
	for _, mapVal := range data {
		for k, v := range mapVal {
			lowerK := setToLower(k)
			if v.typeInfo == "pwd" {
				pwd = `pwd, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("%s")), bcrypt.DefaultCost)`
				pwd = fmt.Sprintf(pwd, lowerK)
				param = fmt.Sprintf("%s\t%s.%s = string(%s)\n", param, name, k, "pwd")
			} else {
				if v.typeInfo != "string" {
					infoVal := setType(v.typeInfo, k, lowerK, name)
					param = fmt.Sprintf("%s%s\n", param, infoVal)
				} else {
					param = fmt.Sprintf("%s\t%s.%s = c.PostForm(\"%s\")\n", param, name, k, lowerK)
				}
			}
		}
	}
	if pwd != "" {
		str = fmt.Sprintf("%s\n%s", pwd, str)
	}
	str = fmt.Sprintf("%s%s", str, param)
	return str
}

func setType(typeName string, keyName string, lowerKey string, name string) string {
	switch typeName {
	case "uint64":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = uint64(%s)", str, name, keyName, keyName)
		return str
	case "uint32":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = uint32(%s)", str, name, keyName, keyName)
		return str
	case "uint16":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = uint16(%s)", str, name, keyName, keyName)
		return str
	case "uint8":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = uint8(%s)", str, name, keyName, keyName)
		return str
	case "int64":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = int64(%s)", str, name, keyName, keyName)
		return str
	case "int32":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = int32(%s)", str, name, keyName, keyName)
		return str
	case "int16":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = int16(%s)", str, name, keyName, keyName)
		return str
	case "int8":
		str := fmt.Sprintf("\t%s,_ := strconv.Atoi(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = int8(%s)", str, name, keyName, keyName)
		return str
	case "Time":
		str := fmt.Sprintf("\t%s := utils.SetStrToTime(c.PostForm(\"%s\"))\n", keyName, lowerKey)
		str = fmt.Sprintf("%s\t%s.%s = %s", str, name, keyName, keyName)
		return str
	}
	return ""
}

func setToLower(k string) string {
	str := "QWERTYUIOPASDFGHJKLZXCVBNM"
	for _, v := range str {
		if strings.Index(k, string(v)) == 0 {
			k = strings.Replace(k, string(v), strings.ToLower(string(v)), -1)
		} else if strings.Index(k, string(v)) > 0 {
			k = strings.Replace(k, string(v), fmt.Sprintf("_%s", strings.ToLower(string(v))), -1)
		}
	}
	return k
}

func getInfo(KindType reflect.Type) string {
	name := KindType.Name()
	str := "id, _ := strconv.Atoi(c.Query(\"id\"))\n"

	str = fmt.Sprintf("%s\tvar %s %s\n", str, name, KindType)
	str = fmt.Sprintf("%s\tmodel.First(&%s,id)\n", str, name)

	return str
}
