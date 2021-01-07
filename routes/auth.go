package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
	"github.com/wuyan94zl/api/pkg/utils"
)

// 注册路由列表
func AuthRouter(router *gin.RouterGroup) {
	utils.AddRoute(router, "GET", "/admin/auth", admin.AuthInfo) // 登录用户信息
	utils.AddRoute(router, "GET", "/admin/menus", admin.Menus)
}
