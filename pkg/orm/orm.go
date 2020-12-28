package orm

import (
	"github.com/wuyan94zl/api/pkg/database"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func GetInstance() *DB {
	return &DB{DB: database.DB}
}

// 设置查询关联
func (db *DB) SetRelationship(relationship []string) *DB {
	for _, v := range relationship {
		db.DB = db.DB.Preload(v)
	}
	return db
}

// 设置查询条件
func (db *DB) SetCondition(condition []Condition) *DB {
	query, values := formatQuery(condition)
	db.DB = getConditionOrm(db.DB, query, values)
	return db
}

// 设置查询分页信息
func (db *DB) SetLimit(offset int, limit int) *DB {
	db.DB = db.DB.Offset(offset).Limit(limit)
	return db
}

// 设置查询排序
func (db *DB) SetOrder(orderBy string) *DB {
	db.DB = db.DB.Order(orderBy)
	return db
}

// 创建数据
func (db *DB) Create(model interface{}) {
	db.DB.Create(model)
}

// 保存更新数据
func (db *DB) Save(model interface{}) {
	db.DB.Save(model)
}

// 删除数据
func (db *DB) Delete(model interface{}) {
	db.DB.Delete(model)
}

// 主键查询一条数据
func (db *DB) First(model interface{}, id interface{}, relationship ...string) {
	db.SetRelationship(relationship).DB.First(model, id)
}

// 查询多条数据
func (db *DB) Get(model interface{}, condition []Condition, relationship ...string) {
	db.SetCondition(condition).SetRelationship(relationship).DB.Find(model)
}

// 查询分页数据
func (db *DB) Paginate(lists *PageList, condition []Condition, relationship ...string) {
	var count int64
	db.SetCondition(condition).DB.Model(lists.Data).Count(&count)
	lists.LastPage = (count / lists.PageSize) + 1
	lists.Total = count
	if count > 0 && lists.LastPage >= lists.CurrentPage {
		offset := (lists.CurrentPage - 1) * lists.PageSize
		db.SetRelationship(relationship).SetLimit(int(offset), int(lists.PageSize)).DB.Find(lists.Data)
	}
}
