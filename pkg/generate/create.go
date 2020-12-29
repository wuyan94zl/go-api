package generate

import (
	"fmt"
	"reflect"
)

// 创建Create方法
func getCreateFuncStr(KindType reflect.Type, fields []map[string]mapValue) string {
	str := `
func Create(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)
`
	data := setValidate(fields)
	str = fmt.Sprintf("%s\n%s", str, data)

	data = `	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}`
	str = fmt.Sprintf("%s\n%s\n", str, data)

	str = fmt.Sprintf("%s\tvar %s %s", str, KindType.Name(), KindType)

	data = getModelData(KindType, fields)
	str = fmt.Sprintf("%s\n%s", str, data)
	setRelationshipStr()
	data = `	orm.GetInstance().Create(&%s)
	utils.SuccessData(c, %s) // 返回创建成功的信息
}`
	data = fmt.Sprintf(data, KindType.Name(), KindType.Name())
	str = fmt.Sprintf("%s%s", str, data)
	return str
}
