package model

import (
	"github.com/wuyan94zl/api/pkg/orm"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id        uint64
	Email     string    `validate:"required,min:6,email"search:"like"`
	Password  string    `validate:"min:6"pwd:"pwd"`
	Name      string    `validate:"required,min:6"search:"like"`
	Roles     []Role    `gorm:"many2many:user_roles;joinForeignKey:UserID"relationship:"manyToMany"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type UserRole struct {
	UserId uint64
	RoleId uint64
}

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

type Permission struct {
	Id          uint64
	Name        string    `validate:"required"`
	Route       string    `validate:"required"`
	MenuId      uint64    `validate:"required"`
	Description string    `validate:"required"`
	CreatedAt   time.Time `gorm:"column:created_at;index"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}

type Menu struct {
	Id          uint64
	ParentId    uint64       `validate:"required,numeric"`
	Name        string       `validate:"required"`
	Route       string       `validate:"required"`
	Description string       `validate:"required"`
	Permissions []Permission `relationship:"hasMany"`
	Menus       []Menu       `gorm:"-"`
	CreatedAt   time.Time    `gorm:"column:created_at;index"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"`
}

// 用户设置角色
func (user *User) SetRole(roleId string) {
	where := make(map[string]interface{})
	where["user_id"] = user.Id
	orm.GetInstance().Where(where).DB.Delete(&UserRole{})

	ids := strings.Split(roleId, ",")
	var userHasRole []UserRole
	for _, id := range ids {
		uid, _ := strconv.Atoi(id)
		userHasRole = append(userHasRole, UserRole{UserId: user.Id, RoleId: uint64(uid)})
	}
	orm.GetInstance().DB.Create(userHasRole)
}

// 设置角色权限菜单
func (role *Role) SetPermissionMenu(permissionId string) {
	ids := strings.Split(permissionId, ",")
	where := make(map[string]interface{})
	where["id"] = orm.Where{Way: "IN",Value: ids}

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
