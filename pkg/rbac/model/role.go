package model

import (
	"github.com/wuyan94zl/api/pkg/orm"
	"strings"
	"time"
)

type Role struct {
	Id          uint64
	Name        string       `validate:"required"`
	Description string       `validate:"required"`
	Menus       []Menu       `gorm:"many2many:role_menus"relationship:"manyToMany"`
	Permissions []Permission `gorm:"many2many:role_permissions"relationship:"manyToMany"`
	CreatedAt   time.Time    `gorm:"column:created_at;index"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"`
}

type RoleHasPermission struct {
	RoleId       uint64
	PermissionId uint64
}

type RoleHasMenu struct {
	RoleId uint64
	MenuId uint64
}

// 设置角色权限菜单
func (role *Role) SetPermissionMenu(permissionId string) {
	ids := strings.Split(permissionId, ",")
	where := make(map[string]interface{})
	where["id"] = orm.Where{Way: "IN", Value: ids}

	var permissions []Permission
	orm.GetInstance().Where(where).Get(&permissions)
	mapV := make(map[uint64]uint64)
	var addPermission []RoleHasPermission
	var addMenu []RoleHasMenu
	for _, v := range permissions {
		addPermission = append(addPermission, RoleHasPermission{RoleId: role.Id, PermissionId: v.Id})
		_, ok := mapV[v.MenuId]
		if !ok {
			mapV[v.MenuId] = v.MenuId
			addMenu = append(addMenu, RoleHasMenu{RoleId: role.Id, MenuId: v.MenuId})
		}
	}
	orm.GetInstance().DB.Create(addPermission)
	orm.GetInstance().DB.Create(addMenu)
}

// 删除角色权限菜单
func (role *Role) DelPermissionMenu() {
	where := make(map[string]interface{})
	where["role_id"] = role.Id
	orm.GetInstance().Where(where).DB.Delete(RoleHasPermission{})
	orm.GetInstance().Where(where).DB.Delete(RoleHasMenu{})
}

func RecursionMenuList(data []Menu, pid uint64, level uint64) []Menu {
	var listTree []Menu
	for _, value := range data {
		if value.ParentId == pid {
			value.Menus = RecursionMenuList(data, value.Id, level+1)
			listTree = append(listTree, value)
		}
	}
	return listTree
}
