package t_100_46_permutations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	var (
		nums     = []int{1, 2, 3}
		expected = [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}
		actual [][]int
	)
	actual = permute(nums)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
