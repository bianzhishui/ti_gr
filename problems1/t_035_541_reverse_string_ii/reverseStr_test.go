package t_035_541_reverse_string_ii

import "testing"

func TestReverseStr(t *testing.T) {
	var (
		s        = "abcdefg"
		k        = 2
		expected = "bacdfeg"
		actual   string
	)
	actual = reverseStr(s, k)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
