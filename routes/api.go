package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
)

// 注册路由列表
func ApiRouter(router *gin.RouterGroup) {
	router.POST("/admin/login", admin.Login) // 登录


	// start admin
	router.POST("/admin/create",admin.Create)
	router.POST("/admin/update",admin.Update)
	router.GET("/admin/delete",admin.Delete)
	router.GET("/admin/info",admin.Info)
	router.POST("/admin/paginate",admin.Paginate)
	// end admin
}