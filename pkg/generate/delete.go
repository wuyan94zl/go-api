package generate

import (
	"fmt"
	"reflect"
)

// 创建删除方法
func getDeleteFuncStr(kind reflect.Type) string {
	name := kind.Name()
	str := `
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
`
	str = fmt.Sprintf("%s\tvar %s %s\n", str, name, kind)
	str = fmt.Sprintf("%s\n\tmodel.First(&%s,%s)\n", str, name, "id")

	data := `	if %s.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}`
	data = fmt.Sprintf(data, name)
	str = fmt.Sprintf("%s%s\n", str, data)

	data = `	model.DeleteOne(&%s)
	utils.SuccessData(c, "删除成功")
}`
	data = fmt.Sprintf(data, name)
	str = fmt.Sprintf("%s%s", str, data)
	return str
}
