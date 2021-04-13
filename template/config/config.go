package config

import (
	"fmt"
	"github.com/gongrongyun/qkgo-cli/template/boot/config"
	"time"
)

func init() {
	fmt.Println("init config")
	config.SetCheckDuration(time.Minute)
	InitHttpConfig()
}

func HttpConfig() *httpConf {
	return httpConfig.GetConfig().(*httpConf)
}

