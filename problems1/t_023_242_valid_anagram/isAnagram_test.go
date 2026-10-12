package t_023_242_valid_anagram

import (
	"testing"
)

func TestIsAnagram(t *testing.T) {
	var (
		s        = "anagram"
		t1       = "nagaram"
		expected = true
		actual   bool
	)
	actual = isAnagram(s, t1)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
