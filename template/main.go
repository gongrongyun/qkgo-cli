package template

import (
	"github.com/gongrongyun/qkgo-cli/template/boot/http"
	"github.com/gongrongyun/qkgo-cli/template/boot/logger"
	mw "github.com/gongrongyun/qkgo-cli/template/boot/middleware"
	"github.com/gongrongyun/qkgo-cli/template/boot/orm"
	_ "github.com/gongrongyun/qkgo-cli/template/config"
	_ "github.com/gongrongyun/qkgo-cli/template/router"
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
