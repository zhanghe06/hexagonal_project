package model

// StockCount 盘点（仅支持同一个仓库不同货架的迁移）
type StockCount struct {
	BaseModel
	WarehouseId     uint64           `gorm:"column:warehouse_id;type:bigint(21) unsigned;default:0;not null;comment:'仓库ID'" json:"warehouse_id"` // 仓库ID
	Warehouse       Warehouse        `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	UserId          uint64           `gorm:"column:user_id;type:bigint(21) unsigned;default:0;not null;comment:'员工ID'" json:"user_id"` // 员工ID
	User            User             `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	StockCountSt    uint8            `gorm:"column:stock_count_st;type:tinyint(4);default:0;not null;comment:'盘点状态'" json:"stock_count_st"` // 盘点状态（0:已停用,1:已启用）
	StockCountItems []StockCountItem `gorm:"foreign_key:stock_count_id"`                                                                    // Has Many
}

func (m *StockCount) TableName() string {
	return "stock_count"
}

// StockCountItem 盘点明细
type StockCountItem struct {
	BaseModel
	StockCountId uint64    `gorm:"column:stock_count_id;type:bigint(21) unsigned;default:0;not null;comment:'盘点ID'" json:"stock_count_id"` // 盘点ID
	InventoryId  uint64    `gorm:"column:inventory_id;type:bigint(21) unsigned;default:0;not null;comment:'存货ID'" json:"inventory_id"`     // 存货ID
	Inventory    Inventory `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	RackId       uint64    `gorm:"column:rack_id;type:bigint(21) unsigned;default:0;not null;comment:'货架ID'" json:"rack_id"` // 货架ID
	Rack         Rack      `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	InventorySt  uint8     `gorm:"column:inventory_st;type:tinyint(4);default:0;not null;comment:'存货状态'" json:"inventory_st"` // 存货状态（1:盘盈,2:盘亏）
	Quantity     uint32    `gorm:"column:quantity;type:int(11) unsigned;default:0;not null;comment:'数量'" json:"quantity"`     // 数量
}

func (m *StockCountItem) TableName() string {
	return "stock_count_item"
}

// TODO 校验数量
