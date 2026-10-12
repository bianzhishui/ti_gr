package t_095_131_palindrome_partitioning

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	var (
		s        = "aab"
		expected = [][]string{
			{"a", "a", "b"},
			{"aa", "b"},
		}
		actual [][]string
	)
	actual = partition(s)
	logstr := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logstr)
	} else {
		t.Fatal(logstr)
	}
}
