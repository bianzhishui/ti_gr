package t_074_113_path_sum_ii

import (
	"reflect"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	var (
		root      = []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 0, 5, 1}
		targetSum = 22
		l         = len(root)
		queue     = make([]*TreeNode, l)
		expected  = [][]int{{5, 4, 11, 2}, {5, 8, 4, 5}}
		actual    [][]int
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
	actual = pathSum(queue[0], targetSum)
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
