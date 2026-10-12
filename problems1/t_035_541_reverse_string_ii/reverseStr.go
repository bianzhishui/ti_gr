package t_035_541_reverse_string_ii

import (
	trs "github.com/bianzhishui/ti_gr/problems1/t_034_344_reverse_string"
)

func reverseStr(s string, k int) string {
	l := len(s)
	sbytes := []byte(s)
	for i := 0; i < l; i += 2 * k {
		if i+k <= l {
			trs.ReverseString(sbytes[i : i+k])
		} else {
			trs.ReverseString(sbytes[i:l])
		}
	}
	return string(sbytes)
}
