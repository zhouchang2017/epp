package services

import (
	models2 "github.com/zhouchang2017/epp/app/modules/inventories/models"
	"github.com/zhouchang2017/epp/app/modules/suppliers/models"
	"github.com/zhouchang2017/epp/infrastructure"
)

type GetRequest struct {
	VariantId uint `json:"variant_id" form:"variant_id"`
	PerPage   uint `json:"per_page" form:"per_page"`
	Page      uint `json:"page" form:"page"`
}

type GetResponse struct {
	CurrentPage uint                      `json:"current_page"`
	Data        []*models.SupplierVariant `json:"data"`
	PerPage     uint                      `json:"per_page"`
	Count       uint                      `json:"count"`
}

// 库存管理接口
type Supplier interface {
	// 获取
	Get(*GetRequest) (response *GetResponse, err error)
}

type SupplierService struct {
}

func (this *SupplierService) Get(params *GetRequest) (response interface{}, err error) {


	build:=infrastructure.GetDB()
	var data []*models.SupplierVariant

	if response,err = infrastructure.Page(build,&data,params.Page,params.PerPage,&models2.Inventory{
		VariantId:params.VariantId,
	}); err != nil {
		return response, err
	}

	return response, nil

}
