package utils

import (
	"encoding/json"
	redis "github.com/wuyan94zl/redigo"
	"sync"
	"time"
)

const MyQueueKey = "wuyan94zl:queue:list"

var mutexW sync.Mutex

type queue struct {
	QueueType string            `json:"queue_type"`
	QueueUnix int64             `json:"queue_unix"`
	QueueData map[string]string `json:"queue_data"`
}

func Push(queueType string, data map[string]string, second ...int64) {
	mutexW.Lock()
	var q queue
	q.QueueType = queueType
	score := GetScore(second...)
	q.QueueUnix = score
	q.QueueData = data
	jsonByte, _ := json.Marshal(q)
	set := redis.SortSet{
		Score:  score,
		Member: string(jsonByte),
	}
	redis.ZAdd(MyQueueKey, set)
	mutexW.Unlock()
}

func GetScore(second ...int64) int64 {
	nano := time.Now().UnixNano()
	if len(second) > 0 {
		nano += second[0] * 1000000000
	}
	return nano / 1000
}
