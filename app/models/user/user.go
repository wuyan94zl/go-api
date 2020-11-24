package user

import (
	"github.com/wuyan94zl/api/app/models"
)

// 定义 表名 users 字段如下
type User struct {
	models.BaseModel
	Email       string
	Password	string
	Name		string
	IdString	string
}