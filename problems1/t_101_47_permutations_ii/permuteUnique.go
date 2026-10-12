package t_101_47_permutations_ii

import "sort"

func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	ints := []int{}
	// used 判断元素是否使用过
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		// 找到符合题意的组合
		if len(ints) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, ints)
			result = append(result, temp)
			return
		}
		for i := 0; i < len(nums); i++ {
			// 当前元素和前一个元素相同 且 前一个元素在本层使用过
			if i > 0 && nums[i] == nums[i-1] && used[i-1] == false {
				continue
			}
			// 说明当前元素在上层使用过
			if used[i] { // == true
				continue
			}
			ints = append(ints, nums[i])
			used[i] = true
			backtrack()
			// 回溯
			ints = ints[:len(ints)-1]
			used[i] = false
		}
	}
	// 去重需要 排序
	sort.Ints(nums)
	backtrack()
	return result
}
