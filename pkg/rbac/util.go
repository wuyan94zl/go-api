package rbac

import (
	"github.com/wuyan94zl/api/pkg/orm"
)

func CheckPermission(id uint64, url string) bool {
	rolePermission := "join role_permissions on role_permissions.permission_id = permissions.id"
	userRole := "join user_roles on user_roles.role_id = role_permissions.role_id"
	query := "user_roles.user_id = ? AND permissions.route = ?"
	type result struct {
		Id uint64
	}
	r := result{}
	orm.GetInstance().DB.Table("permissions").Select("permissions.id").Joins(rolePermission).Joins(userRole).Where(query, id, url).Scan(&r)
	if r.Id > 0 {
		return true
	}
	return false
}
