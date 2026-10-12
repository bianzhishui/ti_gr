package t_005_367_valid_perfect_square

import (
	"testing"
)

func TestIsPerfectSquare(t *testing.T) {
	var (
		num      = 8
		expected = false
		actual   bool
	)
	actual = isPerfectSquare(num)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestIsPerfectSquare1(t *testing.T) {
	var (
		num      = 8
		expected = false
		actual   bool
	)
	actual = isPerfectSquare1(num)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
