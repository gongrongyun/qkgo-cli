package config

import (
	"fmt"
)

const (
	globalConfPath = "conf/global.json"
)

type globalConf struct {
	AppName    string `json:"app_name"`
	OpenPprof  bool   `json:"open_pprof"`
	Port       string `json:"port"`
	PprofToken string `json:"pprof_token"`
	Server     string `json:"server"`
	AppKey     string `json:"app_key"`
	AppSecret  string `json:"app_secret"`
	TBUsername string `json:"tb_username"`
	TBPassword string `json:"tb_password"`
}

var (
	gConf *HotConfig = &HotConfig{
		ConfPtrFirst:  &globalConf{},
		ConfPtrSecond: &globalConf{},
	}
)

func InitGlobalConfig() {
	var err error
	_, err = GetHotLoadConfig(globalConfPath, gConf)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GlobalConfig() *globalConf {
	return gConf.GetConfig().(*globalConf)
}
