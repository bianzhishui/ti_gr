package t_105_455_assign_cookies

import (
	"fmt"
	"testing"
)

func TestFindContentChildren(t *testing.T) {
	var (
		g        = []int{1, 2}
		s        = []int{1, 2, 3}
		expected = 2
		actual   int
	)
	actual = findContentChildren(g, s)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if expected == actual {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
