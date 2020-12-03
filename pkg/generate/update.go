package generate

import (
	"fmt"
	"os"
	"reflect"
)

// 设置更新信息
func setInfo(KindType reflect.Type) string {
	name := KindType.Name()
	param := ""
	for i := 0; i < KindType.NumField(); i++ {
		if !KindType.Field(i).Anonymous {
			if KindType.Field(i).Tag.Get("validate") != "" {
				param = fmt.Sprintf("%s\t%sInfo.%s = c.PostForm(\"%s\")\n", param, name, KindType.Field(i).Name, KindType.Field(i).Name)
			}
		}
	}
	return param
}

// 创建Update方法
func getUpdateFuncStr(file *os.File, kind reflect.Type) string{
	str := `
func Update(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)
`
	// 参数验证信息
	var rlt []map[string]mapValue
	rlt = append(rlt, map[string]mapValue{"Id": {validateInfo: "required,min:1"}})
	fields := getField(rlt, kind)
	data := setValidate(fields)
	str = fmt.Sprintf("%s\n%s", str, data)
	// 参数验证信息
	data = `	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}`
	str = fmt.Sprintf("%s\n%s", str, data)
	// 查询信息
	data = getInfo(kind)
	str = fmt.Sprintf("%s\n\t%s", str, data)
	// 数据不存在
	data = `	if %sInfo.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}`
	data = fmt.Sprintf(data, kind.Name())
	str = fmt.Sprintf("%s%s\n", str, data)

	// 设置更改信息
	data = setInfo(kind)
	str = fmt.Sprintf("%s\n%s", str, data)
	// 更改操作
	data = fmt.Sprintf("model.UpdateOne(%sInfo)", kind.Name())
	str = fmt.Sprintf("%s\n\t%s", str, data)

	data = `
	utils.SuccessData(c, %sInfo) // 返回创建成功的信息
}`
	data = fmt.Sprintf(data, kind.Name())
	str = fmt.Sprintf("%s\n%s", str, data)
	return str
	//file.Write([]byte(str))
}
