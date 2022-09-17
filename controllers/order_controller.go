package controllers

import (
	"assignment2_golang/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GORM

func (idb *InDB) CreateOrders(c *gin.Context) {
	var (
		order  []model.Orders
		result gin.H
	)

	order, err := idb.OrderRepo.CreateOrder(c)
	if err != nil {
		result = gin.H{
			"result": order,
		}
	} else {
		result = gin.H{
			"result": "created successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrderById(c *gin.Context) {
	var (
		order  []model.Orders
		result gin.H
	)
	order, err := idb.OrderRepo.GetOrderById(c)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": order,
			"count":  0,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetOrders(c *gin.Context) {
	var (
		orders []model.Orders
		result gin.H
	)
	orders, err := idb.OrderRepo.GetOrder()
	if len(orders) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
			"error":  err,
		}
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdateOrder(c *gin.Context) {
	var (
		orders []model.Orders
		result gin.H
	)
	orders, err := idb.OrderRepo.UpdateOrder(c)

	if err != nil {
		result = gin.H{
			"result": "Update Failed",
			"count":  orders,
		}
	} else {
		result = gin.H{
			"result": "Update Successfully",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeleteOrder(c *gin.Context) {
	var (
		orders []model.Orders
		// item   model.Items
		result gin.H
	)
	orders, err := idb.OrderRepo.DeleteOrder(c)
	// err2 = idb.DB.Unscoped().Delete(&item).Error // permanent delete with unscoped
	if err != nil {
		result = gin.H{
			"result": "Gagal Menghapus Data",
			"count":  orders,
		}
	} else {
		result = gin.H{
			"result": "Berhasil Menghapus Data",
		}
	}

	// if err2 != nil {
	// 	result = gin.H{
	// 		"result": "Gagal Menghapus Data",
	// 	}
	// } else {
	// 	result = gin.H{
	// 		"result": "Berhasil Menghapus Data",
	// 	}
	// }
	c.JSON(http.StatusOK, result)
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
