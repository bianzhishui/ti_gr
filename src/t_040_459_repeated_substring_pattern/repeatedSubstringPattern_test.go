package t_040_459_repeated_substring_pattern

import (
	"testing"
)

func TestRepeatedSubstringPattern(t *testing.T) {
	var (
		s        = "abab"
		expected = true
		actual   bool
	)
	actual = repeatedSubstringPattern(s)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestRepeatedSubstringPattern1(t *testing.T) {
	var (
		s        = "abab"
		expected = true
		actual   bool
	)
	actual = repeatedSubstringPattern1(s)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
