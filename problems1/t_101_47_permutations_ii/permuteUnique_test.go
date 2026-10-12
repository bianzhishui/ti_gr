package t_101_47_permutations_ii

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	var (
		nums     = []int{1, 1, 2}
		expected = [][]int{
			{1, 1, 2},
			{1, 2, 1},
			{2, 1, 1},
		}
		actual [][]int
	)
	actual = permuteUnique(nums)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
