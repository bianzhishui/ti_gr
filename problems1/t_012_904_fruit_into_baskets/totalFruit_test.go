package t_012_904_fruit_into_baskets

import (
	"testing"
)

func TestTotalFruit(t *testing.T) {
	var (
		fruits   = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
		expected = 5
		actual   int
	)
	actual = totalFruit(fruits)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestTotalFruit1(t *testing.T) {
	var (
		fruits   = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
		expected = 5
		actual   int
	)
	actual = totalFruit1(fruits)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
