package config

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sync"
	"time"
)

// 需要初始化两个配置的结构体指针传入两个配置
type HotConfig struct {
	confUseFirst 	bool
	ConfPtrFirst 	interface{}
	ConfPtrSecond 	interface{}
}

func init() {
	InitGlobalConfig()
	InitLogConfig()
	InitDatabaseConfig()
}

func (hc *HotConfig) setConfig(conf []byte) error {
	if hc.confUseFirst {
		err := json.Unmarshal(conf, hc.ConfPtrSecond)
		if err != nil {
			return err
		}
		hc.confUseFirst = false
	} else {
		err := json.Unmarshal(conf, hc.ConfPtrFirst)
		if err != nil {
			return err
		}
		hc.confUseFirst = true
	}
	return nil
}

func (hc *HotConfig) GetConfig() interface{} {
	if hc.confUseFirst {
		return hc.ConfPtrFirst
	} else {
		return hc.ConfPtrSecond
	}
}

type hotLoadConf struct {
	filename string
	hotConfig *HotConfig
	structType reflect.Type
	loadTime time.Time
}

var (
	mu sync.Mutex
	checkDuration = time.Second
	hotLoadArray = make([]*hotLoadConf, 0, 16)
	startHotLoad = false
)

func SetCheckDuration(duration time.Duration) {
	checkDuration = duration
}

func GetHotLoadConfig(filename string, config *HotConfig) (*HotConfig, error) {
	mu.Lock()
	defer mu.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("config: open file %s error: %v", filename, err)
	}

	fileInfo, _ := file.Stat()
	body := make([]byte, fileInfo.Size())
	n, err := file.Read(body)
	if err != nil {
		return nil, fmt.Errorf("config: read config %s error: %v", filename, err)
	}

	err = config.setConfig(body[:n])
	if err != nil {
		return nil, fmt.Errorf("config: decode config %s error: %v", filename, err)
	}

	if startHotLoad == false {
		go checkReloadConf()
		startHotLoad = true
	}

	hotLoadArray = append(hotLoadArray, &hotLoadConf{
		filename:	filename,
		hotConfig:  config,
		loadTime:   fileInfo.ModTime(),
	})

	return config, nil
}

func GetConfig(filename string, config interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("config: open file %s error: %v", filename, err)
	}

	fileInfo, _ := file.Stat()
	body := make([]byte, fileInfo.Size())
	n, err := file.Read(body)
	if err != nil {
		return fmt.Errorf("config: read config %s error: %v", filename, err)
	}

	err = json.Unmarshal(body[:n], config)
	if err != nil {
		return fmt.Errorf("config: decode config %s error: %v", filename, err)
	}

	return nil
}

func checkReloadConf() {
	for {
		time.Sleep(checkDuration)
		fmt.Println("start check conf...")
		checkAllHotConf()
	}
}

func checkAllHotConf() {
	mu.Lock()
	defer mu.Unlock()
	for _, v := range hotLoadArray {
		file, err := os.Open(v.filename)
		if err != nil {
			continue
		}

		fileInfo, _ := file.Stat()
		if v.loadTime == fileInfo.ModTime() {
			_ = file.Close()
			continue
		}
		fmt.Println("reload conf", v.filename)
		body := make([]byte, fileInfo.Size())
		n, err := file.Read(body)
		if err != nil {
			_ = file.Close()
			continue
		}

		err = v.hotConfig.setConfig(body[:n])
		if err != nil {
			_ = file.Close()
			fmt.Println("reload hot config ", v.filename, " error, err=", err)
			continue
		}
		v.loadTime = fileInfo.ModTime()
		_ = file.Close()
	}
}
