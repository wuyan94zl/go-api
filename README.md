## 分分钟 构建一个完整rbac权限的 go-api 服务

### 使用
工作区直接拉取代码  
`git clone https://github.com/wuyan94zl/go-api`  

### 目录结构
```
app
    controllers #控制器
    middleware  #中间件
    models      #模型
boorstrap       #启动
config          #配置
pkg             #工具包
routes          #路由
.env.example
main.go         #入口
wuyan.go        #自动构建模型的CURD接口入口
```

### 开始
首先把 `.env.example` 改为 `.env` 并配置数据库连接，其他信息随情况修改  
默认项目包含2个接口，见路由文件 routes/api.go 和 routes/auth.go(需要认证的接口)  
1、登录 `api/admin/login`  
2、登录用户信息 `api/admin/auth`  

初始没有用户数据没法登录，使用CURD工具快速开始  
运行 `go run wuyan.go admin` 构建CURD控制器  
运行 `go run wuyan.go admin route` 构建CURD路由  
默认构建路由：  
`api/admin/create (POST)`,`api/admin/update (POST)`,`api/admin/delete (GET)`,`api/admin/info (GET)`,`api/admin/paginate (POST)`。  

完成后 运行 `go run main.go` 启动服务  
使用 `api/admin/create` 接口 POST参数:email,password,name 创建数据。  

再次使用登录接口`api/admin/login` POST email 和 password 登录。

成功后使用token 访问 `api/admin/auth` 获取登录用户信息。

### 构建CURD工具使用
#### 两步准备工作    
1、创建数据表模型
```go
type Admin struct {
	models.BaseModel
	Email    string `validate:"required,min:6,email"search:"like"`
	Password string `validate:"min:6"pwd:"pwd"`
	Name     string `validate:"required,min:6"search:"like"`
}
//valiedate标签控制create和update的参数验证，valiedate参数详情请看 github.com/thedevsaddam/govalidator
//search标签控制分页列表的查询方式（目前支持:like,=,>,<,!=。like为str%）
```  
2、bootstrap/auto_migrate.go `init` 函数中`MigrateStruct` map数据添加模型  
```
MigrateStruct["admin"] = admin.Admin{}
```
#### 使用
1、创建CURD控制器执行 `go run wuyan.go admin`  
> 默认目录 app/controllers/admin/curd.go,admin为你的模型名称小写  
> 执行 go run wuyan.go admin curd console,则路径为app/controllers/console/admin/curd.go
  
2、创建对应路由执行 `go run wuyan.go admin route`
> 默认路由添加在 routes/api.go 路由文件  
> 可执行 go run wuyan.go admin route auth,则路由添加在 routes/auth.go  
> 如果你CURD自定义了目录console,则需要传第四个参数为：go run wuyan.go admin route auth console  
> 如果文件不存在则添加失败，需手动增加，增加方式参考api和auth文件

  
注：该工具暂时只能使用单模型，模型关联（hasOne,hasMany等）暂时还不支持，后续可能会加上  
