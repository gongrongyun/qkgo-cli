package middleware

import "github.com/gin-gonic/gin"

var DefaultMiddleWares = []gin.HandlerFunc {
	Recovery(),

	PrintLog(),

	Recovery(),

	Format(),
}
