package t_043_20_valid_parentheses

import (
	"testing"
)

func TestIsValidTestIsValid(t *testing.T) {
	var (
		s        = "([])"
		expected = true
		actual   bool
	)

	actual = isValid(s)
	if expected != actual {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
