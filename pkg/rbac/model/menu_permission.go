package model

import (
	"time"
)

type Permission struct {
	Id          uint64    `json:"id"gorm:"column:id;primaryKey;autoIncrement;not null"`
	Name        string    `json:"name"validate:"required"`
	Route       string    `gorm:"index"validate:"required"json:"route"`
	MenuId      uint64    `validate:"required"json:"menu_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"json:"updated_at"`
	IsHas       bool      `gorm:"-"json:"is_has"`
}

type Menu struct {
	Id          uint64       `json:"id"gorm:"column:id;primaryKey;autoIncrement;not null"`
	ParentId    uint64       `validate:"required,numeric"json:"parent_id"`
	Name        string       `json:"name"validate:"required"`
	Icon        string       `validate:"required"json:"icon"`
	Route       string       `validate:"required"json:"route"`
	Sort        uint64       `json:"sort"`
	Permissions []Permission `relationship:"hasMany"json:"permissions"`
	Menus       []Menu       `gorm:"-"json:"menus"`
	CreatedAt   time.Time    `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at"json:"updated_at"`
}
