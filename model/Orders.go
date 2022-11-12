package model

import (
	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	Customer_Name string  `json:"customer_name"`
	Ordered_At    string  `json:"ordered_at"`
	Items         []Items `json:"items" gorm:"foreignKey:Order_Id"`
	Person        Result  `json:"person"`
}

type Items struct {
	gorm.Model
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Order_Id    int    `json:"order_id"`
}

type Result struct {
	Person []Person `json:"result"`
}

type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Uuid      string `json:"uuid"`
}
