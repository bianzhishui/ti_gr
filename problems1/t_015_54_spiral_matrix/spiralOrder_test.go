package t_015_54_spiral_matrix

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	var (
		matrix   = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		expected = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		actual   []int
	)
	actual = spiralOrder(matrix)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestSpiralOrder1(t *testing.T) {
	var (
		matrix   = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
		expected = []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		actual   []int
	)
	actual = spiralOrder1(matrix)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
