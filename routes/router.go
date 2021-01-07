package routes

import (
	"github.com/gin-gonic/gin" // 基于 gin 框架
	"github.com/wuyan94zl/api/app/middleware"
	"github.com/wuyan94zl/api/pkg/rbac"
)

// 注册当前
func Register() *gin.Engine {
	router := gin.Default() // 获取路由实例
	router.Use(middleware.Cors())
	// 定义默认普通api组
	api := router.Group("/api")
	ApiRouter(api)

	// 定义默认auth认证api组
	authApi := router.Group("api")
	authApi.Use(middleware.ApiAuth())
	AuthRouter(authApi)

	// 定义默认auth认证api组
	permissionApi := router.Group("api")
	permissionApi.Use(middleware.ApiAuth())
	permissionApi.Use(rbac.PermissionCheck())
	rbac.RegisterRouter(permissionApi)
	PermissionRouter(permissionApi)

	return router // 返回路由
}
