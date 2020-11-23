package main

import (
	"fmt"
	"github.com/wuyan94zl/api/config"
	conf "github.com/wuyan94zl/api/pkg/config"
	"github.com/wuyan94zl/api/routes"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}
func main() {
	router := routes.Register()
	addr := fmt.Sprintf(":%s", conf.GetString("app.port"))
	router.Run(addr)
}
