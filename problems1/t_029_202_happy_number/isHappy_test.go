package t_029_202_happy_number

import "testing"

func TestIsHappy(t *testing.T) {
	var (
		n        = 7
		expected = true
		actual   bool
	)

	actual = isHappy(n)
	if expected == actual {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
