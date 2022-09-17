package model

import (
	"github.com/jinzhu/gorm"
)

type Orders struct {
	gorm.Model
	Customer_Name string  `json:"customer_name"`
	Ordered_At    string  `json:"ordered_at"`
	Items         []Items `json:"items" gorm:"foreignKey:Order_Id; polymorphic:Items; polymorphicValue:master constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	// Person        []Person
}

type Items struct {
	gorm.Model
	Item_Code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
	Order_Id    int    `json:"order_id"`
}

type Person struct {
	// gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  int    `json:"username"`
	Phone     int    `json:"phone"`
	Email     int    `json:"email"`
	Uuid      int    `json:"uuid"`
}
