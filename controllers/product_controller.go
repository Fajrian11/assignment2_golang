package controllers

import (
	"assignment2_golang/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct { // implementasi Controller
	psa service.ProductServiceApi
}

func NewProductController(psa service.ProductServiceApi) *ProductController {
	return &ProductController{psa: psa}
}

func (pc *ProductController) CreateProductControllers(c *gin.Context) {
	res := pc.psa.CreateProductService(c)
	c.JSON(http.StatusCreated, res)
}

func (pc *ProductController) GetAllProductControllers(c *gin.Context) {
	res := pc.psa.GetAllProductService(c)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) GetProductByIdControllers(c *gin.Context) {
	res := pc.psa.GetProductByIdService(c)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) UpdateProductController(c *gin.Context) {
	res := pc.psa.UpdateProductService(c)
	c.JSON(http.StatusOK, res)
}

func (pc *ProductController) DeleteProductController(c *gin.Context) {
	res := pc.psa.DeleteProductService(c)
	c.JSON(http.StatusOK, res)
}
