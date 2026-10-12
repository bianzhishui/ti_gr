package t_010_977_squares_of_a_sorted_array

import (
	"fmt"
	"reflect"
	"testing"
)

func main() {
	var nums = []int{-7, -3, 2, 3, 11}
	fmt.Println(nums)
	fmt.Println(sortedSquares1(nums))
}

func TestSortedSquares(t *testing.T) {
	var (
		nums     = []int{-7, -3, 2, 3, 11}
		expected = []int{4, 9, 9, 49, 121}
		actual   []int
	)
	actual = sortedSquares(nums)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}

func TestSortedSquares1(t *testing.T) {
	var (
		nums     = []int{-7, -3, 2, 3, 11}
		expected = []int{4, 9, 9, 49, 121}
		actual   []int
	)
	actual = sortedSquares1(nums)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
	t.Log("done")
}
