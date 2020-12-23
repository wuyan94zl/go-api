package menu
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/rbac/model"
	"github.com/wuyan94zl/api/pkg/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
)

func Create(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["parent_id"] = []string{"required","numeric"} 
	data["name"] = []string{"required"} 
	data["route"] = []string{"required"} 
	data["description"] = []string{"required"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	var Menu model.Menu
	ParentId,_ := strconv.Atoi(c.PostForm("parent_id"))
	Menu.ParentId = uint64(ParentId)
	Menu.Name = c.PostForm("name")
	Menu.Route = c.PostForm("route")
	Menu.Description = c.PostForm("description")


	model.Create(&Menu)
	utils.SuccessData(c, Menu) // 返回创建成功的信息
}
func Update(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["parent_id"] = []string{"required","numeric"} 
	data["name"] = []string{"required"} 
	data["route"] = []string{"required"} 
	data["description"] = []string{"required"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu model.Menu
	model.First(&Menu,id,"Permissions")
	if Menu.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	ParentId,_ := strconv.Atoi(c.PostForm("parent_id"))
	Menu.ParentId = uint64(ParentId)
	Menu.Name = c.PostForm("name")
	Menu.Route = c.PostForm("route")
	Menu.Description = c.PostForm("description")


	model.UpdateOne(Menu)
	utils.SuccessData(c, Menu) // 返回创建成功的信息
}
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu model.Menu

	model.First(&Menu,id)
	if Menu.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	model.DeleteOne(&Menu)
	utils.SuccessData(c, "删除成功")
}
func Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Menu model.Menu
	model.First(&Menu,id,"Permissions")

	utils.SuccessData(c, Menu)
}
func Paginate(c *gin.Context) {
	var conditions []model.Condition

	var Menu []model.Menu
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	lists := model.Paginate(&Menu, model.PageInfo{Page: int64(page), PageSize: int64(pageSize)}, conditions,"Permissions")
	utils.SuccessData(c, lists)
}