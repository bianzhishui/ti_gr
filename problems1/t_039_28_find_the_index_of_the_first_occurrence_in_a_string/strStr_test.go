package t_039_28_find_the_index_of_the_first_occurrence_in_a_string

import (
	"reflect"
	"testing"
)

func TestStrStr(t *testing.T) {
	var (
		haystack = "sadbutsad"
		needle   = "sad"
		expected = 0
		actual   int
	)
	actual = strStr(haystack, needle)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestGetNext(t *testing.T) {
	var (
		s        = "abababab"
		expected = []int{0, 0, 1, 2, 3, 4, 5, 6}
		actual   []int
	)
	actual = make([]int, len(s))
	getNext(actual, s)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
