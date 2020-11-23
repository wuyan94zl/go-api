package models

import (
	"fmt"
	"github.com/wuyan94zl/api/pkg/database"
	"github.com/wuyan94zl/api/pkg/types"
	"reflect"
	"time"
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"column:created_at;index"`
	UpdatedAt time.Time `gorm:"column:updated_at;index"`

	// 支持 gorm 软删除
	// DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" sql:"index"`
}

type PageInfo struct {
	Page     int
	PageSize int
}

type PageList struct {
	CurrentPage int
	FirstPage   int
	LastPage    int
	PageSize    int
	Total       int
	Data        interface{}
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}

func Create(data interface{}, id interface{}) {
	bType := reflect.TypeOf(data)

	d := database.DB.First(data, id)

	fmt.Println(bType, d)
}

func First(model interface{},where ...interface{})  {
	database.DB.First(model,where...)
}

/**
查询列表
*/
func Lists(model interface{}, where ...interface{}) {
	database.DB.Find(model, where...)
}

/**
* 单表分页查询
*/
func Paginate(model interface{}, pageInfo PageInfo, query string, where ...interface{}) PageList {
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	limit := pageInfo.PageSize
	rom := database.DB
	if query != ""{
		rom = rom.Where(query,where)
	}
	rom.Offset(offset).Limit(limit).Find(model)
	count := 0
	rom.Model(model).Count(&count)
	lastPage := (count / pageInfo.PageSize) + 1
	return PageList{CurrentPage: pageInfo.Page, FirstPage: 1, LastPage: lastPage, PageSize: pageInfo.PageSize, Total: count, Data: model}
}
