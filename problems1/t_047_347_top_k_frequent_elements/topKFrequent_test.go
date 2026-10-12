package t_047_347_top_k_frequent_elements

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	var (
		nums     = []int{1, 1, 1, 1, 2, 2, 2, 3}
		k        = 2
		expected = []int{1, 2}
		actual   []int
	)
	actual = topKFrequent(nums, k)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
