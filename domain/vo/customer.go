package vo

// CustomerCreateReq .
type CustomerCreateReq struct {
	Name string `binding:"required,email" form:"name" json:"name"` // 名称
}

// CustomerUpdateReq .
type CustomerUpdateReq struct {
	Name string `binding:"omitempty" json:"name,omitempty"` // 名称
}

// CustomerGetListReq .
type CustomerGetListReq struct {
	Limit     *int     `binding:"omitempty" form:"limit,omitempty" json:"limit,omitempty"`
	Offset    *int     `binding:"omitempty" form:"offset,omitempty" json:"offset,omitempty"`
	CreatedAt []string `binding:"omitempty" form:"created_at[],omitempty" json:"created_at[],omitempty"` // 创建时间
	Sorter    string   `binding:"omitempty" form:"sorter,omitempty" json:"sorter,omitempty"`             // 排序字段
}
