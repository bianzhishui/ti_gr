package t_093_39_combination_sum

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	result := [][]int{}
	ints := []int{}
	sum := 0
	// 为了剪枝需要先进行排序
	sort.Ints(candidates)
	var backtrack func(index int)
	backtrack = func(index int) {
		// 找到符合题意的组合
		if sum == target {
			temp := make([]int, len(ints))
			copy(temp, ints)
			result = append(result, temp)
			return
		}
		// && sum+candidates[i] <= target; 剪枝
		for i := index; i < len(candidates) && sum+candidates[i] <= target; i++ {
			sum += candidates[i]
			ints = append(ints, candidates[i])
			// 入参为i 因为单个元素可以重复
			backtrack(i)
			// 回溯
			sum -= candidates[i]
			ints = ints[:len(ints)-1]
		}
	}
	backtrack(0)
	return result
}
