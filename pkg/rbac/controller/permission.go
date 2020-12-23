package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/rbac"
	"github.com/wuyan94zl/api/pkg/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func PermissionCreate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["name"] = []string{"required"}
	data["route"] = []string{"required"}
	data["menu_id"] = []string{"required"}
	data["description"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	var Permission rbac.Permission
	Permission.Name = c.PostForm("name")
	Permission.Route = c.PostForm("route")
	MenuId, _ := strconv.Atoi(c.PostForm("menu_id"))
	Permission.MenuId = uint64(MenuId)
	Permission.Description = c.PostForm("description")
	model.Create(&Permission)
	utils.SuccessData(c, Permission) // 返回创建成功的信息
}
func PermissionUpdate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["name"] = []string{"required"}
	data["route"] = []string{"required"}
	data["menu_id"] = []string{"required"}
	data["description"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Permission rbac.Permission
	model.First(&Permission, id)
	if Permission.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	Permission.Name = c.PostForm("name")
	Permission.Route = c.PostForm("route")
	MenuId, _ := strconv.Atoi(c.PostForm("menu_id"))
	Permission.MenuId = uint64(MenuId)
	Permission.Description = c.PostForm("description")
	model.UpdateOne(Permission)
	utils.SuccessData(c, Permission) // 返回创建成功的信息
}
func PermissionDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Permission rbac.Permission

	model.First(&Permission, id)
	if Permission.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	model.DeleteOne(&Permission)
	utils.SuccessData(c, "删除成功")
}
func PermissionInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Permission rbac.Permission
	model.First(&Permission, id)

	utils.SuccessData(c, Permission)
}
func PermissionPaginate(c *gin.Context) {
	var conditions []model.Condition

	var Permission []rbac.Permission
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	lists := model.Paginate(&Permission, model.PageInfo{Page: int64(page), PageSize: int64(pageSize)}, conditions)
	utils.SuccessData(c, lists)
}
