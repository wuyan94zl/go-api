package middleware

import (
	"github.com/gin-gonic/gin"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//tokenString := c.Request.Header.Get("Authorization")
		//admin := admin.Admin{}
		//id, err := admin.AuthToken(tokenString)
		//if err != nil {
		//	utils.SuccessErr(c, 401, err)
		//	c.Abort()
		//	return
		//}
		//mysql.GetInstance().First(&admin, id)
		//// 保存用户到 上下文
		//c.Set("auth", admin)
		//c.Set("auth_id", id)
		c.Next()
	}
}
