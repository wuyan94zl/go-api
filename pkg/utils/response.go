package utils

import (
	"github.com/gin-gonic/gin"
)
// 参数 data interface{} 类型为可接受任意类型
func SuccessData(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  "请求成功",
	})
}
func SuccessErr(c *gin.Context, err_code int, msg interface{}){
	c.JSON(200, gin.H{
		"code": err_code,
		"msg":  msg,
	})
}