package t_047_347_top_k_frequent_elements

import (
	"container/heap"
)

// 小顶堆
type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	// numsMap统计频率
	numsMap := make(map[int]int)
	for _, num := range nums {
		numsMap[num]++
	}
	// 小顶堆
	h := &IHeap{}
	heap.Init(h)
	for key, val := range numsMap {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	// 返回
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}
