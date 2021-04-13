package main

import (
	"qkgo-template/boot/http"
	"qkgo-template/boot/logger"
	mw "qkgo-template/boot/middleware"
	"qkgo-template/boot/orm"
	_ "qkgo-template/config"
	_ "qkgo-template/router"
)

func _init() {
	http.DefaultMiddleWares = mw.DefaultMiddleWares
	logger.InitLog()
	orm.InitOrm()
	http.InitHttp()
}

func _end() {
	orm.EndOrm()
}

func main() {
	_init()
	http.Run()
	defer _end()
}
