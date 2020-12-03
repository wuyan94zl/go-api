package model

import (
	"github.com/wuyan94zl/api/pkg/database"
)

// 创建一条数据
func Create(model interface{}){
	database.DB.Create(model)
}

/** start查询相关 */

// 主键获取一条数据
func First(model interface{},id interface{}){
	database.DB.First(model,id)
}

// 条件获取一条数据
func GetOne(model interface{},condition []Condition){
	rom := orm(condition)
	rom.First(model)
}

// 获取多条数据
func GetAll(model interface{},condition []Condition,limit ...int){
	rom := orm(condition)
	if limit != nil{
		switch len(limit) {
		case 1:
			rom = rom.Limit(limit[0])
		case 2:
			rom = rom.Offset(limit[1]).Limit(limit[0])
		}
	}
	rom.Find(model)
}

// 单表分页查询
func Paginate(model interface{}, pageInfo PageInfo, condition []Condition) PageList {
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	limit := pageInfo.PageSize
	rom := orm(condition)
	var count int64
	rom.Model(model).Count(&count)
	if count > 0 {
		rom.Offset(int(offset)).Limit(int(limit)).Find(model)
	}
	lastPage := (count / pageInfo.PageSize) + 1
	return PageList{CurrentPage: pageInfo.Page, FirstPage: 1, LastPage: lastPage, PageSize: pageInfo.PageSize, Total: count, Data: model}
}

/** end查询相关 */

// 根据条件删除数据
func Delete(model interface{},condition []Condition) bool{
	rom := orm(condition)
	rom.Delete(model)
	return true
}

// 删除当前模型数据
func DeleteOne(model interface{}) bool{
	database.DB.Delete(model)
	return true
}

// 删除当前模型数据
func DeleteById(model interface{},Id interface{}){
	database.DB.Delete(model,Id)
}

/**
更新当前模型数据
 */
func UpdateOne(model interface{}) interface{}{
	database.DB.Save(model)
	return model
}
