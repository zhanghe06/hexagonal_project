package model

// Specification 规格
type Specification struct {
	BaseModel
	ProductId uint64  `gorm:"column:product_id;type:bigint(21) unsigned;default:0;not null;comment:'产品ID'" json:"product_id"` // 产品ID
	Product   Product `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Name      string  `gorm:"column:name;type:varchar(128);default:'';not null;comment:'规格名称'" json:"name"` // 规格名称
}

func (m *Specification) TableName() string {
	return "specification"
}
