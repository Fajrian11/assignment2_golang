package config

import (
	"assignment2_golang/model"

	"github.com/jinzhu/gorm"
)

func DBinit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect database")
	}

	db.AutoMigrate(model.Orders{})
	return db
}
