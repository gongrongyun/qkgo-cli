package router

import (
	"template/boot/http"
	"template/controller/version"
)

func Init()  {
	AddStaticRoute()
	AddRouter()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", version.Version)
}