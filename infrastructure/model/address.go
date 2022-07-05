package model

// Address 地址
type Address struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(128);default:'';NOT NULL" json:"name"` // 内容存储库
}

func (m *Address) TableName() string {
	return "address"
}
