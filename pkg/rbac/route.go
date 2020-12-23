package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/rbac/controller"
)

// 注册路由列表
func RbacRouter(router *gin.RouterGroup) {
	router.POST("/role/create", controller.RoleCreate)
	router.POST("/role/update", controller.RoleUpdate)
	router.GET("/role/delete", controller.RoleDelete)
	router.GET("/role/info", controller.RoleInfo)
	router.POST("/role/paginate", controller.RolePaginate)

	router.POST("/permission/create", controller.PermissionCreate)
	router.POST("/permission/update", controller.PermissionUpdate)
	router.GET("/permission/delete", controller.PermissionDelete)
	router.POST("/permission/paginate", controller.PermissionPaginate)

	router.POST("/menu/create", controller.MenuCreate)
	router.POST("/menu/update", controller.MenuUpdate)
	router.GET("/menu/delete", controller.MenuDelete)
	router.POST("/menu/paginate", controller.MenuPaginate)
}