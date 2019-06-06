package models

// 物流实体
type Shipment struct {
	ID     uint           `gorm:"primary_key"json:"id"`
	Num    string         `json:"num" form:"num"`   // 物流单号
	Name   string         `json:"name" form:"name"` // 物流名称
	Orders []*SupplyOrder `json:"orders,omietempty" gorm:"many2many:supply_shipments;"`
}
