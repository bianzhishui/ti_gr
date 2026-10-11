package t_106_376_wiggle_subsequence

import (
	"fmt"
	"testing"
)

func TestWiggleMaxLength(t *testing.T) {
	var (
		nums     = []int{1, 7, 4, 9, 2, 5}
		expected = 6
		actual   int
	)
	actual = wiggleMaxLength(nums)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if expected == actual {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
