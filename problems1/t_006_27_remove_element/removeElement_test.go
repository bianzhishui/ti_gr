package t_006_27_remove_element

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	var (
		nums     = []int{0, 1, 2, 2, 3, 0, 4, 2}
		val      = 2
		expected = 5
		actual   int
	)
	actual = removeElement(nums, val)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestRemoveElement1(t *testing.T) {
	var (
		nums     = []int{0, 1, 2, 2, 3, 0, 4, 2}
		val      = 2
		expected = 5
		actual   int
	)
	actual = removeElement1(nums, val)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestRemoveElement2(t *testing.T) {
	var (
		nums     = []int{0, 1, 2, 2, 3, 0, 4, 2}
		val      = 2
		expected = 5
		actual   int
	)
	actual = removeElement2(nums, val)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
