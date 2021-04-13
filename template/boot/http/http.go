package http

import (
	"qkgo-template/boot/config"
	"qkgo-template/utils"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

var http *Http
var Router *gin.Engine

type Http struct {
	server *gin.Engine
	port   string
	addr   string
}

var DefaultMiddleWares []gin.HandlerFunc

func InitHttp() {
	http = new(Http)
	http.server = gin.New()
	Router = http.server

	if len(DefaultMiddleWares) > 0 {
		Router.Use(DefaultMiddleWares...)
	}
	// init
	needPprof := config.GlobalConfig().OpenPprof
	if needPprof {
		adminGroup := Router.Group("/tools", func(ctx *gin.Context) {
			pprofToken := config.GlobalConfig().PprofToken
			token, err := ctx.Cookie("pprof_token")
			if err != nil {
				token = ""
			}
			if ctx.DefaultQuery("token", "") != pprofToken && token != pprofToken {
				utils.Failure(ctx, utils.NoPermGetPprof)
				ctx.Abort()
				return
			}
			ctx.SetCookie("pprof_token", pprofToken, 3600, "/", "", false, true)
			ctx.Next()
		})
		pprof.RouteRegister(adminGroup, "pprof")
	}
	http.port = config.GlobalConfig().Port
	http.addr = config.GlobalConfig().Server
}

func Run() {
	err := Router.Run(http.addr + ":" + http.port)
	if err != nil {
		panic(fmt.Errorf("Fatal error run http server: %s\n", err))
	}
}
