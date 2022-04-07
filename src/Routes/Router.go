package Router

import (
	CashSurplusRule "Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", CashSurplusRule.GetHello)
	grp1 := router.Group("/cash-surplus-rule")
	{
		grp1.GET("/", CashSurplusRule.CashSurplusRule)
		grp1.POST("/", CashSurplusRule.CreateRule)
		grp1.GET("/:id", CashSurplusRule.GetRuleByID)
		grp1.PUT("/:id", CashSurplusRule.UpdateRule)
		grp1.DELETE("/:id", CashSurplusRule.DeleteRule)
	}
	return router
}
