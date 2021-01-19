package model

import (
	"fmt"
	"github.com/wuyan94zl/api/pkg/orm"
	"strconv"
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
	Id       uint64     `json:"id"`
	ParentId uint64     `json:"parent_id"`
	Name     string     `json:"name"`
	Route    string     `json:"route"`
	Child    []treeType `json:"child"`
}

var mId uint64 = 100000000

// 获取角色权限菜单
func (role *Role) GetPermissionMenu() ([]treeType, []uint64) {
	permission := make([]uint64, 0)
	for _, v := range role.Permissions {
		permission = append(permission, v.Id)
	}

	var Menu []Menu
	orm.GetInstance().Order("parent_id").Get(&Menu, "Permissions")

	var tree []treeType
	for _, v := range Menu {
		var pId uint64 = 0
		if v.ParentId > 0 {
			pId = mId + v.ParentId
		}
		item := treeType{Id: mId + v.Id, ParentId: pId, Name: v.Name, Route: ""}
		var child []treeType
		for _, p := range v.Permissions {
			childItem := treeType{Id: p.Id, Name: fmt.Sprintf("%s（%s）", p.Name, p.Route), Route: p.Route}
			child = append(child, childItem)
		}
		item.Child = child

		tree = append(tree, item)
	}
	treeData := TreeList(tree, 0, 1)
	return treeData, permission
}

// 设置角色权限菜单
func (role *Role) SetPermissionMenu(permissionId []string) {
	where := make(map[string]interface{})
	where["id"] = orm.Where{Way: "IN", Value: permissionId}
	var permissions []Permission
	orm.GetInstance().Where(where).Get(&permissions)
	var addPermission []RolePermission
	var menuIds []uint64
	for _, v := range permissions {
		addPermission = append(addPermission, RolePermission{RoleId: role.Id, PermissionId: v.Id})
		menuIds = append(menuIds, v.MenuId)
	}
	for _, v := range permissionId {
		gid, _ := strconv.Atoi(v)
		if uint64(gid) > mId {
			menuIds = append(menuIds, uint64(gid)-mId)
		}
	}
	menus := make(map[uint64]RoleMenu)
	parentMenu(menuIds, menus, role.Id)

	var addMenu []RoleMenu
	for _, v := range menus {
		addMenu = append(addMenu, v)
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

// 菜单列表tree
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

// 菜单权限tree
func TreeList(data []treeType, pid uint64, level uint64) []treeType {
	var listTree []treeType
	for _, value := range data {
		if value.ParentId == pid {
			child := TreeList(data, value.Id, level+1)
			if len(child) > 0 {
				value.Child = child
			}
			listTree = append(listTree, value)
		}
	}
	return listTree
}

func parentMenu(menuId []uint64, menuMap map[uint64]RoleMenu, roleId uint64) {
	where := make(map[string]interface{})
	where["id"] = orm.Where{Way: "IN", Value: menuId}
	var menus []Menu
	orm.GetInstance().Where(where).Get(&menus)
	var nid []uint64
	for _, v := range menus {
		if v.ParentId > 0 {
			if _, ok := menuMap[v.ParentId]; !ok {
				nid = append(nid, v.ParentId)
			}
		}
		menuMap[v.Id] = RoleMenu{RoleId: roleId, MenuId: v.Id}
	}
	if len(nid) > 0 {
		parentMenu(nid, menuMap, roleId)
	}
}
