package model

import (
	"github.com/wuyan94zl/api/pkg/jwt"
	"github.com/wuyan94zl/api/pkg/orm"
	"strconv"
	"strings"
)

type User struct {
	jwt.Jwt
	Email    string `json:"email"gorm:"unique"validate:"required,min:6,email"search:"like"`
	Password string `json:"-"validate:"min:6"pwd:"pwd"`
	Name     string `json:"name"validate:"required,min:6"search:"like"`
	Roles    []Role `json:"roles"gorm:"many2many:user_roles;joinForeignKey:UserID"relationship:"manyToMany"`
}

type UserRole struct {
	UserId uint64
	RoleId uint64
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

func (user *User) Menus() []Menu {
	var Menus []Menu

	var roleIds []uint64
	for _, v := range user.Roles {
		roleIds = append(roleIds, v.Id)
	}
	var roleMenus []RoleMenu
	where := make(map[string]interface{})
	where["role_id"] = orm.Where{Way: "IN", Value: roleIds}
	orm.GetInstance().Where(where).Get(&roleMenus)

	var menuIds []uint64
	for _, v := range roleMenus {
		menuIds = append(menuIds, v.MenuId)
	}

	mWhere := make(map[string]interface{})
	mWhere["id"] = orm.Where{Way: "IN", Value: menuIds}
	orm.GetInstance().Where(mWhere).Order("sort ASC").Get(&Menus)

	tree := RecursionMenuList(Menus, 0, 1)

	return tree
}
