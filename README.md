## 基于gin，gorm集成的 golang api 服务
作为一个后端 api 服务除了离不开的 crud 操作，还经常用到 定时任务，延时队列。  
该项目可以快速生成**表结构体**及对应该表的 **crud api** 接口  
还可以轻松快速的使用 **秒级定时任务**  
还可以轻松快速的使用 **秒级延时队列**
## 使用
- 拉取项目：`git clone https://github.com/wuyan94zl/go-api`  
- 修改 config.yml 对应的配置信息    
- 执行 `go run main.go`  

### 目录结构简单说明
```
|-app           app目录
|--command      定时任务代码
|--http         api功能代码
|--middleware   中间件代码
|--queue        延时队列代码
|-artisan       （不需要更改）
|-bootstrap     (不需要更改)
|-config        （不需要更改）
|-pkg           工具包
|-routes        （不需要更改）
```
> 所有开发基本都在app目录（功能代码开发） 或者pkg（工具包开发）
> 开发需要对gin,gorm熟悉


## 代码生成器 artisan
### 安装
`go get -u github.com/wuyan94zl/go-api/artisan`

执行 `artisan -h` 确保 artisan 代码生成器安装成功

## 构建 restful api
第一步：生成`restful api`的`json`配置文件，再跟进需要配置字段信息  
第二步：根据`json`配置文件生成`restful api`代码

### 生成`restful api`的`json`配置文件
命令：`artisan model task`
会创建 `app/http/task/model.api`文件
```json
{
  "package_name": "task",
  "struct_name": "Task",
  "fields": [
    {
      "field": "Id",
      "type_name": "uint64",
      "tags": {
        "json": "id"
      }
    },
	// 增加自定义字段 然后删除该行
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
```
增加以下字段信息
```json
    {
      "field": "UserId",
      "type_name": "uint64",
      "tags": {
        "json": "user_id",
        "auth": "auth"
      }
    },
    {
      "field": "Name",
      "type_name": "string",
      "tags": {
        "json": "name",
        "validate": "required"
      }
    },
    {
      "field": "Description",
      "type_name": "string",
      "tags": {
        "json": "description",
        "validate": "required"
      }
    },
    {
      "field": "Duration",
      "type_name": "int",
      "tags": {
        "json": "duration",
        "validate": "required||numeric"
      }
    },
    {
      "field": "StartTime",
      "type_name": "time.Time",
      "tags": {
        "json": "start_time",
        "validate": "required||date"
      }
    },
    {
      "field": "EndTime",
      "type_name": "time.Time",
      "tags": {
        "json": "end_time",
        "validate": "required||date"
      }
    },
```

### 根据`json`配置文件生成`restful api`代码
命令：`artisan api task`
执行后会在`app/http/task`文件夹生成三个go文件：
- controller.go 控制器代码文件
- model.go 模型代码文件
- service.go 服务代码文件

路由注册：
在`app/http/kernel.go` 文件 `Handle` 函数中增加 `Init()`
```go
import (
	"github.com/wuyan94zl/go-api/app/http/task" // 增加代码
	"github.com/wuyan94zl/go-api/routes" // 增加代码
)
func Handle() {
	task.Init(routes.AuthRouteGroup) // 增加代码 因为需要 auth 权限 所以传参 routes.AuthRouteGroup
}
```

运行程序
命令：`go run main.go`

![go-api artisan](https://cdn.learnku.com/uploads/images/202106/18/29943/hGh6BM602j.png!large)
好了，用postman 试试上面的5个api接口

## 构建定时任务

**命令**：`article cron task`
执行后会创建 `app/command/task` 文件夹,并生成 `handle.go` 文件 内容如下：
```go
package task

import (
	"fmt"
	"time"
)

type Job struct{}

func (j Job) Run() {
	fmt.Println("Execution per minute", time.Now().Format("2006-01-02 15:4:05"))
}

```

根据业务需求在`Run`函数中写入任务代码  
完成后在 app/command/kernel.go 文件 Handle 函数中增加任务调度
```go
import "github.com/wuyan94zl/go-api/app/command/task" // 增加代码
func Handle(c *cron.Cron) {
	//秒 分 时 天 月 年
	c.AddJob("0 * * * * *", task.Job{}) //增加代码
}

```
以上Handle 配置好后 控制台会每分钟执行一次 `handle.go` 中的 `Run` 函数

## 构建延时队列
**命令**：`artisan queue job`
执行后会生成 app/queue/data 文件夹 并生成 app/queue/name/queue.go 文件 内容如下：
```go
package job

import (
	"fmt"
	"github.com/wuyan94zl/go-api/app/queue/utils"
	"time"
)

var queueType = "job"

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
```
在Run函数中写业务代码
> `q.Data` 数据为发送队列是传递的参数
map[string]string 类型，下面代码中的 params数据

**队列注册** app/queue/actions.go
```go
package queue

import "github.com/wuyan94zl/go-api/app/queue/job" // 手动代码

func Action(method string, mapData map[string]string) Queue {
	switch method {
	case "job":// 手动代码
		return job.NewQueue(mapData)// 手动代码
	default:
		return nil
	}
}

```

**发送队列**  
在上面生成的定时任务`Run`函数中增加
```go
params := make(map[string]string)
params["data"] = "queue data params"
// 立即执行队列
job.NewQueue(params).Push()
// 延时10秒后执行队列
job.NewQueue(params).Push(10)
``` 
> 控制台会在定时任务执行后马上执行第一个队列，10秒后执行第二个队列

