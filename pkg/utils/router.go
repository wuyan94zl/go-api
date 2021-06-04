package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteList struct {
	Method string
	Route  string
}

// 路由数
var num = 0

// 所有路由
var AllRoutes = make(map[int]RouteList)

// 添加路由
func AddRoute(router *gin.RouterGroup, method string, path string, handle gin.HandlerFunc) {
	prefix := router.BasePath()
	AllRoutes[num] = RouteList{Method: method, Route: fmt.Sprintf("%s%s",prefix,path)}
	num++
	switch method {
	case "GET":
		router.GET(path, handle)
	case "POST":
		router.POST(path, handle)
	case "Any":
		router.Any(path, handle)
	case "PATCH":
		router.PATCH(path, handle)
	case "DELETE":
		router.DELETE(path, handle)
	case "PUT":
		router.PUT(path, handle)
	}
}
