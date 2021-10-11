package queue

import (
	"container/list"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type BaseQueue struct {
	Time int64
}

func (b BaseQueue) RunTime() int64 {
	return b.Time
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
	for true {
		ele := j.ll.Back()
		if ele == nil {
			break
		}
		queue := ele.Value.(Queue)
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
	Push(second int64)
	Run()
	RunTime() int64
}
