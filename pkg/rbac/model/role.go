package model

import (
	"fmt"
	"github.com/wuyan94zl/api/pkg/orm"
	"time"
)

type Role struct {
	Id          uint64       `json:"id"gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name        string       `json:"name"validate:"required"`
	Description string       `json:"description"validate:"required"`
	Menus       []Menu       `json:"menus"gorm:"many2many:role_menus"relationship:"manyToMany"`
	Permissions []Permission `json:"permissions"gorm:"many2many:role_permissions"relationship:"manyToMany"`
	CreatedAt   time.Time    `json:"created_at"gorm:"column:created_at;index"`
	UpdatedAt   time.Time    `json:"updated_at"gorm:"column:updated_at"`
}

type RolePermission struct {
	RoleId       uint64
	PermissionId uint64
}

type RoleMenu struct {
	RoleId uint64
	MenuId uint64
}
type treeType struct {
	Id    uint64     `json:"id"`
	Name  string     `json:"name"`
	Route string     `json:"route"`
	Child []treeType `json:"child"`
}

// 获取角色权限菜单
func (role *Role) GetPermissionMenu() ([]treeType, []uint64) {
	permission :=  make([]uint64,0)
	for _, v := range role.Permissions {
		permission = append(permission, v.Id)
	}
	var Menu []Menu
	orm.GetInstance().Order("parent_id").Get(&Menu, "Permissions")

	var tree []treeType
	var mId uint64 = 10000000
	for _, v := range Menu {
		item := treeType{Id: mId, Name: v.Name, Route: ""}
		var child []treeType
		for _, p := range v.Permissions {
			childItem := treeType{Id: p.Id, Name: fmt.Sprintf("%s（%s）", p.Name, p.Route), Route: p.Route}
			child = append(child, childItem)
		}
		item.Child = child

		tree = append(tree, item)
		mId++
	}
	return tree, permission
}

// 设置角色权限菜单
func (role *Role) SetPermissionMenu(permissionId []string) {
	where := make(map[string]interface{})
	where["id"] = orm.Where{Way: "IN", Value: permissionId}
	var permissions []Permission
	orm.GetInstance().Where(where).Get(&permissions)
	mapV := make(map[uint64]uint64)
	var addPermission []RolePermission
	var addMenu []RoleMenu

	for _, v := range permissions {
		addPermission = append(addPermission, RolePermission{RoleId: role.Id, PermissionId: v.Id})
		_, ok := mapV[v.MenuId]
		if !ok {
			mapV[v.MenuId] = v.MenuId
			addMenu = append(addMenu, RoleMenu{RoleId: role.Id, MenuId: v.MenuId})
		}
	}
	orm.GetInstance().DB.Create(addPermission)
	orm.GetInstance().DB.Create(addMenu)
}

// 删除角色权限菜单
func (role *Role) DelPermissionMenu() {
	where := make(map[string]interface{})
	where["role_id"] = role.Id
	orm.GetInstance().Where(where).DB.Delete(RolePermission{})
	orm.GetInstance().Where(where).DB.Delete(RoleMenu{})
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
