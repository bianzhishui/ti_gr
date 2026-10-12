package t_037_151_reverse_words_in_a_string

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	var (
		s        = "the sky is blue"
		expected = "blue is sky the"
		actual   string
	)
	actual = reverseWords(s)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestReverseWords1(t *testing.T) {
	var (
		s        = "the sky is blue"
		expected = "blue is sky the"
		actual   string
	)
	actual = reverseWords1(s)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
