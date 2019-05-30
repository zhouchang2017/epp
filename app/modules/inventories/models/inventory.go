package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 库存实体
type Inventory struct {
	ID            uint   `gorm:"primary_key"json:"id"`
	ProductId     uint   `json:"product_id"`
	ProductName   string `json:"product_name"`
	VariantId     uint   `json:"variant_id"`
	VariantName   string `json:"variant_name"`
	WarehouseId   uint   `json:"warehouse_id"`
	WarehouseName string `json:"warehouse_name"`
	Quantity      int64  `json:"quantity"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

var inventroies = []Inventory{
	{ID: 1, ProductId: 1, ProductName: "测试产品1", VariantId: 1, VariantName: "测试产品1 黑色", Quantity: 20, WarehouseId: 1, WarehouseName: "深圳仓库"},
	{ID: 2, ProductId: 1, ProductName: "测试产品1", VariantId: 2, VariantName: "测试产品1 白色色", Quantity: 10, WarehouseId: 1, WarehouseName: "深圳仓库"},
	{ID: 3, ProductId: 3, ProductName: "测试产品3", VariantId: 8, VariantName: "测试产品3 黑色", Quantity: 5, WarehouseId: 1, WarehouseName: "深圳仓库"},
}

func Migrate(db *gorm.DB) {
	fmt.Println("db migrate...")
	table(db)

	data(db)
}

func table(db *gorm.DB) {
	db.AutoMigrate(
		&Inventory{},
	)
}

func data(db *gorm.DB) {
	for _, item := range inventroies {
		db.Create(&item)
	}
}
