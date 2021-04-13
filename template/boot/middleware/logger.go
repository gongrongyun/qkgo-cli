package middleware

import (
	"qkgo-template/boot/logger"
	"qkgo-template/boot/timer"
	"qkgo-template/boot/trace"
	"qkgo-template/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func PrintLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = trace.GetLogId(ctx)
		timer.StartTimer(ctx, "total")
		fmt.Println(ctx.FullPath())
		logger.PushNotice(ctx, "url", ctx.FullPath())
		defer func() {
			timer.StopTimer(ctx, "total")
			logger.Log.ServiceLogEngine.Infoln(getNoticeWithTimer(ctx))
			printAccessLog(ctx)
		}()
		ctx.Next()
	}
}

func printAccessLog(ctx *gin.Context) {
	logger.Log.AccessLogEngine.Infoln(getAccessLogMsg(ctx))
}

var accessLogKeys = []string {
	"logId", "status", "errno", // "errno", "errMsg", "protocol",
	"method", "uri",
	"client_ip", "host", "refer", "cost",
	//"local_ip", "uri", "protocol
}

var handleMap = map[string]func(ctx *gin.Context) string {
	"logId": trace.GetLogId,
	"status": func (ctx *gin.Context) string {
		return strconv.Itoa(ctx.Writer.Status())
	},
	"errno": func (ctx *gin.Context) string {
		if v, ok := ctx.Get(utils.ErrnoKey); ok {
			return v.(string)
		}
		return "0"
	},
	"method": func (ctx *gin.Context) string {
		return ctx.Request.Method
	},
	"uri": func (ctx *gin.Context) string {
		return ctx.Request.RequestURI
	},
	"client_ip": func (ctx *gin.Context) string {
		return ctx.ClientIP()
	},
	"host": func (ctx *gin.Context) string {
		return ctx.Request.Host
	},
	"refer": func (ctx *gin.Context) string {
		return ctx.Request.Referer()
	},
	"cost": func(ctx *gin.Context) string {
		return timer.GetTimer(ctx, "total")
	},
}

func getAccessLogMsg(ctx *gin.Context) string {
	builder := strings.Builder{}

	for _, filed := range accessLogKeys {
		handler := handleMap[filed]
		builder.WriteString(filed)
		builder.WriteByte('[')
		builder.WriteString(handler(ctx))
		builder.WriteByte(']')
	}

	return builder.String()
}

func getNoticeWithTimer(ctx *gin.Context) string {
	logString := logger.GetLogString(ctx)
	timerString := timer.GetTimerString(ctx)

	return logString + timerString
}
