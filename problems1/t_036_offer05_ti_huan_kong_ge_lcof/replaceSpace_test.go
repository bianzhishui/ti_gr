package t_036_offer05_ti_huan_kong_ge_lcof

import (
	"testing"
)

func TestReplaceSpace(t *testing.T) {
	var (
		s        = "We are happy."
		expected = "We%20are%20happy."
		actual   string
	)
	actual = replaceSpace(s)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
