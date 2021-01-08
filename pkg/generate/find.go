package generate

import (
	"fmt"
	"reflect"
)

// 设置查询条件
func setSearch(data []map[string]mapValue) string {
	str := ""
	for _, mapVal := range data {
		for k, v := range mapVal {
			lowerK := setToLower(k)
			if v.searchInfo != "" {
				str = fmt.Sprintf("%s\t%s := c.PostForm(\"%s\")\n", str, k, lowerK)
				str = fmt.Sprintf("%s\tif %s != \"\" {\n", str, k)
				switch v.searchInfo {
				case "=":
					str = fmt.Sprintf("%s\t\twhere[\"%s\"] = %s\n", str, lowerK, k)
				case ">":
					str = fmt.Sprintf("%s\t\twhere[\"%s\"] = orm.Where{Way: \">\", Value:%s}\n", str, lowerK, k)
				case "<":
					//str = fmt.Sprintf("%s\t\tconditions = model.SetCondition(conditions,\"%s\",%s,\"<\")\n", str, lowerK, k)
					str = fmt.Sprintf("%s\t\twhere[\"%s\"] = orm.Where{Way: \"<\", Value:%s}\n", str, lowerK, k)
				case "!=":
					//str = fmt.Sprintf("%s\t\tconditions = model.SetCondition(conditions,\"%s\",%s,\"!=\")\n", str, lowerK, k)
					str = fmt.Sprintf("%s\t\twhere[\"%s\"] = orm.Where{Way: \"!=\", Value:%s}\n", str, lowerK, k)
				case "like":
					v := fmt.Sprintf("fmt.Sprintf(\"%s%s\", %s, \"%s\")", "%s", "%s", k, "%")
					//str = fmt.Sprintf("%s\t\tconditions = model.SetCondition(conditions,\"%s\",%s,\"like\")\n", str, lowerK, v)
					str = fmt.Sprintf("%s\t\twhere[\"%s\"] = orm.Where{Way: \"LIKE\",Value:%s}\n", str, lowerK, v)
				}
				str = fmt.Sprintf("%s\t}\n", str)
			}
		}
	}
	return str
}

// 创建分类列表查询方法
func getPaginateFuncStr(kind reflect.Type, fields []map[string]mapValue) string {
	str := `
func Paginate(c *gin.Context) {
	where := make(map[string]interface{})
`
	data := setSearch(fields)
	str = fmt.Sprintf("%s%s\n", str, data)

	str = fmt.Sprintf("%s\tvar %s []%s\n", str, kind.Name(), kind)

	data = "	page, _ := strconv.Atoi(c.DefaultQuery(\"page\", \"1\"))"
	str = fmt.Sprintf("%s%s\n", str, data)
	relationshipStr := getRelationshipStr()
	str = fmt.Sprintf("\t%s%sorm.SetPageList(&%s, int64(page))\n", str, "	lists := ", kind.Name())
	str = fmt.Sprintf("%s\torm.GetInstance().Where(where).Paginate(lists%s)\n", str, relationshipStr)

	str = fmt.Sprintf("%s\tutils.SuccessData(c, lists)\n}", str)
	return str
}

// 创建详细数据方法
func getInfoFuncStr(kind reflect.Type) string {
	str := `
func Info(c *gin.Context) {
`
	// 查询信息
	data := getInfo(kind)
	str = fmt.Sprintf("%s\t%s\n", str, data)
	name := kind.Name()
	str = fmt.Sprintf("%s	utils.SuccessData(c, %s)\n}", str, name)
	return str
}
