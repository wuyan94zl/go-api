package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wuyan94zl/api/pkg/config"
)

var DB *gorm.DB // 定义 mysql 连接实例
var errDb error

// 初始化 mysql DB 连接实例
func init() {
	// 单例模式获取数据库连接 实例
	var (
		host     = config.GetString("database.mysql.host")
		port     = config.GetString("database.mysql.port")
		database = config.GetString("database.mysql.database")
		username = config.GetString("database.mysql.username")
		password = config.GetString("database.mysql.password")
		charset  = config.GetString("database.mysql.charset")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")
	DB, errDb = gorm.Open("mysql", dsn)
	if errDb != nil {
		panic(errDb)
	}
	DB.LogMode(true) // 开启打印sql 日志
	// defer DB.close() // 持久连接 就不需要关闭了
}
