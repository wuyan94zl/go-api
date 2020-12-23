package generate

import "fmt"

type relationship struct {
	name       string
	structType string
	tag        string
}

var relationshipInfo []relationship

// 查询 关联
func getRelationshipStr() string {
	relationshipStr := ""
	if len(relationshipInfo) > 0 {
		for _, v := range relationshipInfo {
			relationshipStr = fmt.Sprintf(`%s%s"%s"`, relationshipStr, ",", v.name)
		}
	}
	return relationshipStr
}

// 创建更新 关联
func setRelationshipStr() string {
	relationshipStr := ""
	if len(relationshipInfo) > 0 {
		for _, v := range relationshipInfo {
			switch v.tag {
			case "hasOne":
			case "belongsTo":
			case "hasMany":
			case "manyToMany":
				relationshipStr = fmt.Sprintf("%s%s", relationshipStr, setRelationshipManyToManyOne())
			}
		}
	}
	return relationshipStr
}

func setRelationshipManyToManyOne() string {
	relationshipStr := ""

	return relationshipStr
}
