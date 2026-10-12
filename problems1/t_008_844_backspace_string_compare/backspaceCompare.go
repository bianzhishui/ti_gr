package t_008_844_backspace_string_compare

// 重构字符串
// 时间复杂度：O(N+M)
// 空间复杂度：O(N+M)
// N,M 分别为字符串 s 和 t 的长度
func backspaceCompare(s string, t string) bool {
	//return processBackspace([]byte(s)) == processBackspace([]byte(t))
	return build(s) == build(t)
}

func processBackspace(s []byte) string {
	// 定义指针
	var pointer = 0
	// 遍历
	for i := 0; i < len(s); i++ {
		// 若第 i 个元素不等于 35 则赋值给 pointer 指向的位置 pointer++
		if s[i] != 35 {
			s[pointer] = s[i]
			pointer++
			// 若第 i 个元素等于 35 则判断 pointer 是否等于 0
		} else {
			//  若 pointer 不等于 0 则 pointer--
			if pointer != 0 {
				pointer--
			}
		}
	}
	// 返回新字符串
	return string(s[:pointer])
}

// 栈
func build(str string) string {
	// 定义栈
	var s []byte
	// 遍历
	for i := range str {
		// 若第 i 个元素不等于 # 则将它压入栈中
		if str[i] != '#' {
			s = append(s, str[i])
			// 若第 i 个元素等于 # 并且栈长度大于 0 则将最后一个元素弹出
		} else if len(s) > 0 {
			s = s[:len(s)-1]
		}
	}
	// 返回新字符串
	return string(s)
}

// 双指针
func backspaceCompare1(s, t string) bool {
	// 分别记录两个字符串中 # 的个数
	skipS, skipT := 0, 0
	// 分别记录两个字符串的长度 从后向前遍历
	i, j := len(s)-1, len(t)-1
	// 遍历
	for i >= 0 || j >= 0 {
		// 处理第一个字符串从后向前找到未删除的字符
		for i >= 0 {
			// 若字符串第 i 个元素等于 # 则 skipS++ i--
			if s[i] == '#' {
				skipS++
				i--
				// 若字符串第 i 个元素不等于 # 并且 # 个数大于 0 则 skipS-- i--
			} else if skipS > 0 {
				skipS--
				i--
				// 否则找到不需要删除的字符 跳出循环
			} else {
				break
			}
		}
		for j >= 0 {
			if t[j] == '#' {
				skipT++
				j--
			} else if skipT > 0 {
				skipT--
				j--
			} else {
				break
			}
		}
		// 上面两层循环有可能产生 i,j<0 的情况
		if i >= 0 && j >= 0 {
			// 若找到的两个未删除的值相等则 i--,j-- 继续循环
			// 若找到的两个未删除的值不相等则 return false
			if s[i] != t[j] {
				return false
			}
			// 若两个循环一个找到了一个没找到则 return false
			// 若 i,j 同时小于 0 则说明两个循环同时执行完 返回 true
		} else if i >= 0 || j >= 0 {
			return false
		}
		i--
		j--
	}
	return true
}
