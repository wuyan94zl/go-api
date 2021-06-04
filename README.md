## 基于gin，gorm集成的 golang api 服务

## 使用
git clone https://github.com/wuyan94zl/go-api  
修改 config.yml 对应的配置信息  
go run main.go

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


## 代码生成器 artisan 包含以下类

### 安装
`go get -u github.com/wuyan94zl/go-api/artisan`

### model 结构体 代码生成器
命令: `artisan model name`  
该命令会创建 app/http/name 文件夹，并生成app/http/name/model.json 文件，内容如下

```json
{
  "package_name": "article",
  "struct_name": "Article",
  "fields": [
    {
      "field": "Id",
      "type_name": "uint64",
      "tags": {
        "json": "id"
      }
    },
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

package_name：对应命令行中 name 值不做更改  
struct_name：对应模型中的结构体  
fields：对应模型字段

- field:字段名称对应结构体的字段
- type_name：字段类型
- tags：字段标签 对应结构体中的 json 标签

> validate 标签api字段验证规则 如 validate:"required||min:6||email"  
> 表示该字段必填长度不能小于6 必须字email类型，这个验证会自动生成

如上默认生成的结构体为

```go
type Article struct {
    Id        uint64    `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
```

根据需要修改json文件定义模型

### model 结构体 crud api 代码生成器
命令 `artisan api name`
> 该命令需要先执行 artisan model name 生成 model.json 文件才能处理

执行后会在app/http/name文件生成三个go文件：  
controller.go 控制器代码文件  
model.go 模型代码文件  
service.go 服务代码文件  

**路由注册**  
在app/http/kernel.go 文件 Handle 函数中增加 Init()
```go
import "github.com/wuyan94zl/go-api/app/http/name" // 增加代码
    func Handle() {
    name.Init() // 增加代码
}
```

### console 定时任务 代码生成器
命令：`article console name`  
执行后会创建 app/command/name 文件夹,并生成 `handle.go` 文件 内容如下：

```go
package name

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
完成后再 app/command/kernel.go 文件 Handle 函数中增加任务调度
```go
import "github.com/wuyan94zl/go-api/app/command/name" // 增加代码
func Handle(c *cron.Cron) {
	//秒 分 时 天 月 年
	c.AddJob("0 * * * * *", name.Job{}) //增加代码
}

```
### queue 延时队列 代码生成器
命令：`artisan queue data`  
执行后会生成 app/queue/data 文件夹 并生成 app/queue/name/queue.go 文件 内容如下：
```go
package data

import (
	"fmt"
	"time"
	"github.com/wuyan94zl/go-api/app/queue/utils"
)

var queueType = "data"

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
添加队列 params == 上面Run函数中的d.Data (map[string]string 类型)
```go
params := make(map[string]string)
params["data"] = "queue data params"
// 立即执行队列
data.NewQueue(params).Push()
// 延时10秒后执行队列
data.NewQueue(params).Push(10)
```

