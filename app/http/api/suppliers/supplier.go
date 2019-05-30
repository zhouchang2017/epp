package suppliers

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/modules/suppliers/services"
)

func Get(c *gin.Context) {

	s := services.SupplierService{}

	var getRequest services.GetRequest

	if c.ShouldBind(&getRequest) == nil {

	}

	response, _ := s.Get(&getRequest)

	c.JSON(200, response)
}
