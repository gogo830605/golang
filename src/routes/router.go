package router

import (
	"github.com/gin-gonic/gin"
	"controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.GetHello)
	return router
}