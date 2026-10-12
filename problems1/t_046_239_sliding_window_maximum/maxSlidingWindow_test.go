package t_046_239_sliding_window_maximum

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	var (
		nums     = []int{1, 3, -1, -3, 5, 3, 6, 7}
		k        = 3
		expected = []int{3, 3, 5, 5, 6, 7}
		actual   []int
	)
	actual = maxSlidingWindow(nums, k)
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
