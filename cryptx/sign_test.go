package cryptx

import (
	"fmt"
	"testing"
)

func TestSign(t *testing.T) {
	mapd := make(map[string]string, 0)

	// ct := time.Now().UnixMilli()

	mapd["Salt"] = "4g89aojga#hhjj666777744@@ggjjj"
	mapd["Ct"] = "1667448273129"

	// mapd["sn"] = "202314022128580125114891"
	mapd["buy_time"] = "0"
	mapd["status"] = "-2"
	// mapd["phone"] = "13881887711"
	// mapd["floor"] = "13-6-51"
	// mapd["community"] = "玺龙湾"
	// mapd["avatar_url"] = "https://thirdwx.qlogo.cn/mmopen/vi_32/CfkXZvicPCLR6ELo8CSt1afRia54vperaoBTVY9PrKELeh3UI97AkJVYCxZajMdiaL5ETUlxnDkNvnERwQVy8YnNA/132"
	// mapd["gender"] = "0"

	host := "api.zs2020516.com/v1/order/list"

	sign := GetSign(mapd, host)
	fmt.Println(sign)

	// ss := cryptx.MD5([]byte("Ct1667741213950Salt4g89aojga$#hhjj666777744@@ggjjjcategory_id1cursor0api.zs2020516.com/v1/product/category/goods"))
	// fmt.Println(ss)
	// m := mapd[100]
	// fmt.Println(m)
}
