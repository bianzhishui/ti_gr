package t_002_35_search_insert_position

import (
	"testing"
)

func TestSearchInsert(t *testing.T) {
	var (
		nums     = []int{-1, 3, 5, 7}
		target   = 0
		expected = 1
		actual   int
	)
	actual = searchInsert(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}

func TestSearchInsert1(t *testing.T) {
	var (
		nums     = []int{-1, 3, 5, 7}
		target   = 0
		expected = 1
		actual   int
	)
	actual = searchInsert1(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}

func TestSearchInsert2(t *testing.T) {
	var (
		nums     = []int{-1, 3, 5, 7}
		target   = 0
		expected = 1
		actual   int
	)
	actual = searchInsert2(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}
