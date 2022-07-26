package model

// OutBound 出库（仓库看不到价格）
type OutBound struct {
	BaseModel
	Code          string         `gorm:"index:udx_code,unique;column:code;type:varchar(36);default:'';not null;comment:'出库编号'" json:"code"` // 编号
	OrderId       uint64         `gorm:"column:order_id;type:bigint(21) unsigned;default:0;not null;comment:'订单ID'" json:"order_id"`        // 订单ID
	Order         SalesOrder     `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	OutBoundSt    uint8          `gorm:"column:out_bound_st;type:tinyint(4);default:0;not null;comment:'出库状态'" json:"out_bound_st"` // 出库状态（0:已停用,1:已启用）
	OutBoundItems []OutBoundItem `gorm:"foreign_key:out_bound_id"`                                                                  // Has Many
}

func (m *OutBound) TableName() string {
	return "out_bound"
}

// OutBoundItem 出库明细
type OutBoundItem struct {
	BaseModel
	OutBoundId  uint64    `gorm:"column:out_bound_id;type:bigint(21) unsigned;not null;comment:'出库ID'" json:"out_bound_id"` // 出库ID
	OutBound    OutBound  `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	InventoryId uint64    `gorm:"column:inventory_id;type:bigint(21) unsigned;not null;comment:'存货ID'" json:"inventory_id"` // 存货ID
	Inventory   Inventory `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Quantity    uint32    `gorm:"column:quantity;type:int(11) unsigned;not null;comment:'出库数量'" json:"quantity"` // 出库数量
}

func (m *OutBoundItem) TableName() string {
	return "out_bound_item"
}

// TODO 校验出库数量
