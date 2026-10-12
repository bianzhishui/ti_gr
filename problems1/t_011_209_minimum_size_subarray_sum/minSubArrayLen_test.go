package t_011_209_minimum_size_subarray_sum

import (
	"testing"
)

func TestMinSubArrayLen(t *testing.T) {
	var (
		nums     = []int{2, 3, 1, 2, 4, 3}
		target   = 7
		expected = 2
		actual   int
	)
	actual = minSubArrayLen(target, nums)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestMinSubArrayLen1(t *testing.T) {
	var (
		nums     = []int{2, 3, 1, 2, 4, 3}
		target   = 7
		expected = 2
		actual   int
	)
	actual = minSubArrayLen1(target, nums)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestMinSubArrayLen2(t *testing.T) {
	var (
		nums     = []int{2, 3, 1, 2, 4, 3}
		target   = 7
		expected = 2
		actual   int
	)
	actual = minSubArrayLen2(target, nums)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
