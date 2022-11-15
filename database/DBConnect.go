package database

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DBconnecct *gorm.DB

var err error

func DD() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:solar999@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnecct, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
}