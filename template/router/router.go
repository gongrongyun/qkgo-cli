package route

import (
	"template/boot/http"
	"template/controller/version"
)

func init()  {
	AddStaticRoute()
	AddRouter()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", version.Version)
}