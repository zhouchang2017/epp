package inventories

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/modules/supplies/services"
	"net/http"
)

func CreateSupplyOrder(c *gin.Context) {

	var createRequest *services.CreateSupplyOrderRequest

	if c.ShouldBind(&createRequest) == nil {
		// TODO
	}
	// Fake
	createRequest.SupplierId = 1
	createRequest.SupplierName = "大名鼎鼎供应商"
	supplyImp := services.MakeSupplyImp()
	order, err := supplyImp.CreateSupplyOrder(createRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, order)
}
