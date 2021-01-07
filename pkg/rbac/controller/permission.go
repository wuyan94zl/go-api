package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/rbac/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func PermissionCreate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["route"] = []string{"required"}
	data["menu_id"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	var Permission model.Permission
	Permission.Name = c.DefaultPostForm("name","")
	Permission.Route = c.PostForm("route")
	MenuId, _ := strconv.Atoi(c.PostForm("menu_id"))
	Permission.MenuId = uint64(MenuId)
	Permission.Description = c.DefaultPostForm("description","")
	orm.GetInstance().Create(&Permission)
	utils.SuccessData(c, Permission) // 返回创建成功的信息
}
func PermissionUpdate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["route"] = []string{"required"}
	data["menu_id"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Permission model.Permission
	orm.GetInstance().First(&Permission, id)
	if Permission.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	Permission.Name = c.DefaultPostForm("name","")
	Permission.Route = c.PostForm("route")
	MenuId, _ := strconv.Atoi(c.PostForm("menu_id"))
	Permission.MenuId = uint64(MenuId)
	Permission.Description = c.DefaultPostForm("description","")
	orm.GetInstance().Save(Permission)
	utils.SuccessData(c, Permission) // 返回创建成功的信息
}
func PermissionDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Permission model.Permission

	orm.GetInstance().First(&Permission, id)
	if Permission.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	orm.GetInstance().Delete(&Permission)
	utils.SuccessData(c, "删除成功")
}

func PermissionList(c *gin.Context) {
	allRoutes := utils.AllRoutes
	utils.SuccessData(c, allRoutes)
}
