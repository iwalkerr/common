package toolx

import (
	"fmt"
	"testing"
)

func TestFullName(t *testing.T) {
	for i := 0; i < 100; i++ {
		fmt.Println(GetFullName())
	}
}
