package middleware
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/auth"
	"github.com/wuyan94zl/api/pkg/utils"
	"strings"
)
func ApiAuth() gin.HandlerFunc  {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			utils.SuccessErr(c,401,"未登录")
			c.Abort()
			return
		}
		kv := strings.Split(tokenString, " ")
		if kv[0] != "Bearer"{
			utils.SuccessErr(c,401,"Token 错误")
			c.Abort()
			return
		}
		info, err := auth.GetUser(kv[1])
		if err != nil {
			fmt.Println(err)
			utils.SuccessErr(c,401,"登录已失效")
			c.Abort()
			return
		}else {
			// 保存用户到 上下文
			c.Set("admin",info)
			c.Next()
		}
	}
}