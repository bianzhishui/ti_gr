package t_038_offer58ii_zuo_xuan_zhuan_zi_fu_chuan_lcof

import (
	trs "github.com/bianzhishui/ti_gr/src/t_034_344_reverse_string"
)

func reverseLeftWords(s string, n int) string {
	ret := []byte{}
	for i := n; i < len(s); i++ {
		ret = append(ret, s[i])
	}
	for i := 0; i < n; i++ {
		ret = append(ret, s[i])
	}
	return string(ret)
}

func reverseLeftWords1(s string, n int) string {
	sbytes := []byte(s)
	trs.ReverseString(sbytes[:n])
	trs.ReverseString(sbytes[n:])
	trs.ReverseString(sbytes)
	return string(sbytes)
}
