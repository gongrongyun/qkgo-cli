package route

import (
	"qkgo-template/boot/http"
	"qkgo-template/controller/version"
)

func init()  {
	AddStaticRoute()
	AddRouter()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", version.Version)
}