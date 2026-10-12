package t_026_438_find_all_anagrams_in_a_string

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {
	var (
		s        = "abab"
		p        = "ab"
		expected = []int{0, 1, 2}
		actual   []int
	)
	actual = findAnagrams(s, p)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
