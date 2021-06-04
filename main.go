package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/wuyan94zl/go-api/bootstrap"
)

func main() {
	app := bootstrap.Start()
	addr := fmt.Sprintf(":%s", viper.GetString("app.port"))
	go func() {
		err := app.Run(addr)
		if err != nil {
			return
		}
	}()
	bootstrap.Timer()
}
