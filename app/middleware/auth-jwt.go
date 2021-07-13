package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/jwt"
	"github.com/wuyan94zl/go-api/pkg/response"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		jwtData := jwt.Jwt{}
		id, err := jwtData.AuthToken(tokenString)
		if err != nil {
			response.Error(401, err)
			c.Abort()
		}
		// 保存用户到 上下文
		c.Set("auth_id", id)
		c.Next()
	}
}
