package crud

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var serviceTpl = `package {{.packageName}}

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/mysql"
	"strconv"
	"github.com/wuyan94zl/go-api/routes"
)

func Init(route ...*gin.RouterGroup) {
	// 表结构迁移
	MigrateStruct := make(map[string]interface{})
	MigrateStruct["{{.packageName}}"] = {{.modelName}}{}
	mysql.AutoMigrate(MigrateStruct)

	//路由注册
	routeItems := make([]routes.Item, 1)
	routeItems = append(routeItems, routes.Item{Method: "post", Route: "{{.packageName}}/create", Action: Create})
	routeItems = append(routeItems, routes.Item{Method: "get", Route: "{{.packageName}}/lists", Action: Lists})
	routeItems = append(routeItems, routes.Item{Method: "get", Route: "{{.packageName}}/info", Action: Info})
	routeItems = append(routeItems, routes.Item{Method: "put", Route: "{{.packageName}}/update", Action: Update})
	routeItems = append(routeItems, routes.Item{Method: "delete", Route: "{{.packageName}}/delete", Action: Delete})
	routes.Register(routeItems, route...)
}

var ins *Service

type Service struct{}

// GetService 单例
func GetService() *Service {
	if ins == nil {
		ins = &Service{}
	}
	return ins
}

func (s *Service) PaginateData(c *gin.Context) *mysql.PageList {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	data := &[]{{.modelName}}{}
	return mysql.SetPageList(data, int64(page), int64(pageSize))
}

func GetModel() {{.modelName}} {
	model := {{.modelName}}{}
	return model
}

`

func setService(jsonData *jsonStruct) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          getDir(jsonData.PackageName),
		Filename:     "service.go",
		TemplateFile: serviceTpl,
		Data: map[string]string{
			"packageName": jsonData.PackageName,
			"modelName":   jsonData.StructName,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
