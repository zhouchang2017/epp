package models

import "time"

// 供货单明细列表
type SupplyOrderItem struct {
	ID          uint   `gorm:"primary_key"json:"id"`
	ProductId   uint   `json:"product_id" form:"product_id"`
	ProductName string `json:"product_name" form:"product_name"`
	VariantId   uint   `json:"variant_id" form:"variant_id"`
	VariantName string `json:"variant_name" form:"variant_name"`

	Quantity      int64        `json:"quantity" form:"quantity"` // 供货数量
	Status        int32        `json:"status"`                   // 供货单状态
	SupplyOrderId uint         `json:"supply_order_id" form:"supply_order_id"`
	SupplyOrder   *SupplyOrder `json:"supply_order" form:"supply_order"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     time.Time    `json:"updated_at"`
	DeletedAt     *time.Time   `sql:"index"json:"deleted_at"`
}
