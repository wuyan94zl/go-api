package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/rbac/model"
	"github.com/wuyan94zl/api/pkg/utils"
	"strconv"
	"strings"
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
	var Role model.Role
	Role.Name = c.PostForm("name")
	Role.Description = c.PostForm("description")

	orm.GetInstance().Create(&Role)
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
	var Role model.Role
	orm.GetInstance().First(&Role, id)
	if Role.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	Role.Name = c.PostForm("name")
	Role.Description = c.PostForm("description")
	orm.GetInstance().Save(Role)
	utils.SuccessData(c, Role) // 返回创建成功的信息
}
func RoleDelete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Role model.Role

	orm.GetInstance().First(&Role, id)
	if Role.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	orm.GetInstance().Delete(Role)
	where := make(map[string]interface{})
	where["role_id"] = Role.Id
	orm.GetInstance().Where(where).Delete(model.RoleMenu{})
	orm.GetInstance().Where(where).Delete(model.RolePermission{})
	utils.SuccessData(c, "删除成功")
}
func RoleInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Role model.Role
	orm.GetInstance().First(&Role, id, "Menus", "Permissions")
	utils.SuccessData(c, Role)
}

func RolePaginate(c *gin.Context) {
	var Role []model.Role
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	paginate := orm.SetPageList(&Role, int64(page))
	orm.GetInstance().Order("id desc").Paginate(paginate)
	utils.SuccessData(c, paginate)
}

func RoleSelectAll(c *gin.Context) {
	var Role []model.Role
	orm.GetInstance().Get(&Role)
	result := make(map[int]map[string]string)
	for k, v := range Role {
		result[k] = map[string]string{"id": strconv.Itoa(int(v.Id)), "name": v.Name}
	}
	utils.SuccessData(c, result)
}

// 获取
func GetPermissionMenu(c *gin.Context) {
	roleId := c.Query("id")
	role := model.Role{}
	orm.GetInstance().First(&role, roleId, "Permissions")
	tree, has := role.GetPermissionMenu()
	result := make(map[string]interface{})
	result["tree"] = tree
	result["has"] = has
	utils.SuccessData(c, result)
}

// 设置
func RolePermissionMenu(c *gin.Context) {
	permissionStr := c.PostForm("permission_id")
	permissionId := strings.Split(permissionStr, ",")
	roleId := c.Query("id")
	role := model.Role{}
	orm.GetInstance().First(&role, roleId)
	role.DelPermissionMenu()
	role.SetPermissionMenu(permissionId)
	utils.SuccessData(c, "设置成功")
}
