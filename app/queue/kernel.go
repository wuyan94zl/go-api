package queue

import (
	"container/list"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type BaseQueue struct {
	QueueType string            `json:"queue_type"`
	QueueUnix int64             `json:"queue_unix"`
	QueueData map[string]string `json:"queue_data"`
}

// Run job 执行
func (b BaseQueue) Run() {

}

type Job struct {
	ll     *list.List
	llCopy *list.List
}

var JobIns = &Job{ll: list.New(), llCopy: list.New()}
var mutexRun sync.Mutex

func (j *Job) Run() {
	mutexRun.Lock()
	j.llCopy.Init()
	for i := j.ll.Len(); i > 0; i-- {
		ele := j.ll.Back()
		queue := ele.Value.(Queue)
		// 当前队列 的处理时间对比
		if queue.RunTime() <= time.Now().Unix() { // 处理队列
			queue.Run()
		} else { // 延时处理
			j.llCopy.PushFront(queue)
		}
		j.ll.Remove(ele)
	}
	j.ll, j.llCopy = j.llCopy, j.ll
	mutexRun.Unlock()
}

func (j *Job) Push(queue Queue) {
	j.ll.PushFront(queue)
}

func Handle(c *cron.Cron) {
	c.AddJob("* * * * * *", JobIns)
}

type Queue interface {
	Push(second ...int64)
	Run()
	RunTime() int64
}
