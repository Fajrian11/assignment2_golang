package repositories

import (
	"assignment2_golang/helpers"
	"assignment2_golang/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type OrderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return OrderRepo{
		db: db,
	}
}

type OrderRepoApi interface {
	// ORDER
	GetOrder(c *gin.Context) ([]model.Orders, error)
	GetOrderById(c *gin.Context) ([]model.Orders, error)
	CreateOrder(c *gin.Context) ([]model.Orders, error)
	UpdateOrder(c *gin.Context) ([]model.Orders, error)
	DeleteOrder(c *gin.Context) ([]model.Orders, error)
	// USER
	UserRegister(c *gin.Context) (model.User, error)
	UserLogin(c *gin.Context) (error, bool, string)
}

type ConsumingAPIRepo interface {
}

// ORDER

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

var baseURL = "https://hiyaa.site"

// func (or *OrderRepo) DeleteOrder(c *gin.Context) ([]model.Orders, error) {

func ConsumingAPI(c *gin.Context) {
	response, err := http.Get(baseURL + "/data.php?qty=1&apikey=7f8fc96e-de1f-4aab-9c62-3dd1de365e66")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// menampilkan response dari request
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Result
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(len(responseObject.Person))

	// for i := 0; i < len(responseObject.Person); i++ {
	// 	fmt.Println("First Name :", responseObject.Person[i].Firstname)
	// 	fmt.Println("Last Name :", responseObject.Person[i].Lastname)
	// 	fmt.Println("Usernmae :", responseObject.Person[i].Username)
	// 	fmt.Println("Email :", responseObject.Person[i].Email)
	// 	fmt.Println("Phone :", responseObject.Person[i].Phone)
	// 	fmt.Println("UUID :", responseObject.Person[i].Uuid)
	// }
	fmt.Println(responseObject.Person)
}

var (
	appJSON = "application/json"
)

func (or *OrderRepo) UserRegister(c *gin.Context) (model.User, error) {
	ContentType := helpers.GetContentType(c)

	var user = model.User{}
	var GetUser model.User

	JsonUser := GetUser
	if ContentType == appJSON {
		c.ShouldBindJSON(&JsonUser)
	} else {
		c.ShouldBind(&JsonUser)
	}

	// menggunakan JSON
	// err := c.BindJSON(&JsonUser)
	// if err != nil {
	// 	c.AbortWithError(http.StatusBadRequest, err)
	// 	return nil, err
	// }
	err := or.db.Create(&JsonUser).Error
	fmt.Println(JsonUser)
	if err != nil {
		fmt.Println(err.Error())
	}

	return user, nil
}

func (or *OrderRepo) UserLogin(c *gin.Context) (error, bool, string) {
	contentType := helpers.GetContentType(c)
	_ = contentType
	User := model.User{}
	password := ""

	if contentType == appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	// Validate Email
	err := or.db.Debug().Where("email = ?", User.Email).Take(&User).Error
	// Validate Password
	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))
	// Validate Email & Password Jika Berhasil
	token := helpers.GenerateToken(User.ID, User.Email)

	return err, comparePass, token
}
