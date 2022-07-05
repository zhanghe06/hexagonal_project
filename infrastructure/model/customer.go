package model

// Customer 客户
type Customer struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(128);default:'';NOT NULL" json:"name"`
}

func (m *Customer) TableName() string {
	return "customer"
}
