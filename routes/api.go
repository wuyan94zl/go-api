package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
	"github.com/wuyan94zl/api/app/middleware"
)

// 注册路由列表
func ApiRouter(router *gin.Engine) {
	api := router.Group("/api")
	api.POST("/admin/login", admin.Login) // 登录

	// 登录鉴权路由
	auth := router.Group("api")      // 认证路由组
	auth.Use(middleware.ApiAuth())   // 登录认证中间件
	auth.GET("/admin/auth", admin.AuthInfo) // 登录用户信息
}




