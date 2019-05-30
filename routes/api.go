package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/http/api/inventories"
	"github.com/zhouchang2017/epp/app/http/api/suppliers"
)

func ApiRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/suppliers", suppliers.Get)

	r.POST("/inventories", inventories.CreateSupplyOrder)
}
