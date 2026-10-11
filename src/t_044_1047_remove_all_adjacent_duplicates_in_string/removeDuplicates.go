package t_044_1047_remove_all_adjacent_duplicates_in_string

import (
	t "github.com/bianzhishui/ti_gr/src/t_043_20_valid_parentheses"
)

func removeDuplicates(s string) string {
	stack := t.NewStack()
	for _, s1 := range s {
		b := byte(s1)
		if stack.Empty() {
			stack.Push(b)
		} else {
			b1 := stack.Top()
			if b == b1 {
				stack.Pop()
			} else {
				stack.Push(b)
			}
		}
	}
	bytes := []byte{}
	for !stack.Empty() {
		b := stack.Pop()
		bytes = append([]byte{b}, bytes...)
	}
	return string(bytes)
}

func removeDuplicates1(s string) string {
	var stack []byte
	for i := 0; i < len(s); i++ {
		// 栈不空 且 与栈顶元素不等
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			// 弹出栈顶元素 并 忽略当前元素(s[i])
			stack = stack[:len(stack)-1]
		} else {
			// 入栈
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
