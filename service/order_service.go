package service

import (
	"assignment2_golang/model"
	"assignment2_golang/repositories"

	"github.com/gin-gonic/gin"
)

type serviceImpl struct {
	rr repositories.OrderRepoApi
}

func NewOrderService(rr repositories.OrderRepoApi) *serviceImpl { //provie service
	return &serviceImpl{rr: rr}
}

type OrderServiceApi interface {
	GetOrderService(c *gin.Context) gin.H
	GetOrderByIdService(c *gin.Context) gin.H
	CreateOrderService(c *gin.Context) gin.H
	UpdateOrderService(c *gin.Context) gin.H
	DeleteOrderService(c *gin.Context) gin.H
}

func (s serviceImpl) GetOrderService(c *gin.Context) gin.H {
	var (
		orders []model.Orders
		result gin.H
	)
	orders, err := s.rr.GetOrder(c)
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
	return result
}

func (s serviceImpl) GetOrderByIdService(c *gin.Context) gin.H {
	var (
		order  []model.Orders
		result gin.H
	)
	order, err := s.rr.GetOrderById(c)
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
	return result
}

func (s serviceImpl) CreateOrderService(c *gin.Context) gin.H {
	var (
		order  []model.Orders
		result gin.H
	)

	order, err := s.rr.CreateOrder(c)
	if err != nil {
		result = gin.H{
			"result": order,
		}
	} else {
		result = gin.H{
			"result": "created successfully",
		}
	}
	return result
}

func (s serviceImpl) UpdateOrderService(c *gin.Context) gin.H {
	var (
		orders []model.Orders
		result gin.H
	)
	orders, err := s.rr.UpdateOrder(c)

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
	return result
}

func (s serviceImpl) DeleteOrderService(c *gin.Context) gin.H {
	var (
		orders []model.Orders
		// item   model.Items
		result gin.H
	)
	orders, err := s.rr.DeleteOrder(c)
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
	return result
}
