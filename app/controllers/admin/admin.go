package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/admin"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/rbac/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"strings"
)

// 登录
func Login(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)
	data["email"] = []string{"required", "min:6", "email"}
	data["password"] = []string{"required", "between:6,20"}
	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	// 查询获取到用户
	email := c.PostForm("email")
	password := c.PostForm("password")
	where := make(map[string]interface{})
	where["email"] = email
	info := admin.Admin{}
	orm.GetInstance().Where(where).One(&info)
	if info.Id == 0 {
		utils.SuccessErr(c, -1000, "用户名或密码错误")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password)) != nil {
		utils.SuccessErr(c, -1000, "用户名或密码错误")
		return
	}
	// 换取token
	token, err := info.Token()
	if err != nil {
		utils.SuccessErr(c, -1000, "未知错误")
	} else {
		utils.SuccessData(c, token)
	}
}

// 获取登录用户信息
func AuthInfo(c *gin.Context) {
	u := c.MustGet("auth").(admin.Admin)
	utils.SuccessData(c, u)
}

// 设置角色
func SetRole(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("id"))
	roleIds := c.PostForm("role_id")
	ids := strings.Split(roleIds, ",")
	where := make(map[string]interface{})
	where["user_id"] = userId
	orm.GetInstance().Where(where).Delete(model.UserRole{})
	var userRoles []model.UserRole
	for _, id := range ids {
		if id != ""{
			uid, _ := strconv.Atoi(id)
			userRoles = append(userRoles, model.UserRole{UserId: uint64(userId), RoleId: uint64(uid)})
		}
	}
	orm.GetInstance().Create(userRoles)
	utils.SuccessData(c, "ok")
}

// 用户菜单
func Menus(c *gin.Context) {
	id := c.MustGet("auth_id")
	adminInfo := admin.Admin{}
	orm.GetInstance().First(&adminInfo, id, "Roles")
	tree := adminInfo.Menus()
	utils.SuccessData(c, tree)
}
