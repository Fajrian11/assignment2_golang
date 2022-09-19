package database

import (
	"assignment2_golang/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func DBinit(dbHost, dbPort, dbUsername, dbPassword, dbName string) *gorm.DB {
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUsername, dbHost, dbPort, dbName)

	fmt.Println(dsn)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("ERROR: Failed to connect to Database -> %v\n", err)
	}
	db.AutoMigrate(model.Orders{})
	return db
}
