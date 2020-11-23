package user

import (
	"github.com/wuyan94zl/api/app/models"
	"github.com/wuyan94zl/api/pkg/database"
)

// 定义 表名 users 字段如下
type User struct {
	models.BaseModel
	Id	        int
	Email       string
	Password	string
	Name		string
	Age			int
}
// 程序初始化时执行表结构迁移
func init(){
	database.DB.AutoMigrate(&User{})
}