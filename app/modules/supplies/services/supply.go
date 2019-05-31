package services

import (
	"github.com/zhouchang2017/epp/app/modules/supplies/models"
	"github.com/zhouchang2017/epp/app/modules/supplies/repositories"
)

type SupplyImp struct {
	repository *repositories.Supply
}

// 管理员审核供货单
func (this *SupplyImp) ApproveSupplyOrder(id uint, status int32) (err error) {
	panic("implement me")
}

func (this *SupplyImp) CreateSupplyOrder(req *CreateSupplyOrderRequest) (res *CreateSupplyOrderResponse, err error) {
	// decode
	supplyOrder := &models.SupplyOrder{
		Description:  req.Description,
		Status:       req.Status,
		SupplierId:   req.SupplierId,
		SupplierName: req.SupplierName,
		Transport:    req.Transport,
	}

	for _, item := range req.Items {
		i := &models.SupplyOrderItem{
			ProductId:   item.ProductId,
			ProductName: item.ProductName,
			VariantId:   item.VariantId,
			VariantName: item.VariantName,
			Quantity:    item.Quantity,
		}
		supplyOrder.Items = append(supplyOrder.Items, i)
	}

	order, err := this.repository.Create(supplyOrder)
	return &CreateSupplyOrderResponse{order}, err
}

func MakeSupplyImp() *SupplyImp {
	return &SupplyImp{
		repository: repositories.MakeSupplyRepository(),
	}
}
