package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

// 设置curd控制器
func SetCurd(kind interface{}) {
	kindType := reflect.TypeOf(kind)

	// 打开文件操作流
	name := strings.ToLower(kindType.Name())
	dir := getDir(name)
	file, err := os.OpenFile(dir, os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	os.Truncate(dir,0)

	// package
	pkgStr := fmt.Sprintf("package %s%s", name, "\n")
	// import
	impStr := getImportStr(name)
	// create
	createStr := getCreateFuncStr(file, kindType)
	// 有密码字段 import 增加 bcrypt包
	n := strings.Index(createStr,"bcrypt.GenerateFromPassword")
	if n != -1{
		impStr = strings.Replace(impStr,")","	\"golang.org/x/crypto/bcrypt\"\n)",1)
	}
	// update func
	updateStr := getUpdateFuncStr(file, kindType)
	// delete func
	deleteStr := getDeleteFuncStr(file, kindType)
	// info func
	infoStr := getInfoFuncStr(file,kindType)
	// Paginate func
	paginateStr := getPaginateFuncStr(file, kindType)

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
func SetRoute(kind interface{})  {
	name := strings.ToLower(reflect.TypeOf(kind).Name())
	dir, _ := os.Getwd()
	filePath := fmt.Sprintf("%s%s",dir,"\\routes\\api.go")
	data,_ := ioutil.ReadFile(filePath)
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
	api.POST("/%s/create",%s.Create)
	api.POST("/%s/update",%s.Update)
	api.GET("/%s/delete",%s.Delete)
	api.GET("/%s/info",%s.Info)
	api.POST("/%s/paginate",%s.Paginate)
	// end %s
}
`
	addStr = fmt.Sprintf(addStr,name,name,name,name,name,name,name,name,name,name,name,name)
	dataString = strings.Replace(dataString,"}",addStr,1)

	// 包路径
	pkgPath := fmt.Sprintf("%s%s","github.com/wuyan94zl/api/app/controllers/",name)
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
