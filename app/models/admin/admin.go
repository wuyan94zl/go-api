package admin

import "github.com/wuyan94zl/api/pkg/jwt"

type Admin struct {
	jwt.Jwt
	Email    string `json:"email"gorm:"unique"validate:"required,min:6,email"search:"like"`
	Password string `json:"-"validate:"min:6"pwd:"pwd"`
	Name     string `json:"name"validate:"required,min:6"search:"like"`
	Phone    string `validate:"required,min:11,max:11"search:"="json:"phone"`
}
