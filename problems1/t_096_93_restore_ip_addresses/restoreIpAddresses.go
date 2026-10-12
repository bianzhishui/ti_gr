package t_096_93_restore_ip_addresses

func restoreIpAddresses(s string) []string {
	result := []string{}
	// 记录 . 的个数
	pointNum := 0
	str := ""
	// 判断是否符合ip数字 0<=x<=255
	var isIpNum = func(start, end int) bool {
		// 排除 011 类似数
		if end < start || end >= len(s) || (end > start && s[start] == '0') {
			return false
		}
		// 计算 字符串数字和
		sum := 0
		for i := start; i <= end; i++ {
			if s[i] < '0' || s[i] > '9' {
				return false
			}
			sum = sum*10 + int(s[i]-'0')
			if sum > 255 {
				return false
			}
		}
		return true
	}
	var backtrack func(index int)
	backtrack = func(index int) {
		// 若 . 的个数等于3个 则判断第四个数是否符合要求 若符合要求加入result中
		if pointNum == 3 {
			if isIpNum(index, len(s)-1) {
				result = append(result, str+s[index:])
			}
			return
		}
		for i := index; i < len(s); i++ {
			// 判断遍历到的数是否为ip数 若不符合直接break
			if isIpNum(index, i) {
				str += s[index:i+1] + "."
				pointNum++
				backtrack(i + 1)
				// 回溯
				str = str[:len(str)-1-(i+1-index)]
				pointNum--
			} else {
				break
			}
		}
	}
	backtrack(0)
	return result
}
