package Entities

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

func (CashSurplusRule *CashSurplusRule) TableName() string {
	return "cash_surplus_rule"
}
