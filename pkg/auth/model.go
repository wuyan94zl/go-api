package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/jwt"
	"github.com/wuyan94zl/go-api/pkg/response"
	"github.com/wuyan94zl/mysql"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"time"
)

type User struct {
	jwt.Jwt
	Email     string    `json:"email" validate:"required||email"`
	Password  string    `json:"password" validate:"required||min:6"`
	Nickname  string    `json:"nickname" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (st *User) Create(c *gin.Context) {
	st.Email = c.DefaultPostForm("email", "")
	password, _ := bcrypt.GenerateFromPassword([]byte(c.DefaultPostForm("password", "")), bcrypt.DefaultCost)
	st.Password = string(password)
	st.Nickname = c.DefaultPostForm("nickname", "")

	err := mysql.GetInstance().Create(st)
	if err != nil {
		response.Error(500, err.Error())
	}
}

func (st *User) Update(c *gin.Context) {
	st.Email = c.DefaultPostForm("email", "")
	password, _ := bcrypt.GenerateFromPassword([]byte(c.DefaultPostForm("password", "")), bcrypt.DefaultCost)
	st.Password = string(password)
	st.Nickname = c.DefaultPostForm("nickname", "")

	err := mysql.GetInstance().Save(st)
	if err != nil {
		response.Error(500, err.Error())
	}
}

func (st *User) Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	err := mysql.GetInstance().First(st, id)
	if err != nil {
		response.Error(500, err.Error())
	}
}

func (st *User) Delete(c *gin.Context) {
	st.Info(c)
	err := mysql.GetInstance().Delete(st)
	if err != nil {
		response.Error(500, err.Error())
	}
}
