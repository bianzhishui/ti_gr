package t_031_454_4sum_ii

import "testing"

func TestFourSumCount(t *testing.T) {
	var (
		nums1    = []int{1, 2}
		nums2    = []int{-2, -1}
		nums3    = []int{-1, 2}
		nums4    = []int{0, 2}
		expected = 2
		actual   int
	)

	actual = fourSumCount(nums1, nums2, nums3, nums4)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
