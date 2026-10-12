package t_034_344_reverse_string

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	var (
		actual   = []byte{'h', 'e', 'l', 'l', 'o'}
		expected = []byte{'o', 'l', 'l', 'e', 'h'}
	)
	reverseString(actual)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
