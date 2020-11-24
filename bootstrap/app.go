package bootstrap

import (
	"github.com/gin-gonic/gin" // 基于 gin 框架
	"github.com/wuyan94zl/api/routes"
)

func Start() *gin.Engine{
	// 数据库初始化
	autoMigrate()
	router := gin.Default() // 获取路由实例
	routes.ApiRouter(router) // 注册路由
	return router // 返回路由
}