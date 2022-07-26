package model

// Tag 标签
type Tag struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(128);default:'';not null;comment:'标签名称'" json:"name"` // 标签名称
}

func (m *Tag) TableName() string {
	return "tag"
}
