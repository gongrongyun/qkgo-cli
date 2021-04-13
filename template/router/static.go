package route

import (
	"github.com/gongrongyun/qkgo-cli/template/boot/http"
)

// Change the relativePath according to your demand
func AddStaticRoute() {
	http.Router.Static("/static", "./public")
}