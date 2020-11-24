package datastructureexercises

type stack struct {
	data []interface{}
}

//length functtion
func (s *stack) len() int {
	return len(s.data)
}

//pop is O(0)
func (s *stack) pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	out := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return out
}

//push is O(0)
func (s *stack) push(in interface{}) {
	s.data = append(s.data, in)
}

//peek is O(0)
func (s *stack) peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}
