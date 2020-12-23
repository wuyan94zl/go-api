package rbac

import (
	"fmt"
	"github.com/wuyan94zl/api/pkg/model"
	model2 "github.com/wuyan94zl/api/pkg/rbac/model"
)

func CheckPermission(id uint64,url string) bool{
	var roles []model2.Role
	var condition []model.Condition
	model.GetAll(&roles,condition,"Permissions")
	fmt.Println(roles)
	return true
}