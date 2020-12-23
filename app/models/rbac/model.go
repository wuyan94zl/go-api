package rbac

import (
	"github.com/wuyan94zl/api/app/models"
	"github.com/wuyan94zl/api/pkg/rbac/model"
)

type User struct {
	model.User
	Phone string `validate:"required,min:11,max:11"search:"like"`
}

type Role struct {
	models.BaseModel
	Name        string       `validate:"required"`
	Description string       `validate:"required"`
	Menus       []Menu       `gorm:"many2many:role_menus"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}

type Permission struct {
	models.BaseModel
	Name        string `validate:"required"`
	Route       string `validate:"required"`
	MenuId      uint64 `validate:"required"`
	Description string `validate:"required"`
}

type Menu struct {
	models.BaseModel
	Name        string `validate:"required"`
	Route       string `validate:"required"`
	Description string `validate:"required"`
	Permissions []Permission
}
