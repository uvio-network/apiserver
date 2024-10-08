package runtime

import (
	"testing"
)

func Test_Runtime_Math_MaxUint256(t *testing.T) {
	var max string
	{
		max = MaxUint256().String()
	}

	var exp string
	{
		exp = "115792089237316195423570985008687907853269984665640564039457584007913129639935"
	}

	if max != exp {
		t.Fatalf("expected %#v got %#v", exp, max)
	}
}

func Test_Runtime_Math_MidUint256(t *testing.T) {
	var mid string
	{
		mid = MidUint256().String()
	}

	var exp string
	{
		exp = "57896044618658097711785492504343953926634992332820282019728792003956564819967"
	}

	if mid != exp {
		t.Fatalf("expected %#v got %#v", exp, mid)
	}
}
