package config

import (
	"fmt"
)

const (
	databaseConfPath = "conf/database.json"
)

type databaseConf struct {
	ConnectTimeout string `json:"connect_timeout"`
	Dbname         string `json:"dbname"`
	Engine         string `json:"engine"`
	Host           string `json:"host"`
	MysqlParams    string `json:"mysqlParams"`
	Password       string `json:"password"`
	Port           string `json:"port"`
	SslMode		   string `json:"ssl_mode"`
	ReadTimeout    string `json:"read_timeout"`
	User           string `json:"user"`
	WriteTimeout   string `json:"write_timeout"`

	SingularTable  bool   `json:"singular_table"`
	LogMode		   bool   `json:"log_mode"`
	MaxOpenConn    int    `json:"max_open_conn"`
	MaxIdleConn	   int	  `json:"max_idle_conn"`
}

var (
	databaseConfig *HotConfig = &HotConfig{
		ConfPtrFirst:  &databaseConf{},
		ConfPtrSecond: &databaseConf{},
	}
)

func InitDatabaseConfig() {
	var err error
	_, err = GetHotLoadConfig(databaseConfPath, databaseConfig)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DatabaseConfig() *databaseConf {
	return databaseConfig.GetConfig().(*databaseConf)
}
