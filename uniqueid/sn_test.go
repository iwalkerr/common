package uniqueid

import (
	"fmt"
	"testing"
)

func TestGenSn(t *testing.T) {
	sn := GenSn(SN_PREFIX_GOODS_ORDER, 3)
	fmt.Println(sn)
}
