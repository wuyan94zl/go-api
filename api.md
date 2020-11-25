
## 已有接口

### 添加用户

---
###### 请求URL：
> {{go-api-url}}/api/user/create
###### 请求方式
> POST
###### 请求参数
|参数名称|参数类型|默认值|
|-|-|-|
|email|text|admin@wuyan.com|
|password|text|123456|
|name|text|wuyan|
---
### 更新用户

---
###### 请求URL：
> {{go-api-url}}/api/user/update?id=8
###### 请求方式
> POST
###### 请求参数
|参数名称|参数类型|默认值|
|-|-|-|
|email|text|admin@wuyan.com|
|password|text|123456|
|name|text|无言123123|
---
### 删除用户

---
###### 请求URL：
> {{go-api-url}}/api/user/delete?id=3
###### 请求方式
> GET
---
### 一个用户

---
###### 请求URL：
> {{go-api-url}}/api/user/one?id=2
###### 请求方式
> GET
---
### 用户列表

---
###### 请求URL：
> {{go-api-url}}/api/user/list
###### 请求方式
> POST
###### 请求参数
|参数名称|参数类型|默认值|
|-|-|-|
|email|text|admin@wuyan.com|
|name|text|123456|
---
### 分页列表

---
###### 请求URL：
> {{go-api-url}}/api/user/paginate
###### 请求方式
> POST
###### 请求参数
|参数名称|参数类型|默认值|
|-|-|-|
|email|text|admin@wuyan.com|
|name|text|123456|
---
### 用户登录

---
###### 请求URL：
> {{go-api-url}}/api/user/login
###### 请求方式
> POST
###### 请求参数
|参数名称|参数类型|默认值|
|-|-|-|
|email|text|admin@wuyan.com|
|password|text|123456|
---
### 登录用户

---
###### 请求URL：
> {{go-api-url}}/api/user/info
###### 请求方式
> GET
###### 请求token
> bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiRXhwVGltZSI6MTYwNjEzMTQwOCwiZXhwIjoxNjA2MTMxNDA4LCJpc3MiOiJ3dXlhbiJ9.7k42N0mhdVo1yrf7y-1Kbb5HExDsVNquzo3-WbKOCeo
---