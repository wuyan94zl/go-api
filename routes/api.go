package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers"
	"github.com/wuyan94zl/api/app/middleware"
)
// 注册路由列表
func ApiRouter(router *gin.Engine)  {
	api := router.Group("/api")
	api.POST("/user/create", controllers.UserCreate) // 增
	api.GET("/user/delete/:id", controllers.UserDelete) // 删
	api.POST("/user/update/:id", controllers.UserUpdate) // 改
	api.GET("/users", controllers.UserList) //查

	api.POST("/user/login",controllers.UserLogin) // 登录

	auth := router.Group("api") // 认证路由组
	auth.Use(middleware.ApiAuth()) // 登录认证中间件
	auth.GET("user/info",controllers.UserInfo) // 登录用户信息

}

