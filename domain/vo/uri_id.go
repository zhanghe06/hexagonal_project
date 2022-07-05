package vo

type UriIdReq struct {
	ID uint64 `uri:"id" binding:"required"`
}
