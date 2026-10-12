package t_007_26_remove_duplicates_from_sorted_array

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var (
		nums     = []int{1, 1, 2}
		expected = 2
		actual   int
	)
	actual = removeDuplicates(nums)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
