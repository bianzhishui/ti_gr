package t_028_349_intersection_of_two_arrays

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	var (
		nums1    = []int{1, 2, 2, 1}
		nums2    = []int{2, 2}
		expected = []int{2}
		actual   []int
	)
	actual = intersection(nums1, nums2)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
