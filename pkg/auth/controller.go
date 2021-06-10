package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/utils"
	"github.com/wuyan94zl/go-api/pkg/validate"
	"github.com/wuyan94zl/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	model := GetModel()
	validateMap, validateNameMap := validate.MapDataForStruct(model)
	validateMap["password_confirmation"] = []string{"required", "min:6"}
	if ok, msg := validate.MapValidate(c.Request, validateMap, validateNameMap); !ok {
		utils.SuccessErr(c, 401, msg)
		return
	}
	where := make(map[string]interface{})
	where["email"] = c.PostForm("email")
	mysql.GetInstance().Where(where).One(&model)
	if model.Id != 0 {
		utils.SuccessErr(c, 500, "email 已存在")
		return
	}
	if create, err := model.Create(c); !create {
		utils.SuccessErr(c, 500, err)
		return
	}
	utils.SuccessData(c, "注册成功")
}

func Login(c *gin.Context) {
	model := GetModel()
	validateMap, validateNameMap := validate.MapDataForStruct(model)
	delete(validateMap, "nickname")
	if ok, msg := validate.MapValidate(c.Request, validateMap, validateNameMap); !ok {
		utils.SuccessErr(c, 401, msg)
		return
	}
	where := make(map[string]interface{})
	where["email"] = c.PostForm("email")
	mysql.GetInstance().Where(where).One(&model)
	if model.Id == 0 {
		utils.SuccessErr(c, 401, "email不存在或密码错误！")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(c.PostForm("password"))); err != nil {
		utils.SuccessErr(c, 401, "email不存在或密码错误！")
		return
	}
	if loginInfo, err := model.Token(); err != nil {
		utils.SuccessErr(c, 401, "email不存在或密码错误！")
		return
	} else {
		utils.SuccessData(c, loginInfo)
	}
}

func Logout(c *gin.Context) {

}

func Info(c *gin.Context) {
	authUser := GetAuthInfo(c)
	if authUser.Id == 0 {
		utils.SuccessErr(c, 500, "用户不存在")
		return
	}
	GetAuthInfo(c)
	GetAuthInfo(c)
	utils.SuccessData(c, authUser)
}
