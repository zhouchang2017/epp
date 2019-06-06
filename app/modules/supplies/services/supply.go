package services

import (
	"github.com/satori/go.uuid"
	"github.com/zhouchang2017/epp/app/modules/supplies/models"
	"github.com/zhouchang2017/epp/app/modules/supplies/repositories"
	"github.com/zhouchang2017/epp/common"
)

type SupplyImp struct {
	repository *repositories.Supply
}

// 发货
func (this *SupplyImp) Shipment(id uint, req ShipmentRequest) (err error) {
	for _, r := range req {
		this.repository.CreateShipment(id, r.Shipment, r.Items)
	}
	return nil
}

// 管理员审核供货单
func (this *SupplyImp) ApproveSupplyOrder(id uint, status int32) (err error) {
	return this.repository.Update(id, map[string]interface{}{"status": status})
}

// 创建供货单
func (this *SupplyImp) CreateSupplyOrder(req *CreateSupplyOrderRequest) (res *CreateSupplyOrderResponse, err error) {
	// decode
	supplyOrder := &models.SupplyOrder{
		Description:  req.Description,
		Status:       req.Status,
		SupplierId:   &req.SupplierId,
		SupplierName: &req.SupplierName,
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

// 获取订单列表
func (this *SupplyImp) GetSupplyOrderList(req *SupplyFilterRequest, page *common.PageRequest) (res *common.Paginate, err error) {
	// decode
	request := &models.SupplyOrder{}
	request.ID = req.ID
	request.Code = uuid.FromStringOrNil(req.Code)
	request.Status = req.Status
	request.WarehouseId = req.WarehouseId
	request.SupplierId = req.SupplierId
	request.Transport = req.Transport

	return this.repository.GetList(request, page)
}

// 获取订单
func (this *SupplyImp) Get(id uint) (supplyOrder *models.SupplyOrder, err error) {
	return this.repository.Get(id)
}

// 工厂方法
func MakeSupplyImp() *SupplyImp {
	return &SupplyImp{
		repository: repositories.MakeSupplyRepository(),
	}
}
