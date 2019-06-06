package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/zhouchang2017/epp/app/modules/supplies/models"
	"github.com/zhouchang2017/epp/common"
	"github.com/zhouchang2017/epp/infrastructure"
)

const (
	SUPPLY_CANCEL       = iota // 取消
	SUPPLY_UN_COMMIT           // 保存
	SUPPLY_PENDING             // 待审核
	SUPPLY_APPROVED            // 待分配仓库
	SUPPLY_UN_SHIP             // 待发货
	SUPPLY_PART_SHIPPED        // 部分发货
	SUPPLY_SHIPPED             // 待收货
	SUPPLY_COMPLATED           // 已完成
)

const (
	TRANSPORT_NULL = iota // 无需物流
	TRANSPORT_BASE        // 物流运输
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
func (this *Supply) Update(id uint, params map[string]interface{}) (err error) {
	order := &models.SupplyOrder{}

	if err = this.db.First(order, id).Error; err != nil {
		return
	}

	return this.db.Model(order).Updates(params).Error
}

// 获取供货单详情（包含items）
func (this *Supply) Get(id uint) (supplyOrder *models.SupplyOrder, err error) {
	supplyOrder = &models.SupplyOrder{}
	err = this.db.Preload("Items").First(supplyOrder, id).Error
	return supplyOrder, err
}

// 获取供货单列表
func (this *Supply) GetList(request *models.SupplyOrder, page *common.PageRequest) (paginate *common.Paginate, err error) {
	supplyOrderList := []models.SupplyOrder{}
	return infrastructure.Page(this.db, &supplyOrderList, page, request)
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

// 获取供货单物流

// 发货
func (this *Supply) CreateShipment(id uint, shipment *models.Shipment, items []uint) (err error) {
	// 记录物流单
	err = this.db.FirstOrCreate(shipment, shipment).Error
	if err != nil {
		return err
	}
	supplyOrder, err := this.Get(id)
	if err != nil {
		return err
	}

	err = this.db.Model(supplyOrder).Association("Shipments").Append(shipment).Error
	if err != nil {
		return err
	}

	return this.db.Model(&models.SupplyOrderItem{}).Where("id IN ?", items).Update("shipment_id", shipment.ID).Error
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
