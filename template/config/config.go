package config

import (
	"qkgo-template/boot/config"
	"fmt"
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

