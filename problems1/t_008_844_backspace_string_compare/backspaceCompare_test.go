package t_008_844_backspace_string_compare

import (
	"fmt"
	"testing"
)

func main() {
	s, t := "a##c", "#a#c"
	fmt.Println(backspaceCompare1(s, t))
}

func TestBackspaceCompare(t *testing.T) {
	var (
		s        = "a##c"
		t2       = "#a#c"
		expected = true
		actual   bool
	)

	actual = backspaceCompare(s, t2)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestBackspaceCompare1(t *testing.T) {
	var (
		s        = "a##c"
		t2       = "#a#c"
		expected = true
		actual   bool
	)

	actual = backspaceCompare1(s, t2)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
