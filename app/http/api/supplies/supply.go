package supplies

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhouchang2017/epp/app/modules/supplies/services"
	"github.com/zhouchang2017/epp/common"
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
		service := services.MakeSupplyImp()
		supplyOrder, err := service.Get(uint(i))

		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, supplyOrder)
	}

}

func Index(c *gin.Context) {
	getSupplyOrderListRequest := &services.GetSupplyOrderListRequest{&services.SupplyFilterRequest{}, &common.PageRequest{}}

	if err := c.ShouldBindQuery(&getSupplyOrderListRequest); err == nil {
		fmt.Printf("%+v\n", getSupplyOrderListRequest)
		service := services.MakeSupplyImp()
		supplyFilterRequest := &services.SupplyFilterRequest{}
		{
			supplyFilterRequest.ID = getSupplyOrderListRequest.ID
			supplyFilterRequest.SupplierId = getSupplyOrderListRequest.SupplierId
			supplyFilterRequest.Transport = getSupplyOrderListRequest.Transport
			supplyFilterRequest.WarehouseId = getSupplyOrderListRequest.WarehouseId
			supplyFilterRequest.Status = getSupplyOrderListRequest.Status
			supplyFilterRequest.Code = getSupplyOrderListRequest.Code
		}

		page := &common.PageRequest{}
		{
			page.Page = getSupplyOrderListRequest.Page
			page.PerPage = getSupplyOrderListRequest.PerPage
		}
		res, err := service.GetSupplyOrderList(supplyFilterRequest, page)

		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, res)

	} else {
		log.Printf("req: %+v, error: %s", getSupplyOrderListRequest, err.Error())
		c.JSON(http.StatusInternalServerError, err)
	}

}

func Approve(c *gin.Context) {
	id := c.Param("id")
	status := c.PostForm("status")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	if s, err := strconv.Atoi(status); err != nil {
		log.Println(err)
	} else {
		service := services.MakeSupplyImp()
		if err := service.ApproveSupplyOrder(uint(i), int32(s)); err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}

}

func Shipment(c *gin.Context) {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}

	var shipmentRequest services.ShipmentRequest
	if c.ShouldBind(&shipmentRequest) == nil {
		service := services.MakeSupplyImp()
		err := service.Shipment(uint(i), shipmentRequest)
		if err != nil {
			log.Println(err)
			return
		}
		c.JSON(http.StatusNoContent, nil)
	}
}
