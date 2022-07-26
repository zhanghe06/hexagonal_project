package model

// Supplier 供应商
type Supplier struct {
	BaseModel
	Name              string            `gorm:"column:name;type:varchar(128);default:'';not null;comment:'供方名称'" json:"name"`
	SupplierAddresses []SupplierAddress `gorm:"foreign_key:supplier_id"` // Has Many
}

func (m *Supplier) TableName() string {
	return "supplier"
}
