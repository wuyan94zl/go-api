package command

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var tpl = `package {{.package}}

import (
	"fmt"
	"time"
)

type Job struct{}

func (j Job) Run() {
	fmt.Println("Execution per minute", time.Now().Format("2006-01-02 15:4:05"))
}

`

type Command struct {
	Name string
}

func (c *Command) GetDir() string {
	return utils.GetDir("command", c.Name)
}

func (c *Command) Run() error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          c.GetDir(),
		Filename:     "handle.go",
		TemplateFile: tpl,
		Data: map[string]string{
			"package": c.Name,
		},
	})
	if err != nil {
		fmt.Println("err：", err)
		return err
	}
	return nil
}
