package orm

import (
	"fmt"
	"github.com/wuyan94zl/api/pkg/database"
	"github.com/wuyan94zl/api/pkg/logger"
	"gorm.io/gorm"
)

type DB struct {
	DB *gorm.DB
}

func GetInstance() *DB {
	return &DB{DB: database.DB}
}

// 设置查询关联
func (db *DB) Relationship(relationship []string) *DB {
	for _, v := range relationship {
		db.DB = db.DB.Preload(v)
	}
	return db
}

func (db *DB) getQueryValues(where map[string]interface{}) (string, []interface{}) {
	query := ""
	var values []interface{}
	for k, v := range where {
		if w, ok := v.(Where); ok {
			if query == "" {
				query = fmt.Sprintf("%s %s ?", k, w.Way)
			} else {
				query = fmt.Sprintf("%s AND %s %s ?", query, k, w.Way)
			}
			values = append(values, w.Value)
		} else {
			if query == "" {
				query = fmt.Sprintf("%s = ?", k)
			} else {
				query = fmt.Sprintf("%s AND %s = ?", query, k)
			}
			values = append(values, v)
		}
	}
	return query, values
}

func (db *DB) Where(where map[string]interface{}) *DB {
	query, values := db.getQueryValues(where)
	db.DB = db.DB.Where(query, values...)
	return db
}

func (db *DB) Or(where map[string]interface{}) *DB {
	query, values := db.getQueryValues(where)
	db.DB = db.DB.Or(query, values...)
	return db
}

// 设置查询分页信息
func (db *DB) Limit(offset int, limit int) *DB {
	db.DB = db.DB.Offset(offset).Limit(limit)
	return db
}

// 设置查询排序
func (db *DB) Order(orderBy string) *DB {
	db.DB = db.DB.Order(orderBy)
	return db
}

// 创建数据
func (db *DB) Create(model interface{}) {
	logger.SystemError(db.DB.Create(model).Error)
}

// 保存更新数据
func (db *DB) Save(model interface{}) {
	logger.SystemError(db.DB.Save(model).Error)
}

// 删除数据
func (db *DB) Delete(model interface{}) {
	logger.SystemError(db.DB.Delete(model).Error)
}

// 主键查询一条数据
func (db *DB) First(model interface{}, id interface{}, relationship ...string) {
	db.Relationship(relationship).DB.First(model, id)
}

// 查询多条数据
func (db *DB) One(model interface{}, relationship ...string) {
	db.Relationship(relationship).DB.First(model)
}

// 查询多条数据
func (db *DB) Get(model interface{}, relationship ...string) {
	db.Relationship(relationship).DB.Find(model)
}

// 查询分页数据
func (db *DB) Paginate(lists *PageList, relationship ...string) {
	if lists.PageSize == 0 {
		lists.PageSize = 15
	}
	var count int64
	db.DB.Model(lists.Data).Count(&count)
	lists.LastPage = (count / lists.PageSize) + 1
	lists.Total = count
	if count > 0 && lists.LastPage >= lists.CurrentPage {
		offset := (lists.CurrentPage - 1) * lists.PageSize
		db.Relationship(relationship).Limit(int(offset), int(lists.PageSize)).Get(lists.Data)
	}
}
