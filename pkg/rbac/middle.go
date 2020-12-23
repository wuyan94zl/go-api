package rbac

import (
	"github.com/gin-gonic/gin"
)

func PermissionCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.MustGet("auth_id").(int)
		url := c.Request.URL.Path
		CheckPermission(uint64(id), url)
	}
}
