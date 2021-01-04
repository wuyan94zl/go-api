package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/models/admin"
	"github.com/wuyan94zl/api/pkg/orm"
	"github.com/wuyan94zl/api/pkg/utils"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		admin := admin.Admin{}
		id, err := admin.AuthToken(tokenString)
		if err != nil {
			utils.SuccessErr(c, 401, err)
			c.Abort()
			return
		}
		orm.GetInstance().First(&admin, id)
		// 保存用户到 上下文
		c.Set("auth", admin)
		c.Set("auth_id", id)
		c.Next()
	}
}
