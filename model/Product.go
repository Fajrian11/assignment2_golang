package model

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string `json:"title" form:"title" validation:"required"`
	Description string `json:"Description" form:"description" validation:"required"`
	UserID      uint
	User        User
}

// validasi field field di database
func (p *Product) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

// validasi field field di database
func (p *Product) BeforeUpdate() (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}
