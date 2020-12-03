package admin

import (
	"github.com/wuyan94zl/api/app/models"
)

type Admin struct {
	models.BaseModel
	Email    string `validate:"required,min:6,email"search:"like"`
	Password string `validate:"min:6"pwd:"pwd"`
	Name     string `validate:"required,min:6"search:"like"`
}
