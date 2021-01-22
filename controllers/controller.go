package controller

import (
	// "net/http"
	// "fmt"
	"github.com/gin-gonic/gin"
	// "inwin/services"
)

func GetHello(ctx *gin.Context) {
	ctx.Data(200, "text/plain", []byte("Hello, It Home!"))
	// ctx.JSON(http.StatusOK, gin.H{
	// 	"狀態": "ok",
	// })
	// fmt.Println(service.InitDB())
}