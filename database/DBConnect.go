package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DBconnect *gorm.DB

var err error

func DD() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:password@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
}