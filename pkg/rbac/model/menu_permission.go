package model

import (
	"time"
)

type Permission struct {
	Id          uint64
	Name        string    `validate:"required"`
	Route       string    `gorm:"index"validate:"required"`
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
