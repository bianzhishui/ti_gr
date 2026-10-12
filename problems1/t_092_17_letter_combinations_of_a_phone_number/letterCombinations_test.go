package t_092_17_letter_combinations_of_a_phone_number

import (
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	var (
		digits   = "23"
		expected = []string{
			"ad",
			"ae",
			"af",
			"bd",
			"be",
			"bf",
			"cd",
			"ce",
			"cf",
		}
		actual []string
	)
	actual = letterCombinations1(digits)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
