package t_103_51_n_queens

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	var (
		n        = 4
		expected = [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}
		actual [][]string
	)
	actual = solveNQueens(n)
	logprint := fmt.Sprintf("expected=%v, actual=%v", expected, actual)
	if reflect.DeepEqual(expected, actual) {
		t.Log(logprint)
	} else {
		t.Fatal(logprint)
	}
}
