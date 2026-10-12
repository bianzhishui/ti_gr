package t_057_116_populating_next_right_pointers_in_each_node

import (
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	var (
		root     = []int{1, 2, 3, 4, 5, 6, 7}
		l        = len(root)
		queue    = make([]*Node, l)
		expected = [][]int{{1}, {2, 3}, {4, 5, 6, 7}}
		actual   [][]int
	)
	for index, val := range root {
		if val != 0 {
			queue[index] = &Node{Val: val}
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
	node := connect(queue[0])
	for node != nil {
		nums := []int{}
		node2 := node
		for node2 != nil {
			nums = append(nums, node2.Val)
			node2 = node2.Next
		}
		actual = append(actual, nums)
		node = node.Left
	}
	if reflect.DeepEqual(expected, actual) {
		t.Logf("expected=%v, actual=%v", expected, actual)
	} else {
		t.Fatalf("expected=%v, actual=%v", expected, actual)
	}
}
