package globalkey

const (
	OrderAll      = -2 // 获取所有订单
	OrderDelete   = -1 // 订单删除
	OrderWaitPay  = 0  // 等待支付
	OrderClose    = 1  // 订单取消
	OrderSuccess  = 2  // 订单支付成功
	OrderWaitSend = 3  // 等待配送
	OrderDone     = 4  // 订单完成 -- 自提与配送到家
	OrderInvite   = 5  // 订单自提
	OrderDownPay  = 6  // 订单预售
	OrderWaitDeal = 7  // 等待处理
)

const (
	GoodsDown       = 0  // 商品下架
	GoodsEnd        = 10 // 商品结束
	GoodsDownPayEnd = 20 // 商品预购结束
	GoodsUp         = 30 // 商品上架
	GoodsDownPay    = 40 // 商品预购

	GoodsPageNum int64 = 20 // 商品每次获取数据条数
)

const (
	UserAuthTypeSystem  = "system" // 平台内部
	UserAuthTypeSmallWX = "wxMini" // 微信小程序
)

const (
	GoodsHomeDelivery = 1 // 送货上门
	GoodsStoreUp      = 2 // 自提
)

const (
	DeleteOrderTimeSeconds int64 = 60 * 60 * 24 * 15 // 15天 删除订单过期时间
	CloseOrderTimeSeconds        = 60 * 10           // 10分钟订单过期
)
