package t_098_90_subsets_ii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	var (
		nums     = []int{1, 2, 2}
		expected = [][]int{
			{},
			{1},
			{1, 2},
			{1, 2, 2},
			{2},
			{2, 2},
		}
		actual [][]int
	)
	actual = subsetsWithDup(nums)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
