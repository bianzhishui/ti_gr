package t_030_1_two_sum

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	var (
		nums     = []int{2, 7, 11, 15}
		target   = 9
		expected = []int{0, 1}
		actual   []int
	)

	actual = twoSum(nums, target)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
