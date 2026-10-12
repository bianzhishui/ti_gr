package t_004_69_sqrtx

import (
	"testing"
)

func TestMySqrt(t *testing.T) {
	var (
		x        = 0
		expected = 0
		actual   int
	)
	actual = mySqrt(x)
	if expected == actual {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}

func TestMySqrt1(t *testing.T) {
	var (
		x        = 0
		expected = 0
		actual   int
	)
	actual = mySqrt1(x)
	if expected == actual {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}
