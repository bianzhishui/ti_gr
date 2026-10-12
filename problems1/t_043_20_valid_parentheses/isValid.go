package t_043_20_valid_parentheses

type Stack struct {
	bytes []byte
	l     int
}

func NewStack() Stack {
	return Stack{
		bytes: make([]byte, 0),
		l:     0,
	}
}

func (s *Stack) Push(b byte) {
	s.bytes = append(s.bytes, b)
	s.l++
}

func (s *Stack) Pop() (b byte) {
	if s.Empty() {
		return '0'
	}
	s.l--
	val := s.bytes[s.l]
	s.bytes = s.bytes[:s.l]
	return val
}

func (s *Stack) Top() (b byte) {
	if s.Empty() {
		return '0'
	}
	return s.bytes[s.l-1]
}

func (s *Stack) Empty() bool {
	return s.l == 0
}

func isValid(s string) bool {
	stack := NewStack()
	byteMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	for _, s1 := range s {
		b := byte(s1)
		if val, ok := byteMap[b]; ok {
			stack.Push(val)
		} else {
			if stack.Empty() {
				return false
			} else {
				b1 := stack.Pop()
				if b != b1 {
					return false
				}
			}
		}

	}
	return stack.l == 0
}
