package t_093_39_combination_sum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var (
		candidates = []int{2, 3, 6, 7}
		target     = 7
		expected   = [][]int{
			{2, 2, 3},
			{7},
		}
		actual [][]int
	)
	actual = combinationSum(candidates, target)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
