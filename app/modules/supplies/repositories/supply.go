package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/zhouchang2017/epp/app/modules/supplies/models"
	"github.com/zhouchang2017/epp/infrastructure"
)

// 供货数据操作层
type Supply struct {
	db *gorm.DB
}

// 创建供货单
func (this *Supply) Create(supplyOrder *models.SupplyOrder) (*models.SupplyOrder, error) {
	err := this.db.Create(supplyOrder).Error

	this.UpdateTotalQuantity(supplyOrder)

	return supplyOrder, err
}

// 更新供货单
func (this *Supply) Update(supplyOrder *models.SupplyOrder) (err error) {
	return this.db.Save(supplyOrder).Error
}

// 获取供货单 with=items
func (this *Supply) Get(id uint) (supplyOrder *models.SupplyOrder, err error) {
	supplyOrder = &models.SupplyOrder{}
	err = this.db.First(supplyOrder, id).Error
	return supplyOrder, err
}

// 更新订货单冗余总商品数量
func (this *Supply) UpdateTotalQuantity(supplyOrder *models.SupplyOrder) (err error) {
	var total int64
	if total, err = this.sumItemsQuantity(supplyOrder.ID); err != nil {
		return err
	}

	if supplyOrder.Quantity != int64(total) {
		supplyOrder.Quantity = total

		this.db.Save(supplyOrder)
	}

	return nil
}

// 供货单明细列表供货数量求和
func (this *Supply) sumItemsQuantity(id uint) (total int64, err error) {

	rows, err := this.db.Model(models.SupplyOrderItem{}).Select("sum(quantity)").Where("supply_order_id = ?", id).Rows()

	if err != nil {
		return total, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&total); err != nil {
			return total, err
		}

	}

	return total, nil
}

// 工厂方法
func MakeSupplyRepository() *Supply {
	return &Supply{
		db: infrastructure.GetDB(),
	}
}
