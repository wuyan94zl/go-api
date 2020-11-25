package admin

import "github.com/wuyan94zl/api/app/models"

type Admin struct {
	models.BaseModel
	LoginId  string
	LoginPwd string
	Name     string
}
