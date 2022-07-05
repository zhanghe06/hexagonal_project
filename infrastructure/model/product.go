package model

// Product 商品
type Product struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(128);default:'';NOT NULL" json:"name"` // 内容存储库
}

func (m *Product) TableName() string {
	return "product"
}
