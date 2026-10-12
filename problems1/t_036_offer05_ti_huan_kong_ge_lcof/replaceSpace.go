package t_036_offer05_ti_huan_kong_ge_lcof

func replaceSpace(s string) string {
	sbytes := []byte{}
	l := len(s)
	snew := []byte("%20")
	for i := 0; i < l; i++ {
		if s[i] == ' ' {
			sbytes = append(sbytes, snew...)
		} else {
			sbytes = append(sbytes, s[i])
		}
	}
	return string(sbytes)
}
