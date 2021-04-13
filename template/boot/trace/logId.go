package trace

import (
	"qkgo-template/boot/config"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const logIdInContext = "_log_id_unique"

func GetLogId(ctx *gin.Context) string {
	logIdCtx := ctx.Value(logIdInContext)
	logId := ""
	if logIdCtx == nil {
		logIdName := config.LogConfig().Log.LogIDName
		logId := ctx.GetHeader(logIdName)
		if logId == "" {
			logId = generateLogId()
		}
		ctx.Header(logIdName, logId)
		ctx.Set(logIdInContext, logId)
	} else {
		logId = logIdCtx.(string)
	}
	return logId
}

func generateLogId() string {
	return uuid.NewV4().String()
}
