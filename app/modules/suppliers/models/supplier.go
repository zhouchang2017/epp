package models

import "time"

type SupplierVariant struct {
	ID        uint      `gorm:"primary_key"json:"id"`
	Name      string    `json:"name"`
	VariantId uint      `json:"variant_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (this SupplierVariant) TableName() string {
	return "supplier_variants"
}
