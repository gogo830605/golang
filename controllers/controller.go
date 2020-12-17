package controller

import (
	// "net/http"
	"github.com/gin-gonic/gin"
)

func GetHello(ctx *gin.Context) {
	ctx.Data(200, "text/plain", []byte("Hello, It Home!"))
}