package crud

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var controllerTpl = `package {{.packageName}}

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/utils"
	"github.com/wuyan94zl/go-api/pkg/validate"
)

func Lists(c *gin.Context) {
	model := GetModel()
	utils.SuccessData(c, model.Lists(c))
}

func Create(c *gin.Context) {
	model := GetModel()
	if ok, msg := validate.StructValidate(c.Request, model); !ok {
		utils.SuccessErr(c, 401, msg)
		return
	}
	if create, err := model.Create(c); !create {
		utils.SuccessErr(c, 500, err)
		return
	}
	utils.SuccessData(c, "创建成功")
}

func Update(c *gin.Context) {
	model := GetModel()
	if ok, msg := validate.StructValidate(c.Request, model); !ok {
		utils.SuccessErr(c, 401, msg)
		return
	}
	if ok, err := model.Info(c); !ok {
		utils.SuccessErr(c, 500, err)
		return
	}
	if ok, err := model.Update(c); !ok {
		utils.SuccessErr(c, 500, err)
		return
	}
	utils.SuccessData(c, "更新成功")
}

func Info(c *gin.Context) {
	model := GetModel()
	if ok, err := model.Info(c); !ok {
		utils.SuccessErr(c, 500, err)
		return
	}
	utils.SuccessData(c, model)
}

func Delete(c *gin.Context) {
	model := GetModel()
	if ok, err := model.Delete(c); !ok {
		utils.SuccessErr(c, 500, err)
		return
	}
	utils.SuccessData(c, "删除成功")
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
