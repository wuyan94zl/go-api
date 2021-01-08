package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
	"github.com/wuyan94zl/api/pkg/utils"
)

// 注册路由列表
func PermissionRouter(router *gin.RouterGroup) {
	utils.AddRoute(router, "POST", "/admin/role", admin.SetRole)
}
