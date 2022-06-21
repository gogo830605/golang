package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

//
//// DBConfig represents db configuration
//type DBConfig struct {
//	Host     string
//	Port     int
//	User     string
//	DBName   string
//	Password string
//}
//
//func (dbConfig *DBConfig) BuildDBConfig() *DBConfig {
//	Config := DBConfig{
//		Host:     "10.1.30.25",
//		Port:     3306,
//		User:     "root",
//		Password: "password",
//		DBName:   "ark",
//	}
//	return &Config
//}
//
//func (dbConfig *DBConfig) DbURL(*DBConfig) string {
//	return fmt.Sprintf(
//		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
//		dbConfig.User,
//		dbConfig.Password,
//		dbConfig.Host,
//		dbConfig.Port,
//		dbConfig.DBName,
//	)
//}
//
//func (dbConfig *DBConfig) GetDB() (*gorm.DB, error) {
//	db, err := gorm.Open("mysql", dbConfig.DbURL(dbConfig.BuildDBConfig()))
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}

func GetDB() (*gorm.DB, error) {
	dbDriver := "mysql"
	Host := "10.1.30.25"
	Port := 3306
	User := "root"
	Password := "password"
	DBName := "ark"
	db, err := gorm.Open(dbDriver, fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		User,
		Password,
		Host,
		Port,
		DBName,
	))
	if err != nil {
		return nil, err
	}
	return db, nil
}
