package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
	"github.com/wuyan94zl/api/pkg/utils"
)

// 注册路由列表
func ApiRouter(router *gin.RouterGroup) {
	utils.AddRoute(router, "POST", "/admin/login", admin.Login)
}