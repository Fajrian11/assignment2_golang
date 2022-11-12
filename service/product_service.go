package service

import (
	"assignment2_golang/repositories"

	"github.com/gin-gonic/gin"
)

type ProductService struct {
	rr repositories.ProductRepoApi
}

func NewProductService(rr repositories.ProductRepoApi) *ProductService { //provie service
	return &ProductService{rr: rr}
}

type ProductServiceApi interface {
	CreateProductService(c *gin.Context) gin.H
	GetAllProductService(c *gin.Context) gin.H
	GetProductByIdService(c *gin.Context) gin.H
	UpdateProductService(c *gin.Context) gin.H
	DeleteProductService(c *gin.Context) gin.H
}

func (ps ProductService) CreateProductService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Product, err := ps.rr.CreateProduct(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Data Has been created",
			"result":  Product,
		}
	}
	return result
}

func (ps ProductService) GetAllProductService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllProduct, err := ps.rr.GetAllProduct(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"result": GetAllProduct,
			"count":  len(GetAllProduct),
		}
	}
	return result
}

func (ps ProductService) GetProductByIdService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	GetAllProduct, err := ps.rr.GetProductById(c)
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": GetAllProduct,
			"count":  0,
		}
	}
	return result
}

func (ps ProductService) UpdateProductService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	Product, err := ps.rr.UpdateProduct(c)
	if err != nil {
		result = gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Data Has been Updated",
			"result":  Product,
		}
	}
	return result
}

func (ps ProductService) DeleteProductService(c *gin.Context) gin.H {
	var (
		result gin.H
	)

	_, err := ps.rr.DeleteProduct(c)
	if err != nil {
		result = gin.H{
			"result":  "Gagal Menghapus Data",
			"message": err.Error(),
		}
	} else {
		result = gin.H{
			"Success": "Data Has been Deleted",
		}
	}
	return result
}
