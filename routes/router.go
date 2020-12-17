package router

import (
	"github.com/gin-gonic/gin"
	"inwin/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.GetHello)
	return router
}