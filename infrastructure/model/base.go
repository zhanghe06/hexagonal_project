package model

import "gorm.io/plugin/soft_delete"

// BaseModel 基础模型
type BaseModel struct {
	Id        uint64                `gorm:"column:id;type:bigint(21) unsigned auto_increment;primary_key" json:"id"`
	CreatedAt uint64                `gorm:"column:created_at;type:bigint(21) unsigned;not null;comment:'创建时间';auto_create_time" json:"created_at"`                  // 创建时间
	UpdatedAt uint64                `gorm:"column:updated_at;type:bigint(21) unsigned;not null;comment:'更新时间';auto_create_time;auto_update_time" json:"updated_at"` // 更新时间
	DeletedAt soft_delete.DeletedAt `gorm:"column:deleted_at;type:bigint(21) unsigned;default:0;not null;comment:'删除时间'" json:"deleted_at"`                         // 删除时间
	CreatedBy uint64                `gorm:"column:created_by;type:bigint(21) unsigned;default:0;not null;comment:'创建人员'" json:"created_by"`                         // 创建人员
	UpdatedBy uint64                `gorm:"column:updated_by;type:bigint(21) unsigned;default:0;not null;comment:'更新人员'" json:"updated_by"`                         // 更新人员
	DeletedBy uint64                `gorm:"column:deleted_by;type:bigint(21) unsigned;default:0;not null;comment:'删除人员'" json:"deleted_by"`                         // 删除人员
}

// 自增主键设置：AUTO_INCREMENT 是 type 的属性；设置错误会出现：Field 'id' doesn't have a default value
// 定义基础结构体的时候，CreatedAt 和 UpdatedAt 不能包含标签default:0，否则通过结构体创建数据的时候，CreatedAt和UpdatedAt 会一直为0，自动补充时间戳(秒)的时候，不会生效
