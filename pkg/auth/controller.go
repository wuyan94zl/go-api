package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/response"
	"github.com/wuyan94zl/go-api/pkg/validate"
	"github.com/wuyan94zl/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	model := GetModel()
	validateMap, validateNameMap := validate.MapDataForStruct(model)
	validateMap["password_confirmation"] = []string{"required", "min:6"}
	if ok, msg := validate.MapValidate(c.Request, validateMap, validateNameMap); !ok {
		response.Error(401, msg)
	}
	if c.PostForm("password") != c.PostForm("password_confirmation"){
		response.Error(401, "两次密码输入不一致")
	}
	where := make(map[string]interface{})
	where["email"] = c.PostForm("email")
	mysql.GetInstance().Where(where).One(&model)
	if model.Id != 0 {
		response.Error(500, "email 已存在")
	}
	model.Create(c)
	response.Success("注册成功")
}

func Login(c *gin.Context) {
	model := GetModel()
	validateMap, validateNameMap := validate.MapDataForStruct(model)
	delete(validateMap, "nickname")
	if ok, msg := validate.MapValidate(c.Request, validateMap, validateNameMap); !ok {
		response.Error(401, msg)
	}
	where := make(map[string]interface{})
	where["email"] = c.PostForm("email")
	mysql.GetInstance().Where(where).One(&model)
	if model.Id == 0 {
		response.Error(401, "email不存在或密码错误！")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(c.PostForm("password"))); err != nil {
		response.Error(401, "email不存在或密码错误！")
	}
	if loginInfo, err := model.Token(); err != nil {
		response.Error(401, "email不存在或密码错误！")
	} else {
		response.Success(loginInfo)
	}
}

func Logout(c *gin.Context) {

}

func Info(c *gin.Context) {
	authUser := GetAuthInfo(c)
	response.Success(authUser)
}
