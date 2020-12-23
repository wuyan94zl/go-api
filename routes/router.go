package routes
import (
	"github.com/gin-gonic/gin" // 基于 gin 框架
	"github.com/wuyan94zl/api/app/middleware"
	"github.com/wuyan94zl/api/pkg/rbac"
)

type RouteList struct {
	Method string
	Route  string
}

var Post = "POST"
var Get = "GET"

var Routes map[int]RouteList

// 添加路由
func AddRoute(){

}

// 注册当前
func Register() *gin.Engine {
	router := gin.Default() // 获取路由实例
	// 定义默认普通api组
	api := router.Group("/api")

	// 定义默认auth认证api组
	authApi := router.Group("api")
	authApi.Use(middleware.ApiAuth())

	// 定义默认auth认证api组
	permissionApi := authApi
	permissionApi.Use(rbac.PermissionCheck())

	ApiRouter(api)
	AuthRouter(authApi)
	rbac.RbacRouter(permissionApi)

	return router // 返回路由
}