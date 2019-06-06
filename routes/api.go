package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/http/api/inventories"
	"github.com/zhouchang2017/epp/app/http/api/suppliers"
	"github.com/zhouchang2017/epp/app/http/api/supplies"
)

func ApiRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/suppliers", suppliers.Get)

	r.GET("/supplies", supplies.Index)

	r.GET("/supplies/:id", supplies.Show)

	r.PUT("/supplies/:id/approve", supplies.Approve)

	r.POST("/supplies/:id/shipment", supplies.Shipment)

	r.POST("/inventories", inventories.CreateSupplyOrder)
}
