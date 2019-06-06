package models

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)


// 供货单实体
type SupplyOrder struct {
	ID            uint               `gorm:"primary_key"json:"id"`
	Code          uuid.UUID          `json:"code"`                                                   // 供货单编号
	Description   string             `json:"description" gorm:"type:text" form:"description"`        // 供货描述
	WarehouseId   *uint              `json:"warehouse_id" form:"warehouse_id"`                       // 仓库id
	WarehouseName *string            `json:"warehouse_name" form:"warehouse_name"`                   // 仓库名称
	Quantity      int64              `json:"quantity"`                                               // 供货总数量
	Status        int32              `json:"status" form:"status"`                                   // 供货单状态
	SupplierId    *uint              `json:"supplier_id" gorm:"default:null" form:"supplier_id"`     // 供应商id
	SupplierName  *string            `json:"supplier_name" gorm:"default:null" form:"supplier_name"` // 供应商名称
	Transport     int32              `json:"transport" gorm:"defalut:0"`                             // 运输方式
	Items         []*SupplyOrderItem `json:"items,omietempty" form:"items"`
	Shipments     []*Shipment        `json:"shipments,omietempty" gorm :"many2many:supply_shipments;"` // 物流
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	DeletedAt     *time.Time         `sql:"index"json:"deleted_at"`
}

func (this *SupplyOrder) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("Code", uuid.Must(uuid.NewV4()))
	return nil
}
