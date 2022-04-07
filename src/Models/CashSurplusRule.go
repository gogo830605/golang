package Models

import (
	"Config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CashSurplusRule struct {
	Id     int
	SiteId int
	Lang   string
	Title  string
	//Content   string
	Active    int
	Sort      int
	CreatedBy int
}

func (b *CashSurplusRule) TableName() string {
	return "cash_surplus_rule"
}

func GetRules(CashSurplusRule *[]CashSurplusRule) (err error) {
	if err = Config.DB.Find(CashSurplusRule).Error; err != nil {
		return err
	}

	return nil
}

func CreateRule(CashSurplusRule *CashSurplusRule) (err error) {
	if err = Config.DB.Create(CashSurplusRule).Error; err != nil {
		return err
	}
	return nil
}

func GetRuleByID(CashSurplusRule *CashSurplusRule, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(CashSurplusRule).Error; err != nil {
		return err
	}
	return nil
}

func UpdateRule(CashSurplusRule *CashSurplusRule, id string) (err error) {
	fmt.Println(CashSurplusRule)
	Config.DB.Save(CashSurplusRule)
	return nil
}

func DeleteRule(CashSurplusRule *CashSurplusRule, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(CashSurplusRule)
	return nil
}
