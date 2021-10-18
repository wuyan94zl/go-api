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
	running bool
	lock    sync.Mutex
	ll      *list.List
	llCopy  *list.List
}

var JobIns = &Job{running: false, ll: list.New(), llCopy: list.New()}

func (j *Job) Run() {
	if !j.start() {
		return
	}
	defer j.end()
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
}

func (j *Job) Push(queue Queue) {
	j.ll.PushFront(queue)
}

func (j *Job) start() bool {
	j.lock.Lock()
	defer j.lock.Unlock()
	if j.running == false {
		j.running = true
		return true
	}
	return false
}

func (j *Job) end() {
	j.running = false
}

func Handle(c *cron.Cron) {
	c.AddJob("* * * * * *", JobIns)
}

type Queue interface {
	Push(second int64)
	Run()
	RunTime() int64
}
