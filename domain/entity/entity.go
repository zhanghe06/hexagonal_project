package entity

import "time"

// Entity 匿名接口（interface只有方法，没有field）
// 匿名接口的方式不依赖具体实现，可以对任意实现了该接口的类型进行重写。
// 这在写一些公共库时会非常有用，如果你经常看一些库的源码，匿名接口的写法应该会很眼熟。
type Entity interface {
	GetId() (id uint64)
}

type BaseEntity struct {
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	CreatedBy string    `json:"created_by"` // 创建人
	UpdatedBy string    `json:"updated_by"` // 更新人
}
