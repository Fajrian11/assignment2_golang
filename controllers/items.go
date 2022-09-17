package controllers

// model
type Items struct {
	ItemId      int
	ItemCode    string
	Description string
	Quantity    int
	OrderId     int
}

// func (idb *InDBItems) GetItems(c *gin.Context) {
// 	db, err := config.Connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("select * from items")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer rows.Close()

// 	var (
// 		result_items []Items
// 		result       gin.H
// 	)

// 	for rows.Next() {
// 		var each = Items{}
// 		var err = rows.Scan(&each.ItemId, &each.ItemCode, &each.Description, &each.Quantity, &each.OrderId)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}
// 		result_items = append(result_items, each)
// 	}

// 	err = rows.Err()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	if len(result_items) <= 0 {
// 		result = gin.H{
// 			"result": nil,
// 			"count":  0,
// 		}
// 	} else {
// 		result = gin.H{
// 			"result": result_items,
// 			"count":  len(result_items),
// 		}
// 	}
// 	c.JSON(http.StatusOK, result)
// 	// for _, each := range result_order {
// 	// 	fmt.Println("Order ID : ", each.orderId)
// 	// 	fmt.Println("Customer Name : ", each.customerName)
// 	// 	fmt.Println("Order ID : ", each.orderId)
// 	// 	fmt.Println("")
// 	// }

// }
