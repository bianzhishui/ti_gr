package t_033_18_4sum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	var (
		nums     = []int{1, 0, -1, 0, -2, 2}
		target   = 0
		expected = [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}
		actual   [][]int
	)

	actual = fourSum(nums, target)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
