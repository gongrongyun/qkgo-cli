package config

import (
	"fmt"
	"github.com/gongrongyun/qkgo-cli/template/boot/config"
)

const (
	httpConfPath = "conf/http.json"
)

type httpConf struct {
	Timeout int64 `json:"timeout"`
}

var (
	httpConfig *config.HotConfig = &config.HotConfig{
		ConfPtrFirst:  &httpConf{},
		ConfPtrSecond: &httpConf{},
	}
)

func InitHttpConfig() {
	var err error
	_, err = config.GetHotLoadConfig(httpConfPath, httpConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
}
