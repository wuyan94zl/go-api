package models

import (
	"github.com/wuyan94zl/api/pkg/types"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	Id        uint64    `gorm:"column:id;primaryKey;autoIncrement;not null"json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"json:"updated_at"`

	// 支持 gorm 软删除
	// DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" sql:"index"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.Id)
}
