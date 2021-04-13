package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

const ErrnoKey = "_errnoUnique"

func Failure(ctx *gin.Context, errno ErrCode) {
	errMsg, ok := ErrMsgInfo[errno]
	ctx.Set(ErrnoKey, strconv.Itoa(int(errno)))
	if !ok {
		ctx.JSON(int(errno), gin.H{"code": 0, "message": "unknown"})
	} else {
		ctx.JSON(errMsg.HttpCode, gin.H{"code": 0, "message": errMsg.ErrMsg})
	}
}
