package t_094_40_combination_sum_ii

import (
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	var (
		candidates = []int{10, 1, 2, 7, 6, 1, 5}
		target     = 8
		expected   = [][]int{
			{1, 1, 6},
			{1, 2, 5},
			{1, 7},
			{2, 6},
		}
		actual [][]int
	)
	actual = combinationSum2(candidates, target)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
