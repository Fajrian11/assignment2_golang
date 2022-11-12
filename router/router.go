package router

import (
	"assignment2_golang/config"
	"assignment2_golang/controllers"
	"assignment2_golang/database"
	"assignment2_golang/middleware"
	"assignment2_golang/repositories"
	"assignment2_golang/service"

	"github.com/gin-gonic/gin"
)

func StartAPP() *gin.Engine {
	cfg := config.LoadConfig()
	db := database.DBinit(cfg.Database.Host, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
	orderRepo := repositories.NewOrderRepo(db)
	orderService := service.NewOrderService(&orderRepo)
	orderController := controllers.NewOrderController(orderService)
	// Product
	ProductRepo := repositories.NewProductRepo(db)
	ProductService := service.NewProductService(&ProductRepo)
	ProductController := controllers.NewProductController(ProductService)

	router := gin.Default()

	router.GET("/orders", orderController.GetOrdersControllers)
	router.GET("/order/:id", orderController.GetOrderByIdControllers)
	router.POST("/createOrder", orderController.CreateOrderCOntrollers)
	router.PUT("/updateOrder/:id", orderController.UpdateOrderControllers)
	router.DELETE("/deleteOrder/:id", orderController.DeleteOrderControllers)

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", orderController.UserRegisterControllers)
		userRouter.POST("/login", orderController.UserLoginControllers)
	}

	productRouter := router.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/create-product", ProductController.CreateProductControllers)
		productRouter.GET("/", ProductController.GetAllProductControllers)

		productRouter.GET("/:productId", middleware.ProductAuthorization(), ProductController.GetProductByIdControllers)
		productRouter.PUT("/update-product/:productId", middleware.ProductAuthorization(), ProductController.UpdateProductController)
		productRouter.DELETE("/delete-product/:productId", middleware.ProductAuthorization(), ProductController.DeleteProductController)
	}

	return router
}
