package model

// Role 角色
type Role struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(128);default:'';not null;comment:'角色名称'" json:"name"`
}

func (m *Role) TableName() string {
	return "role"
}
