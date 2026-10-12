package t_009_283_move_zeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var (
		nums     = []int{0, 1, 0, 3, 12}
		expected = []int{1, 3, 12, 0, 0}
	)
	moveZeroes(nums)
	if reflect.DeepEqual(expected, nums) {
		t.Logf("expected=%v, actual=%v", expected, nums)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, nums)
	}
	t.Log("done")
}

func TestMoveZeroes1(t *testing.T) {
	var (
		nums     = []int{0, 1, 0, 3, 12}
		expected = []int{1, 3, 12, 0, 0}
	)
	moveZeroes1(nums)
	if reflect.DeepEqual(expected, nums) {
		t.Logf("expected=%v, actual=%v", expected, nums)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, nums)
	}
	t.Log("done")
}
