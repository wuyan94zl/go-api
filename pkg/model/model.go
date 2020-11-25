package model

import (
	"github.com/wuyan94zl/api/pkg/database"
)


//获取一条数据
func GetFirst(model interface{},id interface{}) (interface{},error){
	if err := database.DB.First(model,id).Error; err != nil {
		return model, err
	}
	return model,nil
}

func GetOne(model interface{},condition []Condition) interface{}{
	rom := orm(condition)
	rom.First(model)
	return model
}

/**
获取多条数据
 */
func GetAll(model interface{},condition []Condition,limit ...int) interface{}{
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
	return model
}

/**
根据条件删除数据
 */
func Delete(model interface{},condition []Condition) bool{
	rom := orm(condition)
	rom.Delete(model)
	return true
}
/**
// 删除当前模型数据
 */
func DeleteOne(model interface{}) bool{
	database.DB.Delete(model)
	return true
}
/**
更新当前模型数据
 */
func UpdateOne(model interface{}) interface{}{
	database.DB.Save(model)
	return model
}

/**
单表分页查询
 */
func Paginate(model interface{}, pageInfo PageInfo, condition []Condition) PageList {
	offset := (pageInfo.Page - 1) * pageInfo.PageSize
	limit := pageInfo.PageSize
	rom := orm(condition)
	count := 0
	rom.Model(model).Count(&count)
	if count > 0 {
		rom.Offset(offset).Limit(limit).Find(model)
	}
	lastPage := (count / pageInfo.PageSize) + 1
	return PageList{CurrentPage: pageInfo.Page, FirstPage: 1, LastPage: lastPage, PageSize: pageInfo.PageSize, Total: count, Data: model}
}