package t_090_77_combinations

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	var (
		n        = 4
		k        = 2
		expected = [][]int{
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
		}
		actual [][]int
	)
	actual = combine(n, k)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
