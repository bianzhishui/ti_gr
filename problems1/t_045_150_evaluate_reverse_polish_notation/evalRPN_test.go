package t_045_150_evaluate_reverse_polish_notation

import (
	"testing"
)

func TestEvalRPN(t *testing.T) {
	var (
		tokens   = []string{"2", "1", "+", "3", "*"}
		expected = 9
		actual   int
	)
	actual = evalRPN(tokens)
	if expected != actual {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
