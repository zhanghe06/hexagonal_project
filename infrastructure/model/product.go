package model

// Product 产品
type Product struct {
	BaseModel
	Name  string `gorm:"column:name;type:varchar(128);default:'';not null;comment:'产品名称'" json:"name"`  // 产品名称
	Brand string `gorm:"column:brand;type:varchar(64);default:'';not null;comment:'产品品牌'" json:"brand"` // 产品品牌
}

func (m *Product) TableName() string {
	return "product"
}
