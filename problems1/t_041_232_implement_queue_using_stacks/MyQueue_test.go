package t_041_232_implement_queue_using_stacks

import (
	"testing"
)

func TestMyQueue_Push(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	first := queue.inStack[0]
	last := queue.inStack[len(queue.inStack)-1]
	if first != 1 {
		t.Fatalf("expected=1, actual=%v", first)
	}
	if last != 3 {
		t.Fatalf("expected=3, actual=%v", last)
	}
}

func TestMyQueue_Pop(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	actual := queue.Pop()
	if actual != 1 {
		t.Fatalf("expected=1, actual=%v", actual)
	}
	actual = queue.Pop()
	if actual != 2 {
		t.Fatalf("expected=2, actual=%v", actual)
	}
	actual = queue.Pop()
	if actual != 3 {
		t.Fatalf("expected=3, actual=%v", actual)
	}
	empty := queue.Empty()
	if !empty {
		t.Fatalf("expected=true, actual=%v", actual)
	}
}

func TestMyQueue_Peek(t *testing.T) {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	actual := queue.Peek()
	if actual != 1 {
		t.Fatalf("expected=3, actual=%v", actual)
	}
	empty := queue.Empty()
	if empty {
		t.Fatalf("expected=false, actual=%v", actual)
	}
}

func TestMyQueue_Empty(t *testing.T) {
	queue := Constructor()
	actual := queue.Empty()
	if !actual {
		t.Fatalf("expected=true, actual=%v", actual)
	}
	queue.Push(1)
	actual = queue.Empty()
	if actual {
		t.Fatalf("expected=false, actual=%v", actual)
	}
}
