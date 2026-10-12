package t_054_637_average_of_levels_in_binary_tree

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	var (
		root     = []int{3, 9, 20, 0, 0, 15, 7}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{3 * 100000, 14.5 * 100000, 11 * 100000}
		actual   = make([]int, 0, len(expected))
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &TreeNode{Val: val}
		}
	}
	for index, val := range queue {
		if val == nil {
			continue
		}
		i := index * 2
		if i+1 < l && queue[i+1] != nil {
			val.Left = queue[i+1]
		}
		if i+2 < l && queue[i+2] != nil {
			val.Right = queue[i+2]
		}
	}
	floatList := averageOfLevels(queue[0])
	for _, val := range floatList {
		actual = append(actual, int(val*100000))
	}
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
