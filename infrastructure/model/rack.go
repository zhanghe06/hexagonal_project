package model

// Rack 货架
type Rack struct {
	BaseModel
	Name        string    `gorm:"column:name;type:varchar(128);default:'';not null;comment:'货架名称'" json:"name"`
	WarehouseId uint64    `gorm:"column:warehouse_id;type:bigint(21) unsigned;default:0;not null;comment:'仓库ID'" json:"warehouse_id"` // 仓库ID
	Warehouse   Warehouse `gorm:"foreign_key:warehouse_id"`
}

func (m *Rack) TableName() string {
	return "rack"
}
