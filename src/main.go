package main

import (
	"Models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	//if err != nil {
	//	fmt.Println("Status:", err)
	//}
	//defer Config.DB.Close()

	var CashSurplusRule Models.CashSurplusRule
	result, err := CashSurplusRule.GetRuleByID("3")
	if err != nil {
		return
	}

	fmt.Printf("Data:%+v:", result)

	//router := Router.SetupRouter()
	//router.Run(":8888")
}
