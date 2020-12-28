package rbac

import (
	"github.com/wuyan94zl/api/pkg/orm"
	model2 "github.com/wuyan94zl/api/pkg/rbac/model"
)

func CheckPermission(id uint64,url string) bool{
	var roles []model2.Role
	var condition []orm.Condition
	orm.GetInstance().Get(&roles,condition,"Permissions")
	return true
}