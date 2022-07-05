package model

import "time"

// BaseModel 基础模型
type BaseModel struct {
	Id        uint64    `gorm:"column:id;type:bigint(21) unsigned AUTO_INCREMENT;primary_key" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;default:current_timestamp;NOT NULL;comment:'创建时间'" json:"created_at"`                             // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;default:current_timestamp on update current_timestamp;NOT NULL;comment:'更新时间'" json:"updated_at"` // 更新时间
	CreatedBy string    `gorm:"column:created_by;type:varchar(64);NOT NULL;comment:'创建人员'" json:"created_by"`                                                     // 创建人员
	UpdatedBy string    `gorm:"column:updated_by;type:varchar(64);NOT NULL;comment:'更新人员'" json:"updated_by"`                                                     // 更新人员
	DeletedBy string    `gorm:"column:deleted_by;type:varchar(64);NOT NULL;comment:'删除人员'" json:"deleted_by"`                                                     // 删除人员
	DeletedSt uint8     `gorm:"column:deleted_st;type:tinyint(4);default:0;NOT NULL;comment:'删除状态'" json:"deleted_st"`                                            // 删除状态（0:未删除,1:已删除）
}

// 自增主键设置：AUTO_INCREMENT 是 type 的属性；设置错误会出现：Field 'id' doesn't have a default value
