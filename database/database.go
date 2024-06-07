package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DBM *gorm.DB
var DBS *gorm.DB

func InitDatabase() error {
	dsnMaster := "root:liu050202@tcp(127.0.0.1:3306)/takeout?charset=utf8mb4&parseTime=True&loc=Local"
	dsnSlaver := "root:liu050202@tcp(1.13.245.35)/takeout?charset=utf8mb4&parseTime=True&loc=Local"

	dbm, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{})

	if err != nil {
		log.Printf("link to database master error = %v", err)
		return err
	}
	dbs, err := gorm.Open(mysql.Open(dsnSlaver), &gorm.Config{})

	if err != nil {
		log.Printf("link to database master error = %v", err)
		return err
	}

	log.Printf("DBM = %v", dbm)
	log.Printf("DBS = %v", dbs)

	sqlDBM, err := dbm.DB()
	if err != nil {
		log.Printf("set link pool error = %v", err)
		return err
	}

	sqlDBS, err := dbs.DB()
	if err != nil {
		log.Printf("set link pool error in DBS = %v", err)
		return err
	}

	sqlDBM.SetMaxIdleConns(10)
	sqlDBM.SetMaxOpenConns(100)
	sqlDBS.SetMaxIdleConns(10)
	sqlDBS.SetMaxOpenConns(100)

	DBM = dbm
	DBS = dbs

	return nil
}
