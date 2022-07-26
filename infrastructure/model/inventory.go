package model

// Inventory 存货
type Inventory struct {
	BaseModel
	ProductId uint64  `gorm:"column:product_id;type:bigint(21) unsigned;default:0;not null;comment:'产品ID'" json:"product_id"` // 产品ID
	Product   Product `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	RackId    uint64  `gorm:"column:rack_id;type:bigint(21) unsigned;default:0;not null;comment:'货架ID'" json:"rack_id"` // 货架ID
	Rack      Rack    `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	StockQty  uint32  `gorm:"column:stock_qty;int(11) unsigned;default:0;not null;comment:'存货数量'" json:"stock_qty"` // 存货数量
}

func (m *Inventory) TableName() string {
	return "inventory"
}
