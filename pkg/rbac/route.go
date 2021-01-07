package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/rbac/controller"
	"github.com/wuyan94zl/api/pkg/utils"
)

// 注册路由列表
func RegisterRouter(router *gin.RouterGroup) {
	utils.AddRoute(router, "POST", "/role/create", controller.RoleCreate)
	utils.AddRoute(router, "POST", "/role/update", controller.RoleUpdate)
	utils.AddRoute(router, "GET", "/role/delete", controller.RoleDelete)
	utils.AddRoute(router, "GET", "/role/info", controller.RoleInfo)
	utils.AddRoute(router, "GET", "/role/paginate", controller.RolePaginate)
	utils.AddRoute(router, "GET", "/role/select", controller.RoleSelectAll)
	utils.AddRoute(router, "POST", "/role/menu/permission", controller.RolePermissionMenu)
	utils.AddRoute(router, "GET", "/role/menu/permission", controller.GetPermissionMenu)

	utils.AddRoute(router, "POST", "/permission/create", controller.PermissionCreate)
	utils.AddRoute(router, "POST", "/permission/update", controller.PermissionUpdate)
	utils.AddRoute(router, "GET", "/permission/delete", controller.PermissionDelete)
	utils.AddRoute(router, "GET", "/permission/lists", controller.PermissionList)

	utils.AddRoute(router, "POST", "/menu/create", controller.MenuCreate)
	utils.AddRoute(router, "POST", "/menu/update", controller.MenuUpdate)
	utils.AddRoute(router, "GET", "/menu/delete", controller.MenuDelete)
	utils.AddRoute(router, "GET", "/menu/lists", controller.MenuList)
}
