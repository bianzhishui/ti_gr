package t_034_344_reverse_string

func reverseString(s []byte) {
	l := len(s)
	for i := 0; i < l/2; i++ {
		j := l - 1 - i
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseString(s []byte) {
	reverseString(s)
}
