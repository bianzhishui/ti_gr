package t_013_76_minimum_window_substring

import (
	"testing"
)

func TestMinWindow(t *testing.T) {
	var (
		s        = "ADOBECODEBANC"
		t2       = "ABC"
		expected = "BANC"
		actual   string
	)
	actual = minWindow(s, t2)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
