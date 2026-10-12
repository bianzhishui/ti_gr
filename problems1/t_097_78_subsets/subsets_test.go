package t_097_78_subsets

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	var (
		nums     = []int{1, 2, 3}
		expected = [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 3},
			{1, 3},
			{2},
			{2, 3},
			{3},
		}
		actual [][]int
	)
	actual = subsets(nums)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
