package services

import "github.com/zhouchang2017/epp/app/modules/supplies/models"

// 供货单服务
type Supply interface {
	// 创建供货单
	CreateSupplyOrder(req *CreateSupplyOrderRequest) (res *CreateSupplyOrderResponse, err error)

	// 管理员审核供货单
	ApproveSupplyOrder(id uint, status int32) (err error)
}

type CreateSupplyOrderRequest struct {
	Description  string `json:"description" form:"description"`
	Status       int32  `json:"status" form:"status"`
	SupplierId   uint   `json:"supplier_id" form:"supplier_id"`
	SupplierName string `json:"supplier_name" form:"supplier_name"`
	Transport    int32  `json:"transport" form:"transport"`
	Items        []struct {
		ProductId   uint   `json:"product_id" form:"product_id"`
		ProductName string `json:"product_name" form:"product_name"`
		VariantId   uint   `json:"variant_id" form:"variant_id"`
		VariantName string `json:"variant_name" form:"variant_name"`
		Quantity    int64  `json:"quantity" form:"quantity"`
	} `json:"items" form:"items"`
}

type CreateSupplyOrderResponse struct {
	*models.SupplyOrder
}

type ApproveSupplyOrderRequest models.SupplyOrder

type ApproveSupplyOrderResponse struct {
}
