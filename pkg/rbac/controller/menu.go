package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/model"
	rbac "github.com/wuyan94zl/api/pkg/rbac/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func MenuCreate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["parent_id"] = []string{"required", "numeric"}
	data["name"] = []string{"required"}
	data["route"] = []string{"required"}
	data["description"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	var Menu rbac.Menu
	ParentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	Menu.ParentId = uint64(ParentId)
	Menu.Name = c.PostForm("name")
	Menu.Route = c.PostForm("route")
	Menu.Description = c.PostForm("description")

	model.Create(&Menu)
	utils.SuccessData(c, Menu) // 返回创建成功的信息
}
func MenuUpdate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["parent_id"] = []string{"required", "numeric"}
	data["name"] = []string{"required"}
	data["route"] = []string{"required"}
	data["description"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu rbac.Menu
	model.First(&Menu, id)
	if Menu.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	ParentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	Menu.ParentId = uint64(ParentId)
	Menu.Name = c.PostForm("name")
	Menu.Route = c.PostForm("route")
	Menu.Description = c.PostForm("description")

	model.UpdateOne(Menu)
	utils.SuccessData(c, Menu) // 返回创建成功的信息
}
func MenuDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu rbac.Menu

	model.First(&Menu, id, "Permissions")
	if Menu.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	model.DeleteOne(&Menu)
	model.DeleteOne(Menu.Permissions)
	utils.SuccessData(c, "删除成功")
}
func MenuPaginate(c *gin.Context) {
	var conditions []model.Condition
	var Menu []rbac.Menu
	model.GetAll(&Menu, conditions, "Permissions")
	tree := rbac.RecursionMenuList(Menu, 0, 1)
	utils.SuccessData(c, tree)
}
