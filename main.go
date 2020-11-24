package main

import (
	"fmt"
	"github.com/wuyan94zl/api/bootstrap"
	"github.com/wuyan94zl/api/config"
	conf "github.com/wuyan94zl/api/pkg/config"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}
func main() {
	app := bootstrap.Start()
	addr := fmt.Sprintf(":%s", conf.GetString("app.port"))
	app.Run(addr)
}
