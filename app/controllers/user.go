package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/user"
	"github.com/wuyan94zl/api/pkg/auth"
	"github.com/wuyan94zl/api/pkg/database"
	"github.com/wuyan94zl/api/pkg/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)

// 登录
func UserLogin(c *gin.Context) {
	// 查询获取到用户
	email := c.PostForm("email")
	password := c.PostForm("password")
	var condition []model.Condition
	condition = model.SetCondition(condition,"email",email)
	u := model.GetOne(&user.User{}, condition)
	info := u.(*user.User)
	if info.Id == 0 {
		utils.SuccessErr(c, -1000, "用户名或密码错误")
		return
	}
	if bcrypt.CompareHashAndPassword([]byte(info.Password), []byte(password)) != nil {
		utils.SuccessErr(c, -1000, "用户名或密码错误")
		return
	}
	// 换取token
	token, err := auth.GetToken(info)
	if err != nil {
		utils.SuccessErr(c, -1000, "未知错误")
	} else {
		utils.SuccessData(c, token)
	}
}

// 获取登录用户信息
func UserInfo(c *gin.Context) {
	u := c.MustGet("user").(user.User)
	utils.SuccessData(c, u)
}

// 创建用户
func UserCreate(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")
	pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := user.User{Email: email, Name: name, Password: string(pwd)}
	// database.DB 为数据库的连接实例
	database.DB.Create(&user)
	utils.SuccessData(c, user) // 返回创建成功的信息
}

// 删除用户
func UserDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	u, _ := model.GetFirst(&user.User{}, id)
	info := u.(*user.User)
	if info.Id == 0 {
		utils.SuccessErr(c, -1000, "用户不存在")
		return
	}
	model.DeleteOne(&info)
	utils.SuccessData(c, "删除成功")
}

// 更新用户
func UserUpdate(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	u, _ := model.GetFirst(&user.User{}, id)
	info := u.(*user.User)
	if info.Id == 0 {
		utils.SuccessErr(c, -1000, "用户不存在")
		return
	}
	email := c.PostForm("email")
	name := c.PostForm("name")
	password := c.DefaultPostForm("password", "")
	info.Email = email
	info.Name = name
	if password != "" {
		pwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		info.Password = string(pwd)
	}
	rlt := model.UpdateOne(&info)
	utils.SuccessData(c, rlt) // 返回更新成功的信息
}

// 获取单个用户信息
func UserOne(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	u, _ := model.GetFirst(&user.User{}, id)
	info := u.(*user.User)
	if info.Id == 0 {
		utils.SuccessErr(c, -1000, "用户不存在")
		return
	}
	utils.SuccessData(c, info)
}

// 获取多个用户信息
func UserList(c *gin.Context) {
	var conditions []model.Condition
	email := c.PostForm("email")
	if email != ""{
		conditions = model.SetCondition(conditions,"email",fmt.Sprintf("%s%s", email, "%"),"like")
	}
	name := c.PostForm("name")
	if name != ""{
		conditions = model.SetCondition(conditions,"name",name)
	}
	lists := model.GetAll(&[]user.User{}, conditions)
	utils.SuccessData(c, lists)
}

// 获取分页用户信息
func UserPaginate(c *gin.Context) {
	var conditions []model.Condition
	email := c.PostForm("email")
	if email != ""{
		conditions = model.SetCondition(conditions,"email",fmt.Sprintf("%s%s", email, "%"),"like")
	}
	name := c.PostForm("name")
	if name != ""{
		conditions = model.SetCondition(conditions,"name",name)
	}

	var u []user.User
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	list := model.Paginate(&u, model.PageInfo{Page: page, PageSize: pageSize}, conditions)
	utils.SuccessData(c, list) // 返回查询列表
}
