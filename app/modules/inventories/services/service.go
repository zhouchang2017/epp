package services

import (
	"github.com/zhouchang2017/epp/app/modules/inventories/models"
)

type TakeRequest struct {
	VariantId   uint
	WarehouseId uint
	Quantity    uint
}

type GetRequest struct {
	ProductId   uint `json:"product_id"`
	VariantId   uint `json:"variant_id"`
	WarehouseId uint `json:"warehouse_id"`
	PerPage     uint `json:"per_page"`
	Page        uint `json:"page"`
}

type GetResponse struct {
	CurrentPage uint                `json:"current_page"`
	Data        []*models.Inventory `json:"data"`
	PerPage     uint                `json:"per_page"`
	Count       uint                `json:"count"`
}

/*
	仓库中心服务:
		业务需求:
			供货单入库


*/
/*
供货单入库
流程：
进入创建供货单页面 -> 填写供货说明
				 -> 选择供货变体
                 -> 填写对应变体供货数量
				 -> 保存/提交


*/
// 库存管理接口
type Inventory interface {
	// 入库
	Put([]*models.Inventory) (err error)
	// 出库
	Take([]*TakeRequest) (err error)
	// 调拨

	// 获取
	Get(*GetRequest) (response *GetResponse, err error)
}
