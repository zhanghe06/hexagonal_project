package entity

type Address struct {
	Entity
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// GetId 重写匿名方法
func (e Address) GetId() (id uint64) {
	return e.ID
}
