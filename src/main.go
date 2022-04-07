package main

import (
	"Config"
	"Models"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()

	var CashSurplusRule []Models.CashSurplusRule
	err := Models.GetRules(&CashSurplusRule)
	if err != nil {
		return
	}
	//fmt.Printf("Data:%+v:", CashSurplusRule)
	//router := Router.SetupRouter()
	//router.Run(":8888")
}
