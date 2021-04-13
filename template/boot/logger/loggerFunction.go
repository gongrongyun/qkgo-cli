package logger

import (
"template/boot/trace"
"fmt"
"github.com/gin-gonic/gin"
"strings"
)

type noticeMsg struct {
	noticeKey []string
	noticeValue []string
}

const noticeName = "_noticeMsgUnique"

func getLogIdMsg(ctx *gin.Context) string {
	return fmt.Sprintf("logId[%s]", trace.GetLogId(ctx))
}

func NoticeF(ctx *gin.Context, format string, args ...interface{}) {
	Log.ServiceLogEngine.Infof(getLogIdMsg(ctx) + format, args...)
}

func WarningF(ctx *gin.Context, format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Warningf(getLogIdMsg(ctx) + format, args...)
}

func DebugF(ctx *gin.Context, format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Debugf(getLogIdMsg(ctx) + format, args...)
}

func ErrorF(ctx *gin.Context, format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Errorf(getLogIdMsg(ctx) + format, args...)
}

func FatalF(ctx *gin.Context, format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Fatalf(getLogIdMsg(ctx) + format, args...)
}

func NoticeFNoTrace(format string, args ...interface{}) {
	Log.ServiceLogEngine.Infof(format, args...)
}

func WarningFNoTrace(format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Warningf(format, args...)
}

func DebugFNoTrace(format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Debugf(format, args...)
}

func ErrorFNoTrace(format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Errorf(format, args...)
}

func FatalFNoTrace(format string, args ...interface{}) {
	Log.ServiceWarningLogEngine.Fatalf(format, args...)
}

func PushNotice(ctx *gin.Context, key string, value interface{}) {
	msg := new(noticeMsg)
	tmpMsg, ok := ctx.Get(noticeName)
	strValue := fmt.Sprintf("%v", value)
	if !ok {
		msg.noticeKey = []string{key}
		msg.noticeValue = []string{strValue}
		ctx.Set(noticeName, msg)
		return
	}
	msg = tmpMsg.(*noticeMsg)
	msg.noticeKey = append(msg.noticeKey, key)
	msg.noticeValue = append(msg.noticeValue, strValue)
	ctx.Set(noticeName, msg)
}

func PrintNotice(ctx *gin.Context) {
	Log.ServiceLogEngine.Infoln(GetLogString(ctx))
}

func GetLogString(ctx *gin.Context) string {
	logString := strings.Builder{}

	logString.WriteString(getLogIdMsg(ctx))

	tmpMsg, ok := ctx.Get(noticeName)
	if !ok {
		return logString.String()
	}
	msg := tmpMsg.(*noticeMsg)
	for ind, key := range msg.noticeKey {
		logString.WriteString(key)
		logString.WriteByte('[')
		logString.WriteString(msg.noticeValue[ind])
		logString.WriteByte(']')
	}
	return logString.String()
}

