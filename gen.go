package main

import (
	"fmt"
	"github.com/wuyan94zl/api/bootstrap"
	"github.com/wuyan94zl/api/pkg/generate"
	"os"
)

func main(){

	params := os.Args[1]
	model := os.Args[2]

	if params == "" || model == ""{
		fmt.Println("参数错误")
	}

	switch params {
	case "gene":
		generate.StructCurd(bootstrap.MigrateStruct[model])
	case "route":
		generate.StructRoute(bootstrap.MigrateStruct[model])
	}
}
