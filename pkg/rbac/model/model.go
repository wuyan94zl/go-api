package model

import "time"

type User struct {
	Id        uint64
	Email     string    `validate:"required,min:6,email"search:"like"`
	Password  string    `validate:"min:6"pwd:"pwd"`
	Name      string    `validate:"required,min:6"search:"like"`
	Roles     []Role    `gorm:"many2many:user_roles"relationship:"manyToMany"`
	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
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

func RecursionMenuList(data []Menu, pid uint64, level uint64) []Menu {
	var listTree []Menu
	for _, value := range data {
		if value.ParentId == pid {
			//value.Level = level
			value.Menus = RecursionMenuList(data, value.Id, level+1)
			listTree = append(listTree, value)
		}
	}
	return listTree
}
