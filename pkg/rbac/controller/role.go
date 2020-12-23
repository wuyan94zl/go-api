package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/rbac"
	"github.com/wuyan94zl/api/pkg/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func RoleCreate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["name"] = []string{"required"}
	data["description"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	var Role rbac.Role
	Role.Name = c.PostForm("name")
	Role.Description = c.PostForm("description")

	model.Create(&Role)
	utils.SuccessData(c, Role) // 返回创建成功的信息
}
func RoleUpdate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["name"] = []string{"required"}
	data["description"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Role rbac.Role
	model.First(&Role, id)
	if Role.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	Role.Name = c.PostForm("name")
	Role.Description = c.PostForm("description")

	model.UpdateOne(Role)
	utils.SuccessData(c, Role) // 返回创建成功的信息
}
func RoleDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Role rbac.Role

	model.First(&Role, id)
	if Role.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	model.DeleteOne(&Role)
	model.DeleteOne(Role.Menus)
	model.DeleteOne(Role.Permissions)
	utils.SuccessData(c, "删除成功")
}
func RoleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Role rbac.Role
	model.First(&Role, id, "Menus", "Permissions")

	utils.SuccessData(c, Role)
}

func RolePermission(c *gin.Context){

}

func RolePaginate(c *gin.Context) {
	var conditions []model.Condition

	var Role []rbac.Role
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	lists := model.Paginate(&Role, model.PageInfo{Page: int64(page), PageSize: int64(pageSize)}, conditions)
	utils.SuccessData(c, lists)
}
