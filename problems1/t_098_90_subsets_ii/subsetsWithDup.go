package t_098_90_subsets_ii

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{}
	ints := []int{}
	var backtrack func(index int)
	backtrack = func(index int) {
		temp := make([]int, len(ints))
		copy(temp, ints)
		// 添加组合
		result = append(result, temp)
		for i := index; i < len(nums); i++ {
			// 同一树层 过滤相同的元素
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			ints = append(ints, nums[i])
			backtrack(i + 1)
			// 回溯
			ints = ints[:len(ints)-1]
		}
	}
	// 去重需要排序
	sort.Ints(nums)
	backtrack(0)
	return result
}
