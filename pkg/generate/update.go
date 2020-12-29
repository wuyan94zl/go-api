package generate

import (
	"fmt"
	"reflect"
)

// 创建Update方法
func getUpdateFuncStr(KindType reflect.Type, fields []map[string]mapValue) string {
	str := `
func Update(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)
`
	// 参数验证信息
	var rlt []map[string]mapValue
	rlt = append(rlt, map[string]mapValue{"id": {validateInfo: "required,min:1"}})

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
	data = getInfo(KindType)
	str = fmt.Sprintf("%s\n\t%s", str, data)
	// 数据不存在
	data = `	if %s.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}`
	data = fmt.Sprintf(data, KindType.Name())
	str = fmt.Sprintf("%s%s\n", str, data)

	// 设置更改信息
	data = getModelData(KindType, fields)
	str = fmt.Sprintf("%s\n%s", str, data)
	// 更改操作
	data = fmt.Sprintf("orm.GetInstance().Save(%s)", KindType.Name())
	str = fmt.Sprintf("%s\t%s", str, data)

	data = fmt.Sprintf("\tutils.SuccessData(c, %s) // 返回创建成功的信息\n}", KindType.Name())
	str = fmt.Sprintf("%s\n%s", str, data)
	return str
	//file.Write([]byte(str))
}
