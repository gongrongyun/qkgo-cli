package route

import (
	"github.com/qiankaihua/ginDemo/Boot/Http"
)

// Change the relativePath according to your demand
func AddStaticRoute() {
	Http.Router.Static("/static", "./public")
}