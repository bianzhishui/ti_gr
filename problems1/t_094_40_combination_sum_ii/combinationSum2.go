package t_094_40_combination_sum_ii

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := [][]int{}
	ints := []int{}
	sum := 0
	// 排序
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
		// pre处理同一层相同的元素 避免最后组合出现重复的现象 因为candidates切片中会有重复的元素 但是result中不允许出现相同元素的组合
		// && sum+candidates[i] <= target; 剪枝
		pre := -1
		for i := index; i < len(candidates) && sum+candidates[i] <= target; i++ {
			if pre != -1 && candidates[i] == candidates[pre] {
				continue
			}
			// pre记录前一个元素值
			pre = i
			sum += candidates[i]
			ints = append(ints, candidates[i])
			backtrack(i + 1)
			sum -= candidates[i]
			ints = ints[:len(ints)-1]
		}
	}
	backtrack(0)
	return result
}
