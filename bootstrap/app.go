package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin" // 基于 gin 框架
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"github.com/wuyan94zl/go-api/app/command"
	"github.com/wuyan94zl/go-api/app/http"
	"github.com/wuyan94zl/go-api/app/queue"
	"github.com/wuyan94zl/go-api/app/queue/test"
	"github.com/wuyan94zl/go-api/config"
	"github.com/wuyan94zl/go-api/routes"
	"github.com/wuyan94zl/mysql"
	"github.com/wuyan94zl/redigo"
)

func init() {
	config.InitConfig("config")
	mysqlInit()
	redisInit()
}

func Start() *gin.Engine {
	http.Handle()
	return routes.Route
}

func Timer() {
	if viper.GetString("name") == "main"{
		c := cron.New(cron.WithSeconds())
		command.Handle(c)
		queue.Handle(c)
		c.Start()
		fmt.Println("[Timer-debug] Start cron on server")
		fmt.Println("[Timer-debug] Start queue on server")
		test.NewQueue(make(map[string]string)).Push(10)
		test.NewQueue(make(map[string]string)).Push(2)
		test.NewQueue(make(map[string]string)).Push(1)
		test.NewQueue(make(map[string]string{"dd"""})).Push()
	}
	select {}
}

func redisInit() {
	r := redis.Config{
		Host:        fmt.Sprintf("%s:%s", viper.GetString("redis.host"), viper.GetString("redis.port")),
		Password:    viper.GetString("redis.password"),
		MaxActive:   viper.GetInt("redis.max_active"),
		MaxIdle:     viper.GetInt("redis.max_idle"),
		IdleTimeout: viper.GetDuration("redis.idle_timeout"),
	}
	redis.ConRedis(r)
}

func mysqlInit() {
	c := mysql.Config{
		Host:           viper.GetString("mysql.host"),
		Port:           viper.GetUint32("mysql.port"),
		Username:       viper.GetString("mysql.username"),
		Password:       viper.GetString("mysql.password"),
		Database:       viper.GetString("mysql.database"),
		Charset:        viper.GetString("mysql.charset"),
		MaxConnect:     100,
		MaxIdleConnect: 25,
		MaxLifeSeconds: 300,
	}
	mysql.ConMysql(c)
}
