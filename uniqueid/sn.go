package uniqueid

import (
	"fmt"
	"fsmall/common/toolx"
	"strings"
	"time"
)

// 生成sn单号
type SnPrefix string

const (
	SN_PREFIX_GOODS_ORDER   SnPrefix = "HSO" //商品订单前缀
	SN_PREFIX_THIRD_PAYMENT SnPrefix = "PMT" //第三方支付流水记录前缀
)

// 生成单号
func GenSn(snPrefix SnPrefix, userId int64) string {
	date := time.Now().Format(`20060201150405.000`)
	date = strings.ReplaceAll(date, ".", "")

	return date + toolx.Krand(2, toolx.KC_RAND_KIND_NUM) + fmt.Sprintf("%d", userId) + toolx.Krand(3, toolx.KC_RAND_KIND_NUM)
}
