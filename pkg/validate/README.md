
基于 `github.com/thedevsaddam/govalidator` 封装的一个验证器

### 安装
gitee  `go get gitee.com/wuyan94zl/govalidate`  
github `go get github.com/wuyan94zl/govalidate`  

### 使用
定义模型struct
```go
type blog struct {
	Id      uint64 `json:"id"`
	Title   string `json:"title"validate:"required||min:12||max:32"fieldName:"博客标题"`
	Content string `json:"content"validate:"required"fieldName:"博客内容"`
	View    uint64 `json:"view"validate:"numeric"fieldName:"浏览数"`
}
```
> 三个字段标签说明  
> json 定义后 验证时字段名称就为小写了，未定义json 验证字段首字母大写  
> validate 验证规则，多个规则以 || 隔开  
> fieldName 验证字段别名，以Title为例：`Title字段 不能为空`（默认），`博客标题 不能为空`（定义fieldName）

```go
// r 为 *http.Request

// struct 默认验证
ok, msg := govalidate.StructValidate(r,blog{})
if !ok {
    // 验证不通过，msg 为提示消息
}

// 需要临时增加或删除验证规则,title 最大长度不限制 和 view 必传
mapData, fieldMap := MapDataForStruct(blog{})
// view 添加 required 规则
mapData["view"] = append(mapData["view"], "required")
// title 删除 required（slice删除元素）
mapData["title"] = mapData["title"][0:1] // 删除第三个元素 max:32

// map自定义数据验证
ok, msg := govalidate.MapValidate(r, mapData, fieldMap)
if !ok {
    // 验证不通过，msg 为提示消息
}
```
 