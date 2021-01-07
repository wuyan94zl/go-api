package admin

import (
	"github.com/wuyan94zl/api/pkg/rbac/model"
)

type Admin struct {
	model.User
	Phone string `validate:"required,min:11,max:11"search:"="json:"phone"`
}
