package supplies

import (
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/modules/supplies/repositories"
	"log"
	"net/http"
	"strconv"
)

func Show(c *gin.Context) {
	id := c.Param("id")
	log.Println("请求id:", id)
	if id != "" {
		i, e := strconv.Atoi(id)
		if e != nil {
			log.Println(e)
		}
		service := repositories.MakeSupplyRepository()
		supplyOrder, err := service.Get(uint(i))

		if err != nil {
			log.Println(err)
		}

		c.JSON(http.StatusOK, supplyOrder)

	}

}
