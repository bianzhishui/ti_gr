package t_095_131_palindrome_partitioning

func partition(s string) [][]string {
	result := [][]string{}
	slen := len(s)
	str := []string{}
	var backtrack func(index int)
	var isPalindrome func(start, end int) bool
	// 判断是否回文
	isPalindrome = func(start, end int) bool {
		for start < end {
			if s[start] != s[end] {
				return false
			}
			start++
			end--
		}
		return true
	}
	backtrack = func(index int) {
		// 找到符合要求的字符组合 遍历到的index位置为末尾说明遍历完找到了符合题意的字符组合
		if index >= slen {
			temp := make([]string, len(str))
			copy(temp, str)
			result = append(result, temp)
			return
		}
		for i := index; i < slen; i++ {
			// 判断回文 不是回文不会继续统计下去进入下一轮循环
			if isPalindrome(index, i) {
				str = append(str, string(s[index:i+1]))
				backtrack(i + 1)
				str = str[:len(str)-1]
			}
		}
	}
	backtrack(0)
	return result
}
