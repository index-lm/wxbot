package hb

import (
	"fmt"
	"testing"
)

func TestC(t *testing.T) {
	var a int64 = 111
	var p = 6
	hb := computeHb(a, p)
	var tt int64 = 0
	for _, v := range hb {
		tt += v
	}

	fmt.Println(hb)
	fmt.Println(tt)
}
