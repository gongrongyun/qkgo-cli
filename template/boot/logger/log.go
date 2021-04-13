package logger

import (
	"github.com/gongrongyun/qkgo-cli/template/boot/config"
	"github.com/gongrongyun/qkgo-cli/template/boot/rotateFile"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

type _log struct {
	AccessLogEngine *logrus.Logger
	ServiceLogEngine *logrus.Logger
	ServiceWarningLogEngine *logrus.Logger
}

var Log *_log

func InitLog() {
	dir := config.LogConfig().Log.Dir
	appName := config.GlobalConfig().AppName
	exist, err := pathExists(dir)
	if err != nil {
		panic(fmt.Errorf("Fatal error check config dir: %s\n", err))
	}
	if !exist {
		errMkdir := os.Mkdir(dir, os.ModePerm)
		if errMkdir != nil {
			panic(fmt.Errorf("Fatal error check config dir: %s\n", errMkdir))
		}
	}
	Log = new(_log)
	dir = trimSuffix(dir, "/")

	// init log engine
	initLogEngine(dir, "access", &Log.AccessLogEngine, false)
	initLogEngine(dir, appName, &Log.ServiceLogEngine, false)
	initLogEngine(dir, appName + ".wf", &Log.ServiceWarningLogEngine, true)
	if Log.AccessLogEngine == nil || Log.ServiceWarningLogEngine == nil || Log.ServiceLogEngine == nil {
		panic("init log error, nil get")
	}
}

func initLogEngine(dir, logName string, logger **logrus.Logger, needLine bool)  {
	file, err := rotateFile.Open(dir + "/" + logName + ".log") // os.OpenFile(dir + "/" + logName + ".log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0755)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if needLine {
		*logger = &logrus.Logger{
			Out:       file,
			Level:     getLogLevel(),
			Formatter: NewTextFormat(),
			Hooks:     logrus.LevelHooks{},
		}
		filenameHook := NewLineHook()
		filenameHook.Field = "line"
		(*logger).Hooks.Add(filenameHook)
	} else {
		*logger = &logrus.Logger{
			Out:       file,
			Level:     getLogLevel(),
			Formatter: NewTextFormat(),
		}
	}
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s) - len(suffix)]
	}
	return s
}

func getLogLevel() logrus.Level {
	logLevelString := config.LogConfig().Log.Level
	switch logLevelString {
	case "Debug":
		return logrus.DebugLevel
	case "Info":
		return logrus.InfoLevel
	case "Warning":
		return logrus.WarnLevel
	case "Error":
		return logrus.ErrorLevel
	case "Panic":
		return logrus.PanicLevel
	case "Fatal":
		return logrus.FatalLevel
	default:
		return logrus.WarnLevel
	}
}
