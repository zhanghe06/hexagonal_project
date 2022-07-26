package model

// SupplierAddress 供方地址（同一个供方支持多个地址）
type SupplierAddress struct {
	BaseModel
	SupplierId uint64   `gorm:"column:supplier_id;type:bigint(21) unsigned;default:0;not null;comment:'供方ID'" json:"supplier_id"` // 供方ID
	Supplier   Supplier `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
	Address    string   `gorm:"column:address;type:varchar(128);default:'';not null;comment:'地址'" json:"address"`      // 地址
	Contact    string   `gorm:"column:contact;type:varchar(32);default:'';not null;comment:'姓名'" json:"contact"`       // 姓名
	Phone      string   `gorm:"column:phone;type:varchar(128);default:'';not null;comment:'电话'" json:"phone"`          // 电话
	Email      string   `gorm:"column:email;type:varchar(128);default:'';not null;comment:'邮箱'" json:"email"`          // 邮箱
	DefaultSt  uint8    `gorm:"column:default_st;type:tinyint(4);default:0;not null;comment:'默认状态'" json:"default_st"` // 默认状态（0:常规,1:默认）
}

func (m *SupplierAddress) TableName() string {
	return "supplier_address"
}
