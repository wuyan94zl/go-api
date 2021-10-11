package queue

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var tpl = `package {{.package}}

import (
	"fmt"
	"github.com/wuyan94zl/go-api/app/queue"
	"time"
)

type Queue struct {
	queue.BaseQueue
}

func NewQueue() Queue {
	return Queue{}
}

func (q Queue) Push(second int64) {
	q.Time = time.Now().Unix() + second
	queue.JobIns.Push(q)
}

func (q Queue) Run() {
	fmt.Println("执行队列程序：{{.package}}", q.Time)
	time.Sleep(1 * time.Second)
}

`

type Command struct {
	Name string
}

func (c *Command) GetDir() string {
	return utils.GetDir("queue", c.Name)
}

func (c *Command) Run() error {
	err := utils.GenFile(utils.FileGenConfig{
		Dir:          c.GetDir(),
		Filename:     "queue.go",
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
