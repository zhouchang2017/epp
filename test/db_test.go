package test

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhouchang2017/epp/app/modules/supplies/models"
	_ "github.com/zhouchang2017/epp/config"
	"github.com/zhouchang2017/epp/infrastructure"
	"testing"
)

func TestSumQuantity(t *testing.T) {
	infrastructure.Init()
	var id = 8
	quantity := 0
	rows, err := infrastructure.GetDB().Model(models.SupplyOrderItem{}).Select("sum(quantity)").Where("supply_order_id = ?", id).Rows()

	if err != nil {

	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&quantity); err != nil {
			panic(err)
		}

	}

	fmt.Println(quantity)
}

func TestDBQueryRelations(t *testing.T) {
	infrastructure.Init()
	var id = 1
	db:=infrastructure.GetDB()
	data:=&models.SupplyOrder{}
	db.Preload("Items", func(db *gorm.DB) *gorm.DB {
		return db.Limit(1)
	}).First(data, id)
	fmt.Println(data)
}
