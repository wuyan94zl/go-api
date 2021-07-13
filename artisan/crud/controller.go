package crud

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var controllerTpl = `package {{.packageName}}

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/response"
	"github.com/wuyan94zl/go-api/pkg/validate"
)

func Lists(c *gin.Context) {
	model := GetModel()
	response.Success(model.Lists(c))
}

func Create(c *gin.Context) {
	model := GetModel()
	if ok, msg := validate.StructValidate(c.Request, model); !ok {
		response.Error(401, msg)
	}
	model.Create(c)
	response.Success("创建成功")
}

func Update(c *gin.Context) {
	model := GetModel()
	if ok, msg := validate.StructValidate(c.Request, model); !ok {
		response.Error(401, msg)
	}
	model.Info(c)
	model.Update(c)
	response.Success("更新成功")
}

func Info(c *gin.Context) {
	model := GetModel()
	model.Info(c)
	response.Success(model)
}

func Delete(c *gin.Context) {
	model := GetModel()
	model.Delete(c)
	response.Success("删除成功")
}

`

func setController(jsonData *jsonStruct) error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          getDir(jsonData.PackageName),
		Filename:     "controller.go",
		TemplateFile: controllerTpl,
		Data: map[string]string{
			"packageName": jsonData.PackageName,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
