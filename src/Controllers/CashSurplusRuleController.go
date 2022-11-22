package CashSurplusRuleController

import (
	"Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHello(ctx *gin.Context) {
	ctx.Data(200, "text/plain", []byte("Hello, It Home!"))
}

func CashSurplusRule(ctx *gin.Context) {
	var CashSurplusRule Models.CashSurplusRule
	result, err := CashSurplusRule.GetRules()
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func CreateRule(ctx *gin.Context) {
	err := Models.CreateRule(ctx)
	if err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, "ok")
	}
}

func GetRuleByID(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	var CashSurplusRule Models.CashSurplusRule
	result, err := CashSurplusRule.GetRuleByID(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, result)
	}
}

func UpdateRule(ctx *gin.Context) {
	var CashSurplusRule Models.CashSurplusRule
	id := ctx.Params.ByName("id")
	err := Models.GetRuleByID(&CashSurplusRule, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, CashSurplusRule)
	}
	ctx.BindJSON(&CashSurplusRule)
	err = Models.UpdateRule(&CashSurplusRule, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, CashSurplusRule)
	}
}

func DeleteRule(ctx *gin.Context) {
	var CashSurplusRule Models.CashSurplusRule
	id := ctx.Params.ByName("id")
	err := Models.DeleteRule(&CashSurplusRule, id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
