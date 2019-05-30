package inventories

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/modules/inventories/models"
	"github.com/zhouchang2017/epp/app/modules/inventories/repositories"
	"net/http"
)

func CreateSupplyOrder(c *gin.Context) {

	var supplyOrder *models.SupplyOrder

	if c.ShouldBind(&supplyOrder) == nil {
		// TODO
	}

	supplyRep := &repositories.Supply{}
	order, err := supplyRep.Create(supplyOrder)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, order)
}
