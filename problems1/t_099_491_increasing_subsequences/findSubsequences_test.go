package t_099_491_increasing_subsequences

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFindSubsequences(t *testing.T) {
	var (
		nums     = []int{4, 6, 7, 7}
		expected = [][]int{
			{4, 6},
			{4, 6, 7},
			{4, 6, 7, 7},
			{4, 7},
			{4, 7, 7},
			{6, 7},
			{6, 7, 7},
			{7, 7},
		}
		actual [][]int
	)
	actual = findSubsequences(nums)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
