package t_041_232_implement_queue_using_stacks

// 队列
type MyQueue struct {
	// 栈实现队列
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  make([]int, 0),
		outStack: make([]int, 0),
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)

}

func (this *MyQueue) Pop() int {
	if this.Empty() {
		return 0
	}
	ret := this.Peek()
	this.outStack = this.outStack[:len(this.outStack)-1]
	return ret
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	if len(this.outStack) == 0 {
		l := len(this.inStack)
		for l > 0 {
			this.outStack = append(this.outStack, this.inStack[l-1])
			this.inStack = this.inStack[:l-1]
			l = len(this.inStack)
		}
	}
	ret := this.outStack[len(this.outStack)-1]
	return ret
}

func (this *MyQueue) Empty() bool {
	return len(this.inStack) == 0 && len(this.outStack) == 0
}

/**
* Your MyQueue object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(x);
* param_2 := obj.Pop();
* param_3 := obj.Peek();
* param_4 := obj.Empty();
 */
