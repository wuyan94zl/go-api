package main

import (
	"github.com/wuyan94zl/api/bootstrap"
	"github.com/wuyan94zl/api/pkg/generate"
	"os"
)

func main() {
	lenNum := len(os.Args)
	if lenNum < 2 {
		panic("参数错误")
	}
	model := os.Args[1]

	caseVal := ""
	if lenNum > 2 {
		caseVal = os.Args[2]
	}

	uri := ""
	if lenNum > 3 {
		uri = os.Args[3]
	}

	pkgUri := ""
	if lenNum > 4 {
		pkgUri = os.Args[4]
	}

	switch caseVal {
	case "route":
		generate.SetRoute(bootstrap.MigrateStruct[model],uri,pkgUri)
	default:
		generate.SetCurd(bootstrap.MigrateStruct[model],uri)
	}
}
