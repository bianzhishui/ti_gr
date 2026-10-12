package t_033_18_4sum

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	// 数组排序
	sort.Ints(nums)
	numsl := len(nums)
	if len(nums) < 4 || target < 0 && nums[numsl-1] < target {
		return ret
	}
	for i := 0; i < numsl-3; i++ {
		n1 := nums[i]
		// 剪枝
		if n1 > target && n1 > 0 && target > 0 {
			break
		}
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		for j := i + 1; j < numsl-2; j++ {
			l, r := j+1, numsl-1
			n2 := nums[j]
			n1n2 := n1 + n2
			// 二级剪枝
			if n1n2 > target && n1n2 > 0 && target > 0 {
				break
			}
			if j > i+1 && n2 == nums[j-1] {
				continue
			}
			for l < r {
				n3, n4 := nums[l], nums[r]
				sum := n1 + n2 + n3 + n4
				if sum > target {
					r--
				} else if sum < target {
					l++
				} else {
					ret = append(ret, []int{n1, n2, n3, n4})
					for l < r && n3 == nums[l] {
						l++
					}
					for l < r && n4 == nums[r] {
						r--
					}
				}
			}
		}
	}
	return ret
}
