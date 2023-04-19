package toolx

import (
	"fmt"
	"testing"
)

func TestMd5ByString(t *testing.T) {
	s := Md5ByString("AAA")
	t.Log(s)
}

func TestKrand(t *testing.T) {
	number := Krand(5, KC_RAND_KIND_NUM)
	lower := Krand(5, KC_RAND_KIND_LOWER)
	upper := Krand(5, KC_RAND_KIND_UPPER)
	all := Krand(5, KC_RAND_KIND_ALL)

	fmt.Println(number)
	fmt.Println(lower)
	fmt.Println(upper)
	fmt.Println(all)
}
