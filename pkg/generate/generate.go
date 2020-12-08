package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// 设置curd控制器
func SetCurd(kind interface{},uri string) {
	kindType := reflect.TypeOf(kind)
	// 打开文件操作流
	name := strings.ToLower(kindType.Name())
	dir := getDir(name,uri)
	file, err := os.OpenFile(dir, os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	os.Truncate(dir,0)

	// package
	pkgStr := fmt.Sprintf("package %s%s", name, "\n")
	// import
	impStr := getImportStr(kindType.PkgPath())
	var fields []map[string]mapValue
	fields = getField(fields,kindType)
	// create
	createStr := getCreateFuncStr(kindType,fields)
	// 有密码字段 import 增加 bcrypt包
	n := strings.Index(createStr,"bcrypt.GenerateFromPassword")
	if n != -1{
		impStr = strings.Replace(impStr,")","	\"golang.org/x/crypto/bcrypt\"\n)",1)
	}
	// update func
	updateStr := getUpdateFuncStr(kindType,fields)
	// delete func
	deleteStr := getDeleteFuncStr(kindType)
	// info func
	infoStr := getInfoFuncStr(kindType)
	// Paginate func
	paginateStr := getPaginateFuncStr(kindType,fields)

	// 合并
	rightStr := fmt.Sprintf("%s%s%s%s%s%s%s",pkgStr,impStr,createStr,updateStr,deleteStr,infoStr,paginateStr)
	// 写入
	_, err = file.Write([]byte(rightStr))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("write file successful")
}

// 设置路由
func SetRoute(kind interface{},uri string,pkgUri string)  {
	name := strings.ToLower(reflect.TypeOf(kind).Name())
	dir, _ := os.Getwd()
	if uri == ""{
		uri = "api"
	}
	filePath := fmt.Sprintf("%s\\routes\\%s.go",dir,uri)
	data,err := ioutil.ReadFile(filePath)
	if err != nil {
		panic("文件不存在")
	}
	dataString := string(data)

	// 设置关键字
	setKeyword := fmt.Sprintf("// start %s",name)
	num := strings.Index(dataString,setKeyword)
	// 找到设置关键字 跳过设置
	if num != -1{
		return
	}

	// 写入路由信息
	addStr := `
	// start %s
	router.POST("/%s/create",%s.Create)
	router.POST("/%s/update",%s.Update)
	router.GET("/%s/delete",%s.Delete)
	router.GET("/%s/info",%s.Info)
	router.POST("/%s/paginate",%s.Paginate)
	// end %s
}
`
	addStr = fmt.Sprintf(addStr,name,name,name,name,name,name,name,name,name,name,name,name)
	dataString = strings.Replace(dataString,"}",addStr,1)

	// 包路径
	if pkgUri != ""{
		pkgUri = fmt.Sprintf("%s/",pkgUri)
	}
	pkgPath := fmt.Sprintf("%s%s%s","github.com/wuyan94zl/api/app/controllers/",pkgUri,name)
	num = strings.Index(dataString,pkgPath)
	// 没有找到包路径，添加包路径
	if num == -1{
		pkg := fmt.Sprintf("%s\n\t\"%s\"","import (",pkgPath)
		dataString = strings.Replace(dataString,"import (",pkg,1)
	}
	file, err := os.OpenFile(filePath, os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte(dataString))
	fmt.Println("路由写入成功")
}
