package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

const (
	RawBodyKey		= "raw_body"
	FormatBodyKey	= "format_body"
)

func Format() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO:start timer
		defer func() {
			// TODO:end timer
			ctx.Next()
			// TODO: print errno
			//if code, ok := ctx.ErrNo.(uint32); ok && code != 0 {
			//	//TODO:warning
			//	fmt.Printf("raw_body: \n---\n%s\n---\n", ctx.GetString(RawBodyKey))
			//	//TODO:debug
			//	fmt.Printf("format_body: %s", ctx.GetString(FormatBodyKey))
			//}
		}()

		body, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			// TODO: warning
			fmt.Printf("format: read body error: %v", err)
		}
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		ctx.Set(RawBodyKey, string(body))
		// 暂时没有 format 需求
	}
}
