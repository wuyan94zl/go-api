package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/admin"
	"github.com/wuyan94zl/api/pkg/auth"
	"github.com/wuyan94zl/api/pkg/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// 登录
func Login(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)
	data["email"] = []string{"required","min:6","email"}
	data["password"] = []string{"required","between:6,20"}
	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	// 查询获取到用户
	email := c.PostForm("email")
	password := c.PostForm("password")
	var condition []model.Condition
	condition = model.SetCondition(condition,"email",email)
	info := admin.Admin{}
	model.GetOne(&info, condition)
	if info.Id == 0 {
		utils.SuccessErr(c, -1000, "用户名或密码错误")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password)) != nil {
		utils.SuccessErr(c, -1000, "用户名或密码错误")
		return
	}
	// 换取token
	token, err := auth.GetToken(&info)
	if err != nil {
		utils.SuccessErr(c, -1000, "未知错误")
	} else {
		utils.SuccessData(c, token)
	}
}

// 获取登录用户信息
func AuthInfo(c *gin.Context) {
	u := c.MustGet("admin").(admin.Admin)
	utils.SuccessData(c, u)
}