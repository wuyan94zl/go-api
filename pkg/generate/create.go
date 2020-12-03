package generate

import (
	"fmt"
	"os"
	"reflect"
)

// 创建Create方法
func setCreate(file *os.File, kind reflect.Type) string{
	str := `
func Create(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)
`
	var rlt []map[string]mapValue
	fields := getField(rlt, kind)
	data := setValidate(fields)
	str = fmt.Sprintf("%s\n%s", str, data)

	data = `	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}`
	str = fmt.Sprintf("%s\n%s", str, data)

	data = getModelData(kind)
	str = fmt.Sprintf("%s\n\t%s", str, data)

	data = `
	model.Create(&%s)
	utils.SuccessData(c, %s) // 返回创建成功的信息
}`
	data = fmt.Sprintf(data, kind.Name(), kind.Name())
	str = fmt.Sprintf("%s\n%s", str, data)
	return str
	//file.Write([]byte(str))
}
