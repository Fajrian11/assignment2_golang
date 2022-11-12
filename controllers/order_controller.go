package controllers

import (
	"assignment2_golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderController struct { // implementasi Controller
	ors service.OrderServiceApi
}

func NewOrderController(ors service.OrderServiceApi) *OrderController {
	return &OrderController{ors: ors}
}

func (oc *OrderController) GetOrdersControllers(c *gin.Context) {
	res := oc.ors.GetOrderService(c)
	c.JSON(http.StatusOK, res)
}

func (oc *OrderController) GetOrderByIdControllers(c *gin.Context) {
	res := oc.ors.GetOrderByIdService(c)
	c.JSON(http.StatusOK, res)
}

func (oc *OrderController) CreateOrderCOntrollers(c *gin.Context) {
	res := oc.ors.CreateOrderService(c)
	c.JSON(http.StatusOK, res)
}

func (oc *OrderController) UpdateOrderControllers(c *gin.Context) {
	res := oc.ors.UpdateOrderService(c)
	c.JSON(http.StatusOK, res)
}

func (oc *OrderController) DeleteOrderControllers(c *gin.Context) {
	res := oc.ors.DeleteOrderService(c)
	c.JSON(http.StatusOK, res)
}

func (oc *OrderController) UserRegisterControllers(c *gin.Context) {
	res := oc.ors.UserRegisterService(c)
	c.JSON(http.StatusOK, res)
}
func (oc *OrderController) UserLoginControllers(c *gin.Context) {
	res := oc.ors.UserLoginService(c)
	c.JSON(http.StatusOK, res)
}

// func (p *PersonAPI) GetPerson(c *gin.Context) {
// 	var (
// 		persons []model.Person
// 		result  gin.H
// 	)
// 	persons, err := p.PersonRepo.GetPerson()
// 	if len(persons) <= 0 {
// 		result = gin.H{
// 			"result": nil,
// 			"count":  0,
// 			"error":  err,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": persons,
// 			"count":  len(persons),
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)
// }

// golang sql

// func (idb *InDBOrders) GetOrders(c *gin.Context) {
// 	db, err := config.Connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query(`
// 		select
// 		orders.orderId,
// 		orders.customerName,
// 		orders.orderedAt,
// 		items.itemId,
// 		items.itemCode,
// 		items.description,
// 		items.quantity
// 		from orders join items on orders.orderId = items.orderId`)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	var (
// 		result_order []Orders
// 		result       gin.H
// 	)

// 	for rows.Next() {
// 		var each = Orders{}
// 		var err = rows.Scan(
// 			&each.OrderId,
// 			&each.CustomerName,
// 			&each.OrderedAt,
// 			&each.ItemId,
// 			&each.ItemCode,
// 			&each.Description,
// 			&each.Quantity,
// 		)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		result_order = append(result_order, each)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	if len(result_order) <= 0 {
// 		result = gin.H{
// 			"result": nil,
// 			"count":  0,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": result_order,
// 			"count":  len(result_order),
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)
// 	// for _, each := range result_order {
// 	// 	fmt.Println("Order ID : ", each.OrderId)
// 	// 	fmt.Println("Customer Name : ", each.CustomerName)
// 	// 	fmt.Println("Ordered At : ", each.OrderedAt)
// 	// 	fmt.Println("Order ID : ", each.ItemId)
// 	// 	fmt.Println("Order ID : ", each.ItemCode)
// 	// 	fmt.Println("Order ID : ", each.Description)
// 	// 	fmt.Println("Order ID : ", each.Quantity)
// 	// 	fmt.Println("")
// 	// }
// }
