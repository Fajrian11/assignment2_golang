package config

import (
	"assignment2_golang/model"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() model.Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	serverPort := os.Getenv("SERVICE_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	config := model.Config{
		ServerPort: serverPort,
		Database: model.Database{
			Host:     dbHost,
			Port:     dbPort,
			Username: dbUsername,
			Password: dbPassword,
			Name:     dbName,
		},
	}
	return config
}

// func DBinit() *gorm.DB {
// 	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/orders_by?charset=utf8&parseTime=True&loc=Local")
// 	if err != nil {
// 		panic("Failed to connect database")
// 	}

// 	db.AutoMigrate(model.Orders{})
// 	return db
// }
