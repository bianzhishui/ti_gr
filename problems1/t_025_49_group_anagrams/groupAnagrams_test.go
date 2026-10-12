package t_025_49_group_anagrams

import (
	"reflect"
	"testing"
)

// TODO: 无序的
func TestGroupAnagrams(t *testing.T) {
	var (
		strs     = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		expected = [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
		actual   [][]string
	)
	actual = groupAnagrams(strs)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
