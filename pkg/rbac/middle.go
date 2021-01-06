package rbac

import (
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/pkg/utils"
)

func PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("auth_id").(uint64)
		url := c.Request.URL.Path
		if id == 1 || CheckPermission(id, url){
			c.Next()
		}else{
			utils.SuccessErr(c, 403, "您没有权限访问")
			c.Abort()
		}
	}
}
