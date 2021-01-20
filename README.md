## golang的rbac权限api管理服务

第一步： git clone https://github.com/wuyan94zl/GoApiServer  

第二部： cp .env.example .env 并 修改 .env 对应的配置信息

第三步： go run wuyan.go admin 和 go run wuyan.go admin route permission  

第四步： go run main.go  

第五步：执行数据库数据初始化sql文件（见项目根目录）

查看演示 [http://gorbacui.wuyan94zl.cn](http://gorbacui.wuyan94zl.cn)  

以上操作就是演示地址中的 api 所有功能

## 构建CURD工具使用
> 需求：增加文章管理功能
1、创建数据表模型
```go
type Article struct {
	Id          uint64 `gorm:"column:id;primaryKey;autoIncrement;not null"json:"id"`
	Title       string `validate:"required,min:10,max:50"search:"like"json:"title"`
	Description string `validate:"required,min:10,max:200"json:"description"`
	Content     string `validate:"required"json:"content"`
	View        uint64
	AdminId     uint64      `validate:"required,numeric"json:"admin_id"`
	Admin       admin.Admin `gorm:"-"relationship:"belongTo"json:"admin"`
	CreatedAt   time.Time   `gorm:"column:created_at;index"json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at"json:"updated_at"`
}
//valiedate标签控制create和update的参数验证，valiedate参数详情请看 github.com/thedevsaddam/govalidator
//如上创建的时候title和description必填且长度在10-50和10-200
//admin_id 必填且为数字

//search标签控制分页列表的查询方式（目前支持:like,=,>,<,!=。like为str%）
//如上分页列表接口接收title关键字like查询

//relationship 为关联，如上查询的时候会关联出admin用户信息
```  
2、bootstrap/auto_migrate.go `init` 函数中`MigrateStruct` map数据添加模型  
```go
MigrateStruct["article"] = Article{}
```
3、执行 `go run wuyan.go article`  和 `go run wuyan.go article route permission`  

> 以上操作后会增加文章的增、删、改、详细、分页数据5个接口  
增和改接口中字段验证为：  
title：必填，长度在10到50之间  
description：必填，长度在10到200之间  
content：必填  
admin_id：必填，必须是数字  
详细和分页数据中：  
数据会关联查询admin信息  
