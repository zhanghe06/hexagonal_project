package model

// User 用户
type User struct {
	BaseModel
	Name   string `gorm:"column:name;type:varchar(128);default:'';not null;comment:'用户名称'" json:"name"`
	RoleId uint64 `gorm:"column:role_id;type:bigint(21) unsigned;default:0;not null;comment:'角色ID'" json:"role_id"` // 角色ID
	Role   Role   `gorm:"constraint:on_update:cascade,on_delete:cascade;"`
}

func (m *User) TableName() string {
	return "user"
}
