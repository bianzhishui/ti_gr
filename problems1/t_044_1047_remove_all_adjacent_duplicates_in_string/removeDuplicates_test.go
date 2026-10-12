package t_044_1047_remove_all_adjacent_duplicates_in_string

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var (
		s        = "abbaca"
		expected = "ca"
		actual   string
	)
	actual = removeDuplicates(s)
	if expected != actual {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}

func TestRemoveDuplicates1(t *testing.T) {
	var (
		s        = "abbaca"
		expected = "ca"
		actual   string
	)
	actual = removeDuplicates1(s)
	if expected != actual {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
