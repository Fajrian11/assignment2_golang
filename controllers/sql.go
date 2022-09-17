package controllers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type InDBOrders struct {
	*sql.DB
	OrderId      int
	CustomerName string
	OrderedAt    string
}

type InDBInputOrders struct {
	gorm.Model
	*gorm.DB
	OrderId      int
	CustomerName string
	OrderedAt    string
}

type InDBItems struct {
	*sql.DB
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}
