package t_003_34_find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	var (
		nums     = []int{5, 7, 7, 8, 8, 10}
		target   = 5
		expected = []int{0, 0}
		actual   []int
	)
	actual = searchRange(nums, target)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}

func TestSearchRange1(t *testing.T) {
	var (
		nums     = []int{5, 7, 7, 8, 8, 10}
		target   = 5
		expected = []int{0, 0}
		actual   []int
	)
	actual = searchRange1(nums, target)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}
