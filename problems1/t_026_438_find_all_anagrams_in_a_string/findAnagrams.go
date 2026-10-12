package t_026_438_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	rList := []int{}
	// 两个字符串长度
	sLen, pLen := len(s), len(p)
	if sLen < pLen {
		return rList
	}
	// 记录 26 个字符 差值
	var charList [26]int
	// 遍历 p 长度的 s 和 p
	for i, char := range p {
		charList[s[i]-'a']++
		charList[char-'a']--
	}
	// 差值 记录的事不等于 0 的字符个数
	differ := 0
	for _, char := range charList {
		if char != 0 {
			differ++
		}
	}
	// 若 differ 为 0 说明 两字符串起始位置开始为异位词
	if differ == 0 {
		rList = append(rList, 0)
	}
	// 继续向后遍历 滑动窗口
	for i, char := range s[:sLen-pLen] {
		// 窗口起始位置的字符个数等于 1 则 differ 减一
		if charList[char-'a'] == 1 {
			differ--
			// 窗口起始位置的字符个数等于 0 则 differ 加一
		} else if charList[char-'a'] == 0 {
			differ++
		}
		// 窗口起始位置的字符个数减一
		charList[char-'a']--
		// 窗口末尾位置的字符个数等于 -1 则 differ 减一
		if charList[s[i+pLen]-'a'] == -1 {
			differ--
			// 窗口末尾位置的字符串个数等于 0 则 differ 加一
		} else if charList[s[i+pLen]-'a'] == 0 {
			differ++
		}
		// 窗口末尾位置的字符个数加一
		charList[s[i+pLen]-'a']++
		// 若此时 differ 值等于 0 则添加到返回切片中 记录的值为 i+1
		if differ == 0 {
			rList = append(rList, i+1)
		}
	}
	return rList
}
