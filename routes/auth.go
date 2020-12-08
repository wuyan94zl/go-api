package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
)

// 注册路由列表
func AuthRouter(router *gin.RouterGroup) {
	router.GET("/admin/auth", admin.AuthInfo) // 登录用户信息
}
