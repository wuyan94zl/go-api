package model

import (
	"github.com/wuyan94zl/api/pkg/database"
	"gorm.io/gorm"
)
// 构建relationship查询
func setRelationship(rom *gorm.DB, relationship []string) *gorm.DB {
	for _, v := range relationship {
		rom = rom.Preload(v)
	}
	return rom
}

// 创建数据
func Create(model interface{}) {
	database.DB.Create(model)
}

/** start查询相关 */

// 主键获取一条数据
func First(model interface{}, id interface{}, relationship ...string) {
	rom := database.DB
	rom = setRelationship(rom, relationship)
	rom.First(model, id)
}

// 条件获取一条数据
func GetOne(model interface{}, condition []Condition, relationship ...string) {
	rom := orm(condition)
	rom = setRelationship(rom, relationship)
	rom.First(model)
}

// 获取limit部分数据
func GetLimit(model interface{}, condition []Condition, limit int, offset int, relationship ...string) {
	rom := orm(condition)
	rom = rom.Offset(offset).Limit(limit)
	rom = setRelationship(rom, relationship)
	rom.Find(model)
}

// 获取所有数据
func GetAll(model interface{}, condition []Condition, relationship ...string) {
	rom := orm(condition)
	rom = setRelationship(rom, relationship)
	rom.Find(model)
}

// 单表分页查询
func Paginate(model interface{}, pageInfo PageInfo, condition []Condition, relationship ...string) PageList {
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	limit := pageInfo.PageSize
	rom := orm(condition)
	var count int64
	rom.Model(model).Count(&count)
	if count > 0 {
		rom = rom.Offset(int(offset)).Limit(int(limit))
		rom = setRelationship(rom, relationship)
		rom.Find(model)
	}
	lastPage := (count / pageInfo.PageSize) + 1
	return PageList{CurrentPage: pageInfo.Page, FirstPage: 1, LastPage: lastPage, PageSize: pageInfo.PageSize, Total: count, Data: model}
}

/** end查询相关 */

// 根据条件删除数据
func Delete(model interface{}, condition []Condition) bool {
	rom := orm(condition)
	rom.Delete(model)
	return true
}

// 删除当前模型数据
func DeleteOne(model interface{}) bool {
	database.DB.Delete(model)
	return true
}

// 删除当前模型数据
func DeleteById(model interface{}, Id interface{}) {
	database.DB.Delete(model, Id)
}

/**
更新当前模型数据
*/
func UpdateOne(model interface{}) interface{} {
	database.DB.Save(model)
	return model
}
