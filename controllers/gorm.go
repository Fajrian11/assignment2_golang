package controllers

import (
	"assignment2_golang/repositories"
)

type InDB struct {
	OrderRepo repositories.OrderRepo
}

type PersonAPI struct {
	PersonRepo repositories.Person
}
