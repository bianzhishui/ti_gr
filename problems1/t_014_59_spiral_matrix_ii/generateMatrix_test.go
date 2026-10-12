package t_014_59_spiral_matrix_ii

import (
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	var (
		n        = 2
		expected = [][]int{{1, 2}, {4, 3}}
		actual   [][]int
	)
	actual = generateMatrix(n)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestGenerateMatrix1(t *testing.T) {
	var (
		n        = 2
		expected = [][]int{{1, 2}, {4, 3}}
		actual   [][]int
	)
	actual = generateMatrix1(n)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
