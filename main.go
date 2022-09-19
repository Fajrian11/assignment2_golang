package main

import (
	"assignment2_golang/config"
	"assignment2_golang/controllers"
	"assignment2_golang/database"
	"assignment2_golang/repositories"
	"assignment2_golang/service"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var baseURL = "https://hiyaa.site"

func main() {
	// FetchUser()
	cfg := config.LoadConfig()
	db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
	orderRepo := repositories.NewOrderRepo(db)
	orderService := service.NewOrderService(&orderRepo)
	orderController := controllers.NewOrderController(orderService)
	// repoPerson := repositories.NewOrderPerson()
	// thirdparty := &controllers.PersonAPI{PersonRepo: repoPerson}

	router := gin.Default()

	router.GET("/orders", orderController.GetOrdersControllers)
	router.GET("/order/:id", orderController.GetOrderByIdControllers)
	router.POST("/createOrder", orderController.CreateOrderCOntrollers)
	router.PUT("/updateOrder/:id", orderController.UpdateOrderControllers)
	router.DELETE("/deleteOrder/:id", orderController.DeleteOrderControllers)
	// router.GET("/order/person", thirdparty.GetPerson)

	// router.GET("/items", inDBItems.GetItems)

	router.Run(":3000")

	// inDBOrders.GetOrders()
	// inDBOrders.CreateOrders()

}
