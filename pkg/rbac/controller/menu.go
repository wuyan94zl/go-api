package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/rbac/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func MenuCreate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["parent_id"] = []string{"numeric"}
	data["name"] = []string{"required"}
	data["route"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	var Menu model.Menu
	ParentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	Menu.ParentId = uint64(ParentId)
	Menu.Name = c.PostForm("name")
	Menu.Route = c.PostForm("route")
	Menu.Icon = c.DefaultPostForm("icon","")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	Menu.Sort = uint64(sort)
	orm.GetInstance().Create(&Menu)
	utils.SuccessData(c, Menu) // 返回创建成功的信息
}
func MenuUpdate(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["parent_id"] = []string{"numeric"}
	data["name"] = []string{"required"}
	data["route"] = []string{"required"}

	validate := utils.Validator(c.Request, data)
	if validate != nil {
		utils.SuccessErr(c, 403, validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu model.Menu
	orm.GetInstance().First(&Menu, id)
	if Menu.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	ParentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	Menu.ParentId = uint64(ParentId)
	Menu.Name = c.PostForm("name")
	Menu.Route = c.PostForm("route")
	Menu.Icon = c.DefaultPostForm("icon","")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	Menu.Sort = uint64(sort)
	orm.GetInstance().Save(Menu)
	utils.SuccessData(c, Menu) // 返回创建成功的信息
}
func MenuDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu model.Menu

	orm.GetInstance().First(&Menu, id, "Permissions")
	if Menu.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	orm.GetInstance().Delete(&Menu)
	orm.GetInstance().Delete(Menu.Permissions)
	utils.SuccessData(c, "删除成功")
}
func MenuList(c *gin.Context) {
	var Menu []model.Menu
	orm.GetInstance().Order("sort").Get(&Menu, "Permissions")
	tree := model.RecursionMenuList(Menu, 0, 1)
	utils.SuccessData(c, tree)
}
