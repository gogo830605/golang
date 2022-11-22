package Router

import (
	CashSurplusRuleController "Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", CashSurplusRuleController.GetHello)
	grp1 := router.Group("/cash-surplus-rule")
	{
		grp1.GET("/", CashSurplusRuleController.CashSurplusRule)
		grp1.POST("/", CashSurplusRuleController.CreateRule)
		grp1.GET("/:id", CashSurplusRuleController.GetRuleByID)
		grp1.PUT("/:id", CashSurplusRuleController.UpdateRule)
		grp1.DELETE("/:id", CashSurplusRuleController.DeleteRule)
	}
	return router
}
