package route

import (
	"github.com/gongrongyun/qkgo-cli/template/boot/http"
	"github.com/gongrongyun/qkgo-cli/template/controller/version"
)

func init()  {
	AddStaticRoute()
	AddRouter()
}

// Add your route here
func AddRouter() {
	http.Router.GET("/", version.Version)
}