package admin

import (
	"fmt"
	"github.com/wuyan94zl/api/pkg/database"
)

type Admin struct {
	LoginId  string
	LoginPwd string
	Name     string
}

func init(){
	fmt.Println("admin automigreate")
	database.DB.AutoMigrate(&Admin{})
}
