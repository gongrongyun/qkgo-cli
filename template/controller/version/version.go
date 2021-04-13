package version

import (
	"github.com/gin-gonic/gin"
)

func Version(ctx *gin.Context)  {
	ctx.JSON(200, map[string]interface{}{
		"version": "v1.0",
	})
}
