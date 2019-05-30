package repositories

import (
	"github.com/zhouchang2017/epp/app/modules/inventories/models"
	"github.com/zhouchang2017/epp/infrastructure"
)

// 供货数据操作层
type Supply struct {
}

// 创建供货单
// description



func (this *Supply) Create(supplyOrder *models.SupplyOrder) (*models.SupplyOrder, error) {
	err := infrastructure.GetDB().Create(supplyOrder).Error
	return supplyOrder, err
}
