## go-api

### 使用
拉取代码到go工作区  
`git clone https://github.com/wuyan94zl/go-api`  

编译代码

### 代码片段
```go
// 定义查询数据
var conditions []model.Condition
// 设置name右模糊查询
conditions = model.SetCondition(conditions,"name","无言%","like")
// 设置age大于30
conditions = model.SetCondition(conditions,"age","30",">")
// 设置sex大于0（0：男，1：女） 第4个参数默认 =
conditions = model.SetCondition(conditions,"sex","0")

//执行查询0
user := model.GetOne(&user.User{}, conditions)
// select * from users where name like '无言%' and age > 30 and sex = 0 limit 1
//执行查询1
users := model.GetAll(&[]user.User{}, conditions)
// sql：select * from users where name like '无言%' and age > 30 and sex = 0
//执行查询2
users := model.GetAll(&[]user.User{}, conditions,10)
// sql：select * from users where name like '无言%' and age > 30 and sex = 0 limit 10
//执行查询3
users := model.GetAll(&[]user.User{}, conditions,10,5)
// sql：select * from users where name like '无言%' and age > 30 and sex = 0 LIMIT 10 OFFSET 5
//执行查询4（分页）
users := list := model.Paginate(&[]user.User, model.PageInfo{Page: 2, PageSize: 15}, conditions)
// sql：select count(*) from users where name like '无言%' and age > 30 and sex = 0 limit 1
// sql：select * from users where name like '无言%' and age > 30 and sex = 0 limit 15 OFFSET 15

```
