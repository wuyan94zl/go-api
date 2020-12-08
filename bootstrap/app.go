package bootstrap

import (
	"github.com/gin-gonic/gin" // 基于 gin 框架
	"github.com/wuyan94zl/api/routes"
)

func Start() *gin.Engine{
	// 数据库初始化
	autoMigrate()
	// 路由注册
	router := routes.Register()
	return router // 返回路由
}