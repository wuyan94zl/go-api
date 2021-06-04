package model

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var tpl = `{
  "package_name": "{{.package}}",
  "struct_name": "{{.struct}}",
  "fields": [
    {
      "field": "Id",
      "type_name": "uint64",
      "tags": {
        "json": "id"
      }
    },
    {
      "field": "CreatedAt",
      "type_name": "time.Time",
      "tags": {
        "json": "created_at"
      }
    },
    {
      "field": "UpdatedAt",
      "type_name": "time.Time",
      "tags": {
        "json": "updated_at"
      }
    }
  ]
}
`

type Command struct {
	Name string
}

func (c *Command) GetDir() string {
	return utils.GetDir("http", c.Name)
}

func (c *Command) Run() error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          c.GetDir(),
		Filename:     "model.json",
		TemplateFile: tpl,
		Data: map[string]string{
			"package": c.Name,
			"struct":  c.Name,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
