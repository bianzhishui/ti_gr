package t_032_15_3sum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	var (
		nums     = []int{-1, 0, 1, 2, -1, -4}
		expected = [][]int{{-1, -1, 2}, {-1, 0, 1}}
		actual   [][]int
	)

	actual = threeSum(nums)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
