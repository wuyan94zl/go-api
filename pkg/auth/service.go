package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/routes"
	"github.com/wuyan94zl/mysql"
)

func Init(route ...*gin.RouterGroup) {
	// 表结构迁移
	MigrateStruct := make(map[string]interface{})
	MigrateStruct["auth"] = User{}
	mysql.AutoMigrate(MigrateStruct)

	//路由注册
	routeItems := make([]routes.Item, 1)
	routeItems = append(routeItems, routes.Item{Method: "post", Route: "auth/register", Action: Register})
	routeItems = append(routeItems, routes.Item{Method: "post", Route: "auth/login", Action: Login})
	routes.Register(routeItems, route...)
	//routes.AuthRouteGroup = routes.Group("api",middleware.ApiAuth())
	authRouteItems := make([]routes.Item, 1)
	authRouteItems = append(authRouteItems, routes.Item{Method: "post", Route: "auth/logout", Action: Logout})
	authRouteItems = append(authRouteItems, routes.Item{Method: "get", Route: "auth/info", Action: Info})
	routes.Register(authRouteItems, routes.AuthRouteGroup)
}

func GetModel() User {
	model := User{}
	return model
}

func GetAuthInfo(c *gin.Context) User {
	info, err := c.Get("auth_info")
	if !err {
		user := User{}
		mysql.GetInstance().First(&user, c.MustGet("auth_id"))
		c.Set("auth_info",user)
		return user
	} else {
		return info.(User)
	}
}
