package timer

import (
	"github.com/gongrongyun/qkgo-cli/template/boot/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

const timerName = "_timerMsgUnique"

// TODO: 目前多线程不安全，对于map访问没有加锁控制
type timer struct {
	keys []string
	timers map[string]time.Time
	durations map[string]string
	//TODO: timer 加锁?
	//updateLock sync.Mutex
}

func StartTimer(ctx *gin.Context, name string) {
	timerLocal := new(timer)
	tmpTimer, ok := ctx.Get(timerName)
	if !ok {
		timerLocal.keys = []string{name}
		timerLocal.timers = make(map[string]time.Time)
		timerLocal.durations = make(map[string]string)
		timerLocal.timers[name] = time.Now()
		ctx.Set(timerName, timerLocal)
		return
	}
	timerLocal = tmpTimer.(*timer)
	timerLocal.timers[name] = time.Now()
	if hasTimer(timerLocal, name) {
		return
	}
	timerLocal.keys = append(timerLocal.keys, name)
	//ctx.Set(timerName, timerLocal)
}

func StopAllTimer(ctx *gin.Context) {
	tmpTimer, ok := ctx.Get(timerName)
	if !ok {
		return
	}
	timerLocal := tmpTimer.(*timer)
	for k, _ := range timerLocal.timers {
		stopTimerLocal(timerLocal, k)
	}
}

func StopTimer(ctx *gin.Context, name string) {
	tmpTimer, ok := ctx.Get(timerName)
	if !ok {
		return
	}
	timerLocal := tmpTimer.(*timer)
	if !hasTimer(timerLocal, name) {
		logger.ErrorF(ctx, fmt.Sprintf("Timer Not Has Key: %s", name))
		return
	}
	stopTimerLocal(timerLocal, name)
}

func GetTimer(ctx *gin.Context, name string) string {
	tmpTimer, ok := ctx.Get(timerName)
	if !ok {
		return ""
	}
	timerLocal := tmpTimer.(*timer)
	if !hasTimer(timerLocal, name) {
		logger.ErrorF(ctx, fmt.Sprintf("Timer Has Not Has Key: %s", name))
		return ""
	}
	if value, ok := timerLocal.durations[name]; ok {
		return value
	}
	return time.Since(timerLocal.timers[name]).String()
}

func GetAllTimer(ctx *gin.Context) ([]string, map[string]string) {
	tmpTimer, ok := ctx.Get(timerName)
	if !ok {
		return []string{}, map[string]string{}
	}
	timerLocal := tmpTimer.(*timer)
	return timerLocal.keys, timerLocal.durations
}

func GetTimerString(ctx *gin.Context) string {
	logString := strings.Builder{}

	StopAllTimer(ctx)
	keys, times := GetAllTimer(ctx)

	for _, key := range keys {
		if timeString, ok := times[key]; ok {
			logString.WriteString(key)
			logString.WriteByte('[')
			logString.WriteString(timeString)
			logString.WriteByte(']')
		}
	}

	return logString.String()
}

func hasTimer(timerLocal *timer, name string) bool {
	if _, ok := timerLocal.timers[name]; ok {
		return true
	}
	return false
}

func stopTimerLocal(timerLocal *timer, name string) {
	if _, ok := timerLocal.timers[name]; !ok {
		return
	}
	if _, ok := timerLocal.durations[name]; ok {
		return
	}
	timerLocal.durations[name] = time.Since(timerLocal.timers[name]).String()
}
