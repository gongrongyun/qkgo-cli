package router

import (
	"template/boot/http"
)

// Change the relativePath according to your demand
func AddStaticRoute() {
	http.Router.Static("/static", "./public")
}