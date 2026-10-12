package t_061_226_invert_binary_tree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	var (
		root     = []int{4, 2, 7, 1, 3, 6, 9}
		l        = len(root)
		queue    = make([]*TreeNode, l)
		expected = []int{4, 7, 2, 9, 6, 3, 1}
		actual   []int
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
	node := invertTree1(queue[0])
	// 层序遍历
	if node != nil {
		queue2 := []*TreeNode{}
		queue2 = append(queue2, node)
		count := 1
		for len(queue2) != 0 {
			for i := 0; i < count; i++ {
				node2 := queue2[0]
				queue2 = queue2[1:]
				actual = append(actual, node2.Val)
				if node2.Left != nil {
					queue2 = append(queue2, node2.Left)
				}
				if node2.Right != nil {
					queue2 = append(queue2, node2.Right)
				}
			}
			count = len(queue2)
		}
	}
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
