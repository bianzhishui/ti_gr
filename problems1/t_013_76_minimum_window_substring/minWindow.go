package t_013_76_minimum_window_substring

// 滑动窗口
func minWindow(s string, t string) string {
	// t 字符串类型 map
	tTypeMap := map[byte]int{}
	// s 字符串类型 map
	sTypeMap := map[byte]int{}
	// 生成 t 字符串类型 : 个数
	for i := 0; i < len(t); i++ {
		tTypeMap[t[i]]++
	}
	// 字符串左右指针
	sLeft, sRight := -1, -1
	// s 字符串长度
	sLen := len(s)
	// 临时变量 记录当前大小
	tmpLen := sLen + 1
	// 检查函数
	// 逻辑: 遍历 tTypeMap 若 sTypeMap 中对应的记录数小于 tTypeMap 中记录数 则不符合题意 返回 false 否则 遍历完 tTypeMap 返回 true
	check := func() bool {
		for k, v := range tTypeMap {
			if sTypeMap[k] < v {
				return false
			}
		}
		return true
	}
	// 遍历字符串 s
	for left, right := 0, 0; right < sLen; right++ {
		// 若遍历到的字符在 tTypeMap 中 则 sTypeMap 记录数 ++
		if tTypeMap[s[right]] > 0 {
			sTypeMap[s[right]]++
		}
		// 检查通过符合题意 则更改临时变量 修改返回的左右指针 并且左指针右移 sTypeMap 相应记录数 --
		for check() && left <= right {
			if right-left+1 < tmpLen {
				tmpLen = right - left + 1
				sLeft = left
				sRight = right + 1
			}
			if _, ok := sTypeMap[s[left]]; ok {
				sTypeMap[s[left]]--
			}
			left++
		}
	}
	// 若返回值左指针为 -1 则说明没有取到符合题意的字符串 返回空字符串
	if sLeft == -1 {
		return ""
	}
	// 否则返回左右指针对应的字符串
	return s[sLeft:sRight]
}
