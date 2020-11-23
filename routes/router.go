package routes
import (
	"github.com/gin-gonic/gin" // 基于 gin 框架
)
// 注册当前
func Register() *gin.Engine {
	router := gin.Default() // 获取路由实例
	ApiRouter(router) // 注册路由
	return router // 返回路由
}