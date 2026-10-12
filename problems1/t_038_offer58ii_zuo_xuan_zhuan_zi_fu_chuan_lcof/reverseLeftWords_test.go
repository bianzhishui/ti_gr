package t_038_offer58ii_zuo_xuan_zhuan_zi_fu_chuan_lcof

import (
	"testing"
)

func TestReverseLeftWords(t *testing.T) {
	var (
		s        = "abcdefg"
		k        = 2
		expected = "cdefgab"
		actual   string
	)
	actual = reverseLeftWords(s, k)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestReverseLeftWords1(t *testing.T) {
	var (
		s        = "abcdefg"
		k        = 2
		expected = "cdefgab"
		actual   string
	)
	actual = reverseLeftWords1(s, k)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
