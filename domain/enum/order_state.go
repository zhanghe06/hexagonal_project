package enum

type OrderState int

// 订单状态（0:pending准备,1:waiting等待,2:process进行,3:success成功,4:failure失败）
const (
	OrderStatePending = iota
	OrderStateWaiting
	OrderStateProcess
	OrderStateSuccess
	OrderStateFailure
)

var OrderStateMap = map[OrderState]string{
	OrderStatePending: "准备",
	OrderStateWaiting: "等待",
	OrderStateProcess: "进行",
	OrderStateSuccess: "成功",
	OrderStateFailure: "失败",
}

func (t OrderState) DisplayName() string {
	return OrderStateMap[t]
}
