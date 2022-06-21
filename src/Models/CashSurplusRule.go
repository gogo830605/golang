package Models

import (
	"Config"
	"Entities"
	"github.com/gin-gonic/gin"
)

type CashSurplusRule struct {
}

func (CashSurplusRule CashSurplusRule) TableName() string {
	return "cash_surplus_rule"
}

func (CashSurplusRule CashSurplusRule) GetRules() ([]Entities.CashSurplusRule, error) {
	db, err := Config.GetDB()
	defer db.Close()

	if err != nil {
		return nil, err
	} else {
		var cashSurplusRules []Entities.CashSurplusRule
		db.Table("cash_surplus_rule").Select("*").Scan(&cashSurplusRules)
		return cashSurplusRules, nil
	}
}

func CreateRule(ctx *gin.Context) (err error) {
	db, err := Config.GetDB()
	defer db.Close()

	if err != nil {
		return err
	} else {
		var cashSurplusRules Entities.CashSurplusRule
		ctx.BindJSON(cashSurplusRules)
		db.Table("cash_surplus_rule").Select("*").Create(&cashSurplusRules)
		return nil
	}

	//if err = Config.DB.Create(CashSurplusRule).Error; err != nil {
	//	return err
	//}
	//return nil
}

func (CashSurplusRule CashSurplusRule) GetRuleByID(id string) (Entities.CashSurplusRule, error) {
	db, err := Config.GetDB()
	defer db.Close()

	if err != nil {
		return Entities.CashSurplusRule{}, err
	} else {
		var cashSurplusRules Entities.CashSurplusRule
		db.Table("cash_surplus_rule").Select("*").First(&cashSurplusRules, id)
		return cashSurplusRules, nil
	}
}

//
//func UpdateRule(CashSurplusRule *CashSurplusRule, id string) (err error) {
//	fmt.Println(CashSurplusRule)
//	Config.DB.Save(CashSurplusRule)
//	return nil
//}
//
//func DeleteRule(CashSurplusRule *CashSurplusRule, id string) (err error) {
//	Config.DB.Where("id = ?", id).Delete(CashSurplusRule)
//	return nil
//}
