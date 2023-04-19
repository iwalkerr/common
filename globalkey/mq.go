package globalkey

const (
	// 关闭过期的商品
	ScheduleCloseUpExpireGoods = "schedule:expire_goods:close"
	// 删除过期的订单
	ScheduleDeleteExpireOrder = "schedule:expire_order:delete"
	// 商家的定时结算记录
	// ScheduleRecordBusinessSettle = "schedule:business_settle:record"
	// 延迟关闭订单
	DeferFinishGoodsOrder = "defer:goods_order:finish"
)

// 任务队列优先级名称
const (
	Critical = "critical"
	Default  = "default"
	Low      = "low"
)

type MsgPayload struct {
	Key   string
	Value string
}

// 关闭延时订单
type DeferCloseGoodsOrderPayload struct {
	Sn string `json:"sn"`
}
