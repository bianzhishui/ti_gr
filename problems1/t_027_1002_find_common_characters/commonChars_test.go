package t_027_1002_find_common_characters

import (
	"reflect"
	"testing"
)

func TestCommonChars(t *testing.T) {
	var (
		words    = []string{"bella", "label", "roller"}
		expected = []string{"e", "l", "l"}
		actual   []string
	)
	actual = commonChars(words)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
