package stackexercises

type stringStack struct {
	data []string
}

//pop is O(0)
func (s *stringStack) pop() string {
	if len(s.data) == 0 {
		return ""
	}
	out := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return out
}

//push is O(0)
func (s *stringStack) push(in string) {
	s.data = append(s.data, in)
}

//peek is O(0)
func (s *stringStack) peek() string {
	if len(s.data) == 0 {
		return ""
	}
	return s.data[len(s.data)-1]
}
