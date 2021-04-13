package main

import (
	"template/boot/http"
	"template/boot/logger"
	mw "template/boot/middleware"
	"template/boot/orm"
	_ "template/config"
	"template/router"
)

func _init() {
	http.DefaultMiddleWares = mw.DefaultMiddleWares
	logger.InitLog()
	orm.InitOrm()
	http.InitHttp()
	router.Init()
}

func _end() {
	orm.EndOrm()
}

func main() {
	_init()
	http.Run()
	defer _end()
}
