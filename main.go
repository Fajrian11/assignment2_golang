package main

import (
	"assignment2_golang/config"
	"assignment2_golang/controllers"
	"assignment2_golang/repositories"

	"github.com/gin-gonic/gin"
)

var baseURL = "https://hiyaa.site"

func main() {
	// FetchUser()
	db := config.DBinit()

	repoOrder := repositories.NewOrderRepo(db)
	inDB := &controllers.InDB{OrderRepo: repoOrder}
	// repoPerson := repositories.NewOrderPerson()
	// thirdparty := &controllers.PersonAPI{PersonRepo: repoPerson}

	router := gin.Default()

	router.GET("/orders", inDB.GetOrders)
	router.GET("/order/:id", inDB.GetOrderById)
	router.POST("/createOrder", inDB.CreateOrders)
	router.PUT("/updateOrder/:id", inDB.UpdateOrder)
	router.DELETE("/deleteOrder/:id", inDB.DeleteOrder)
	// router.GET("/order/person", thirdparty.GetPerson)

	// router.GET("/items", inDBItems.GetItems)

	router.Run(":3000")

	// inDBOrders.GetOrders()
	// inDBOrders.CreateOrders()

}
