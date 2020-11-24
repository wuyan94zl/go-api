package bootstrap

import (
	"github.com/wuyan94zl/api/app/models/admin"
	"github.com/wuyan94zl/api/app/models/user"
	"github.com/wuyan94zl/api/pkg/database"
)

func autoMigrate()  {
	database.SetMysqlDB()
	database.DB.AutoMigrate(user.User{})
	database.DB.AutoMigrate(admin.Admin{})
}