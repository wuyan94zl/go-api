package queue

import (
	"encoding/json"
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/wuyan94zl/go-api/app/queue/utils"
	redis "github.com/wuyan94zl/redigo"
	"strconv"
	"sync"
)

const MyQueueKey = "wuyan94zl:queue:list"

type BaseQueue struct {
	QueueType string            `json:"queue_type"`
	QueueUnix int64             `json:"queue_unix"`
	QueueData map[string]string `json:"queue_data"`
}
type Job struct {
}

var mutexRun sync.Mutex

func (j *Job) Run() {
	mutexRun.Lock()
	jobData := j.pop()
	if len(jobData) > 0 {
		var queueData BaseQueue
		err := json.Unmarshal([]byte(jobData[0].Member), &queueData)
		if err != nil {
			return
		}
		queue := Action(queueData.QueueType, queueData.QueueData)
		queue.Run()
		_, err = redis.ZRemByScore(utils.MyQueueKey, jobData[0].Score, jobData[0].Score)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	mutexRun.Unlock()
}

func (j Job) pop() []redis.SortSet {
	jobData, _ := redis.ZRangeByScore(MyQueueKey, "0", strconv.FormatInt(utils.GetScore(), 10), "0", "1")
	return jobData
}

func Handle(c *cron.Cron) {
	c.AddJob("* * * * * *", &Job{})
}

type Queue interface {
	Push(second ...int64)
	Run()
}
