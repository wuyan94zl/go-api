package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/app/middleware"
)

var Route = gin.Default()
var RouteGroup *gin.RouterGroup
var AuthRouteGroup *gin.RouterGroup

type Item struct {
	Method string
	Route  string
	Action gin.HandlerFunc
}

func init() {
	Route.Use(middleware.Cors())
	RouteGroup = Route.Group("api")
	AuthRouteGroup = Group("api",middleware.ApiAuth())
}

func Group(relativePath string, middleware ...gin.HandlerFunc) *gin.RouterGroup {
	routerGroup := Route.Group(relativePath)
	routerGroup.Use(middleware...)
	return routerGroup
}

func Register(routes []Item, group ...*gin.RouterGroup) {
	var route *gin.RouterGroup
	if len(group) > 0 {
		route = group[0]
	} else {
		route = RouteGroup
	}
	for _, v := range routes {
		switch v.Method {
		case "get":
			route.GET(v.Route, v.Action)
		case "post":
			route.POST(v.Route, v.Action)
		case "put":
			route.PUT(v.Route, v.Action)
		case "delete":
			route.DELETE(v.Route, v.Action)
		}
	}
}
