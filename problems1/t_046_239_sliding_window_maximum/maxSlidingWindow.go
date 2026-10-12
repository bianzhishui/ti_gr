package t_046_239_sliding_window_maximum

// 单调队列
type MyQueue struct {
	queue []int
}

func NewMyQueue() MyQueue {
	return MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Push(n int) {
	for !m.Empty() && n > m.GetBackVal() {
		m.PopBack()
	}
	m.PushBack(n)
}

func (m *MyQueue) PushBack(n int) {
	m.queue = append(m.queue, n)
}

func (m *MyQueue) Pop(n int) {
	if !m.Empty() && n == m.GetFrontVal() {
		m.PopFront()
	}
}

func (m *MyQueue) GetFrontVal() (n int) {
	return m.queue[0]
}

func (m *MyQueue) PopFront() {
	m.queue = m.queue[1:]
}

func (m *MyQueue) GetBackVal() (n int) {
	return m.queue[len(m.queue)-1]
}

func (m *MyQueue) PopBack() {
	m.queue = m.queue[:len(m.queue)-1]
}

func (m *MyQueue) GetMaxVal() int {
	return m.GetFrontVal()
}

func (m *MyQueue) Empty() bool {
	return len(m.queue) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	ret := make([]int, 0, len(nums)-k+1)
	q := NewMyQueue()
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	ret = append(ret, q.GetMaxVal())
	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		ret = append(ret, q.GetMaxVal())
	}
	return ret
}
