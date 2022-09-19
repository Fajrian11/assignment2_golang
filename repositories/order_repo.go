package repositories

import (
	"assignment2_golang/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return OrderRepo{db: db}
}

type OrderRepoApi interface {
	GetOrder(c *gin.Context) ([]model.Orders, error)
	GetOrderById(c *gin.Context) ([]model.Orders, error)
	CreateOrder(c *gin.Context) ([]model.Orders, error)
	UpdateOrder(c *gin.Context) ([]model.Orders, error)
	DeleteOrder(c *gin.Context) ([]model.Orders, error)
}

func (or *OrderRepo) GetOrder(c *gin.Context) ([]model.Orders, error) {
	var order = []model.Orders{}

	err := or.db.Model(&model.Orders{}).Preload("Items").Find(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return order, nil
}

func (or *OrderRepo) GetOrderById(c *gin.Context) ([]model.Orders, error) {
	var order = []model.Orders{}

	id := c.Param("id")
	err := or.db.Where("id = ?", id).Model(&model.Orders{}).Preload("Items").First(&order).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return order, nil
}

func (or *OrderRepo) CreateOrder(c *gin.Context) ([]model.Orders, error) {
	var order = []model.Orders{}
	var GetOrder model.Orders
	// var getItem model.Items

	// menggunakan JSON
	JsonOrder := GetOrder
	err := c.BindJSON(&JsonOrder)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	// menggunakan PostForm

	// customer_name := c.PostForm("customer_name")
	// ordered_at := c.PostForm("ordered_at")
	// item_code := c.PostForm("item_code")
	// description := c.PostForm("description")
	// quantity := c.PostForm("quantity")

	// GetOrder.Customer_Name = JsonOrder.Customer_Name
	// GetOrder.Ordered_At = JsonOrder.Ordered_At
	// getItem.Item_Code = item_code
	// getItem.Description = description
	// getItem.Quantity = quantity

	err = or.db.Create(&model.Orders{
		Customer_Name: JsonOrder.Customer_Name,
		Ordered_At:    JsonOrder.Ordered_At,
		Items:         JsonOrder.Items,
	}).Error
	fmt.Println(JsonOrder)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return order, nil
}

func (or *OrderRepo) UpdateOrder(c *gin.Context) ([]model.Orders, error) {
	var order = []model.Orders{}
	var GetOrder model.Orders
	// var newOrder model.Orders
	// var item model.Items

	// menggunakan JSON
	JsonOrder := GetOrder
	err := c.BindJSON(&JsonOrder)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}

	id := c.Param("id")

	// menggunakan Form data

	// customer_name := c.PostForm("customer_name")
	// ordered_at := c.PostForm("ordered_at")
	// item_code := c.PostForm("item_code")
	// description := c.PostForm("description")
	// quantity := c.PostForm("quantity")

	err = or.db.First(&GetOrder, id).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	// newOrder.Customer_Name = customer_name
	// newOrder.Ordered_At = ordered_at
	// item.Item_Code = item_code
	// item.Description = description
	// item.Quantity = quantity

	err = or.db.Model(&GetOrder).Updates(&model.Orders{
		Customer_Name: JsonOrder.Customer_Name,
		Ordered_At:    JsonOrder.Ordered_At,
		Items:         JsonOrder.Items,
	}).Error

	return order, nil
}

func (or *OrderRepo) DeleteOrder(c *gin.Context) ([]model.Orders, error) {
	var order = []model.Orders{}
	var getOrder model.Orders

	id := c.Param("id")
	err := or.db.First(&getOrder, id).Error
	if err != nil {
		fmt.Println("data tidak ditemukan!")
		return nil, err
	}

	// err2 := idb.DB.Where("order_id = ?", id).First(&item).Error
	// if err2 != nil {
	// 	result = gin.H{
	// 		"result": "Data Items Not Found",
	// 	}
	// }

	// err = or.db.Unscoped().Delete(&getOrder).Error // permanent delete with unscoped
	err = or.db.Exec(`
	DELETE orders,items 
	FROM orders 
	INNER JOIN items 
	WHERE orders.id = items.order_id 
	AND items.order_id = ?`, id).Error
	if err != nil {
		fmt.Println("Data Gagal Dihapus!")
		return nil, err
	}
	// err2 = idb.DB.Unscoped().Delete(&item).Error // permanent delete with unscoped
	return order, nil
}
