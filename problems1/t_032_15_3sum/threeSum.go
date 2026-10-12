package t_032_15_3sum

import "sort"

func threeSum(nums []int) [][]int {
	ret := [][]int{}
	if len(nums) < 3 {
		return ret
	}
	numsl := len(nums)
	// 排序nums
	sort.Ints(nums)
	// 三个指针 i l r
	for i := 0; i < numsl-2; i++ {
		l, r := i+1, numsl-1
		n1 := nums[i]
		// 如果最左值大于0 break
		if n1 > 0 {
			break
		}
		// 如果当前最左值与上一个值相等 说明已经遍历过 continue
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		// 调整 l r
		for l < r {
			n2, n3 := nums[l], nums[r]
			sum := n1 + n2 + n3
			// 三数相加大于0 左移最右指针
			if sum > 0 {
				r--
				// 三数相加小于0 右移中间指针
			} else if sum < 0 {
				l++
				// 三数相加等于0 符合条件 保存至 ret数组
			} else {
				ret = append(ret, []int{n1, n2, n3})
				// 保存之后 右移中间指针 pass掉与n2值相同的元素
				for l < r && n2 == nums[l] {
					l++
				}
				// 左移最右指针 pass掉与n3值相同的元素
				for l < r && n3 == nums[r] {
					r--
				}
			}
		}
	}
	return ret
}
