package t_001_704_binary_search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	var (
		nums     = []int{-1, 0, 3, 5, 9, 12, 18}
		target   = 9
		expected = 4
		actual   int
	)
	actual = search(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}

func TestSearch1(t *testing.T) {
	var (
		nums     = []int{-1, 0, 3, 5, 9, 12, 18}
		target   = 9
		expected = 4
		actual   int
	)
	actual = search1(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}

func TestSearch2(t *testing.T) {
	var (
		nums     = []int{-1, 0, 3, 5, 9, 12, 18}
		target   = 9
		expected = 4
		actual   int
	)
	actual = search2(nums, target)
	if actual == expected {
		t.Logf("expected=%d, actual=%d", expected, actual)
	} else {
		t.Fatalf("expected=%d, actual=%d", expected, actual)
	}
	t.Log("done")
}
