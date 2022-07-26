package model

// Warehouse 仓库
type Warehouse struct {
	BaseModel
	Name    string `gorm:"column:name;type:varchar(128);default:'';not null;comment:'仓库名称'" json:"name"`               // 仓库名称
	OwnerId uint64 `gorm:"column:owner_id;type:bigint(21) unsigned;default:0;not null;comment:'仓管ID'" json:"owner_id"` // 仓管ID
	Owner   User   `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
}

func (m *Warehouse) TableName() string {
	return "warehouse"
}
