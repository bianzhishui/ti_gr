package t_042_225_implement_stack_using_queues

import (
	"testing"
)

func TestMyStack_Push(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	first := stack.queue[len(stack.queue)-1]
	last := stack.queue[0]
	if first != 3 {
		t.Fatalf("expected=3, actual=%v", first)
	}
	if last != 1 {
		t.Fatalf("expected=1, actual=%v", first)
	}
}

func TestMyStack_Pop(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	actual := stack.Pop()
	if actual != 3 {
		t.Fatalf("expected=3, actual=%v", actual)
	}
	actual = stack.Pop()
	if actual != 2 {
		t.Fatalf("expected=2, actual=%v", actual)
	}
	actual = stack.Pop()
	if actual != 1 {
		t.Fatalf("expected=1, actual=%v", actual)
	}
}

func TestMyStack_Top(t *testing.T) {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	actual := stack.Top()
	if actual != 3 {
		t.Fatalf("expected=3, actual=%v", actual)
	}
	actual = len(stack.queue)
	if actual != 3 {
		t.Fatalf("expected=3, actual=%v", actual)
	}
}

func TestMyStack_Empty(t *testing.T) {
	stack := Constructor()
	actual := stack.Empty()
	if !actual {
		t.Fatalf("expected=true, actual=%v", actual)
	}
	stack.Push(1)
	actual = stack.Empty()
	if actual {
		t.Fatalf("expected=false, actual=%v", actual)
	}
}
