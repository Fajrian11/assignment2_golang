package repositories

import (
	"assignment2_golang/helpers"
	"assignment2_golang/model"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type ProductRepo struct {
	db *gorm.DB
}

func NewProductRepo(db *gorm.DB) ProductRepo {
	return ProductRepo{
		db: db,
	}
}

type ProductRepoApi interface {
	CreateProduct(c *gin.Context) (model.Product, error)
	UpdateProduct(c *gin.Context) (model.Product, error)
	DeleteProduct(c *gin.Context) (model.Product, error)
	GetAllProduct(c *gin.Context) ([]model.Product, error)
	GetProductById(c *gin.Context) ([]model.Product, error)
}

func (pr *ProductRepo) CreateProduct(c *gin.Context) (model.Product, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Product := model.Product{}
	UserID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = UserID

	err := pr.db.Debug().Create(&Product).Error

	return Product, err
}

func (pr *ProductRepo) GetAllProduct(c *gin.Context) ([]model.Product, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := model.Product{}
	UserID := uint(userData["id"].(float64))

	Product.UserID = UserID

	var GetAllProduct = []model.Product{}
	err := pr.db.Model(&model.Product{}).Find(&GetAllProduct).Error

	return GetAllProduct, err
}

func (pr *ProductRepo) GetProductById(c *gin.Context) ([]model.Product, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := model.Product{}
	productId, _ := strconv.Atoi(c.Param("productId"))
	UserID := uint(userData["id"].(float64))

	Product.UserID = UserID
	Product.ID = uint(productId)

	var GetAllProduct = []model.Product{}
	err := pr.db.Where("id = ?", productId).Model(&model.Orders{}).First(&GetAllProduct).Error

	return GetAllProduct, err
}

func (pr *ProductRepo) UpdateProduct(c *gin.Context) (model.Product, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Product := model.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		c.ShouldBindJSON(&Product)
	} else {
		c.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := pr.db.Model(&Product).Where("id = ?", productId).Updates(model.Product{
		Title:       Product.Title,
		Description: Product.Description,
	}).Error

	return Product, err
}

func (pr *ProductRepo) DeleteProduct(c *gin.Context) (model.Product, error) {
	userData := c.MustGet("userData").(jwt.MapClaims)
	Product := model.Product{}

	productId, _ := strconv.Atoi(c.Param("productId"))
	userID := uint(userData["id"].(float64))

	Product.UserID = userID
	Product.ID = uint(productId)

	err := pr.db.Exec(`
	DELETE products 
	FROM products 
	WHERE products.id = ?`, productId).Error

	return Product, err
}
