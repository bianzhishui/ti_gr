package t_091_216_combination_sum_iii

import (
	"reflect"
	"testing"
)

func TestCombinationSum3(t *testing.T) {
	var (
		k        = 3
		n        = 7
		expected = [][]int{
			{1, 2, 4},
		}
		actual [][]int
	)
	actual = combinationSum3(k, n)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
