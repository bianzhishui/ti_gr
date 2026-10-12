package t_037_151_reverse_words_in_a_string

import (
	trs "github.com/bianzhishui/ti_gr/problems1/t_034_344_reverse_string"
)

func reverseWords(s string) string {
	// string -> []byte{}
	sbytes := []byte(s)
	// 快慢指针
	slow, fast := 0, 0
	// s长度
	l := len(s)
	// 快指针遍历
	for ; fast < l; fast++ {
		// 快指针位置不为空格
		if sbytes[fast] != ' ' {
			// 慢指针不为开头则添加空格
			if slow != 0 {
				sbytes[slow] = ' '
				slow++
			}
			// 将快指针指向的单词移动到慢指针位置
			for fast < l && sbytes[fast] != ' ' {
				sbytes[slow] = s[fast]
				slow++
				fast++
			}
		}
	}
	start := 0
	// 遍历每个单词进行翻转
	for i := 0; i <= slow; i++ {
		if i == slow || sbytes[i] == ' ' {
			trs.ReverseString(sbytes[start:i])
			start = i + 1
		}
	}
	// 整个字符串翻转
	trs.ReverseString(sbytes[:slow])
	return string(sbytes[:slow])
}

func reverseWords1(s string) string {
	sbytes := []byte{}
	l := len(s)
	for fast := l - 1; fast >= 0; fast-- {
		if s[fast] != ' ' {
			slow := fast + 1
			for fast >= 0 && s[fast] != ' ' {
				fast--
			}
			sbytes = append(sbytes, s[fast+1:slow]...)
			sbytes = append(sbytes, ' ')
		}
	}
	return string(sbytes[:len(sbytes)-1])
}
