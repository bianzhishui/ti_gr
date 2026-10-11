package t_042_225_implement_stack_using_queues

type MyStack struct {
	queue []int
}

func Constructor() MyStack {
	return MyStack{
		queue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.queue = append(this.queue, x)
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	l := len(this.queue) - 1
	for l > 0 {
		this.Push(this.queue[0])
		this.queue = this.queue[1:]
		l--
	}
	value := this.queue[0]
	this.queue = this.queue[1:]
	return value
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	value := this.Pop()
	this.Push(value)
	return value
}

func (this *MyStack) Empty() bool {
	return len(this.queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
