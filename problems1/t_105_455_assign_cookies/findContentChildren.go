package t_105_455_assign_cookies

import "sort"

func findContentChildren(g []int, s []int) int {
	result := 0
	// 排序
	sort.Ints(s)
	sort.Ints(g)
	// 找到最大饼干 也就是排序后最后一个饼干
	index := len(s) - 1
	for i := len(g) - 1; i >= 0 && index >= 0; i-- {
		// 如果这个饼干大于等于g[i]孩子的胃口 则给他
		if s[index] >= g[i] {
			result++
			index--
		}
	}
	return result
}
