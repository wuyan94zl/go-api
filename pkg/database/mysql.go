package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/wuyan94zl/api/pkg/config"
	"strconv"
	"time"
)

var DB *gorm.DB // 定义 mysql 连接实例
var errDb error

// 初始化 mysql DB 连接实例
func SetMysqlDB() {
	// 单例模式获取数据库连接 实例
	var (
		host              = config.GetString("database.mysql.host")
		port              = config.GetString("database.mysql.port")
		database          = config.GetString("database.mysql.database")
		username          = config.GetString("database.mysql.username")
		password          = config.GetString("database.mysql.password")
		charset           = config.GetString("database.mysql.charset")
		maxConnect, _     = strconv.Atoi(config.GetString("database.mysql.max_open_connections"))
		maxIdleConnect, _ = strconv.Atoi(config.GetString("database.mysql.max_idle_connections"))
		maxLifeSeconds, _ = strconv.Atoi(config.GetString("database.mysql.max_life_seconds"))
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database, charset, true, "Local")
	DB, errDb = gorm.Open("mysql", dsn)
	sqlDB := DB.DB()
	sqlDB.SetMaxOpenConns(maxConnect)     // 设置最大连接数
	sqlDB.SetMaxIdleConns(maxIdleConnect) //
	sqlDB.SetConnMaxLifetime(time.Duration(maxLifeSeconds))
	if errDb != nil {
		panic(errDb)
	}
	DB.LogMode(true) // 开启打印sql 日志
}
