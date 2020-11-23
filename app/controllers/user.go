package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models"
	"github.com/wuyan94zl/api/app/models/user"
	"github.com/wuyan94zl/api/pkg/auth"
	"github.com/wuyan94zl/api/pkg/database"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func UserLogin(c *gin.Context)  {
	user := user.User{}
	// 查询获取到用户
	email := c.PostForm("email")
	password := c.PostForm("password")
	database.DB.Where("email = ?",email).First(&user)

	if user.Password != password {
		utils.SuccessErr(c,-1000,"用户名或密码错误")
		return
	}

	// 换取token
	token, err := auth.GetToken(&user)
	if err != nil{
		utils.SuccessErr(c,-1000,"未知错误")
	}else{
		utils.SuccessData(c,token)
	}
}

func UserInfo(c *gin.Context)  {
	u := c.MustGet("Authuser").(user.User)
	utils.SuccessData(c,u)
}

func UserCreate(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	name := c.PostForm("name")
	user := user.User{Email: email,Name: name,Password: password}
	// database.DB 为数据库的连接实例
	database.DB.Create(&user)
	utils.SuccessData(c,user) // 返回创建成功的信息
}
func UserDelete(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SuccessErr(c,-1000,"用户ID参数错误")
		return
	}
	user := user.User{}
	database.DB.First(&user,id)
	if user.Id != id {
		// 用户不存在，禁止删除
		utils.SuccessErr(c,-1000,"用户不存在")
		return
	}
	database.DB.Delete(user)
	utils.SuccessData(c,"删除成功")
}
func UserUpdate(c *gin.Context) {
	id,err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.SuccessErr(c,-1000,"用户ID参数错误")
		return
	}
	email := c.PostForm("email")
	name := c.PostForm("name")

	user := user.User{}
	database.DB.First(&user, id)

	if user.Id != id {
		// 用户不存在，禁止修改
		utils.SuccessErr(c,-1000,"用户不存在")
		return
	}
	user.Email = email
	user.Name = name
	database.DB.Save(&user)
	utils.SuccessData(c,user) // 返回更新成功的信息
}
func UserList(c *gin.Context) {
	id := c.Query("id")

	query := ""
	if id != "" {
		query = "id > ?"
	}

	var user []user.User
	list := models.Paginate(&user,models.PageInfo{Page: 1, PageSize: 10},query,id)
	utils.SuccessData(c,list) // 返回查询列表
}