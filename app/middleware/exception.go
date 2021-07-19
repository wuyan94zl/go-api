package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/go-api/pkg/logger"
	"github.com/wuyan94zl/go-api/pkg/response"
)

func Exception(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			switch data := r.(type) {
			case *response.Response:
				if data.Code == 200 {
					successData(c, data.Data)
				} else {
					successErr(c, data.Code, data.Message)
				}
			default:
				logger.Error(r)
				successErr(c, 500, fmt.Sprintf("%s%v", "系统错误：", r))
			}
			c.Abort()
		}
	}()
	c.Next()
}

func successData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "请求成功",
	})
}

func successErr(c *gin.Context, errCode int, msg interface{}) {
	c.JSON(500, gin.H{
		"code": errCode,
		"msg":  msg,
	})
}
