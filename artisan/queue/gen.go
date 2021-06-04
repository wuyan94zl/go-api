package queue

import (
	"fmt"
	"github.com/wuyan94zl/go-api/artisan/utils"
)

var tpl = `package {{.package}}

import (
	"fmt"
	"time"
	"github.com/wuyan94zl/go-api/app/queue/utils"
)

var queueType = "{{.package}}"

func NewQueue(data map[string]string) *Queue {
	return &Queue{
		Data: data,
	}
}

type Queue struct {
	Data map[string]string
}

func (q *Queue) Push(second ...int64) {
	utils.Push(queueType, q.Data, second...)
}
func (q *Queue) Run() {
	fmt.Println("执行队列程序 参数为：", q.Data)
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
