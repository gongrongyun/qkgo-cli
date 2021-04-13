package orm

import (
	"template/boot/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var useConnPool bool

func InitOrm()  {
	engine := config.DatabaseConfig().Engine
	dbEngine, err := gorm.Open(engine, getParams(engine))
	if err != nil {
		panic(fmt.Errorf("Fatal error open database error [err=%s]\n", err))
	}
	db = dbEngine
	db.SingularTable(config.DatabaseConfig().SingularTable) // 禁止表名复数
	db.LogMode(config.DatabaseConfig().LogMode) // 会打印执行的 mysql 语句
	useConnPool = true
	db.DB().SetMaxOpenConns(config.DatabaseConfig().MaxOpenConn) //设置数据库连接池最大连接数
	db.DB().SetMaxIdleConns(config.DatabaseConfig().MaxIdleConn) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
}

func GetDB() *gorm.DB {
	return db
}

func EndOrm() {
	if useConnPool {
		return
	}
	err := db.Close()
	if err != nil {
		panic(fmt.Errorf("Fatal error close database error [err=%s]\n", err))
	}
}

func getParams(engine string) string {
	switch engine {
	case "mysql":
		host := config.DatabaseConfig().Host
		port := config.DatabaseConfig().Port
		dbname := config.DatabaseConfig().Dbname
		username := config.DatabaseConfig().User
		password := config.DatabaseConfig().Password
		mysqlParams := config.DatabaseConfig().MysqlParams
		connectTimeout := config.DatabaseConfig().ConnectTimeout
		readTimeout := config.DatabaseConfig().ReadTimeout
		writeTimeout := config.DatabaseConfig().WriteTimeout
		if mysqlParams == "" {
			mysqlParams = fmt.Sprintf("timeout=%s&readTimeout=%s&writeTimeout=%s",
				connectTimeout, readTimeout, writeTimeout)
		} else {
			mysqlParams = mysqlParams + fmt.Sprintf("&timeout=%s&readTimeout=%s&writeTimeout=%s",
				connectTimeout, readTimeout, writeTimeout)
		}
		params := fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
			username, password, host, port, dbname, mysqlParams)
		fmt.Println(params)
		return params
	case "sqlite3":
		params := config.DatabaseConfig().Dbname
		return params
	case "postgres":
		host := config.DatabaseConfig().Host
		port := config.DatabaseConfig().Port
		dbname := config.DatabaseConfig().Dbname
		username := config.DatabaseConfig().User
		password := config.DatabaseConfig().Password
		sslMode := config.DatabaseConfig().SslMode
		params := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host, port, username, password, dbname, sslMode)
		return params
	default:
		panic(fmt.Errorf("Fatal error getting database params: %s\n", engine))
	}
}
