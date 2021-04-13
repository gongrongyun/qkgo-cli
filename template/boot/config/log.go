package config

import (
	"fmt"
)

const (
	logConfPath = "conf/log.json"
)

type logConf struct {
	Log struct {
		Dir       string `json:"dir"`
		Level     string `json:"level"`
		LogIDName string `json:"log_id_name"`
	} `json:"log"`
	Rotate struct {
		FileSuffixFormat string `json:"file_suffix_format"`
		ReverseTime      string `json:"reverse_time"`
		RotateTime       string `json:"rotate_time"`
	} `json:"rotate"`
}

var (
	logConfig *HotConfig = &HotConfig{
		ConfPtrFirst:  &logConf{},
		ConfPtrSecond: &logConf{},
	}
)

func InitLogConfig() {
	var err error
	_, err = GetHotLoadConfig(logConfPath, logConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func LogConfig() *logConf {
	return logConfig.GetConfig().(*logConf)
}
