package rbac

import (
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/rbac/model"
)

func CheckPermission(id uint64,url string) bool{
	var roles []model.Role
	orm.GetInstance().Get(&roles,"Permissions")
	return true
}