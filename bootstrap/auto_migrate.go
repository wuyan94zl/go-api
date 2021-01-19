package bootstrap

import (
	"github.com/wuyan94zl/api/app/models/admin"
	"github.com/wuyan94zl/api/pkg/database"
)

var MigrateStruct map[string]interface{}

// 初始化表结构体
func init() {
	MigrateStruct = make(map[string]interface{})
	MigrateStruct["admin"] = admin.Admin{}
}

func autoMigrate() {
	database.SetMysqlDB()
	for _, v := range MigrateStruct {
		_ = database.DB.AutoMigrate(v)
	}
}
