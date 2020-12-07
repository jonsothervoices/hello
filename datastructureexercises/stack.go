package datastructureexercises

type stack struct {
	data []interface{}
}
type minstack struct {
	data    []interface{}
	compare Comparer
	min     stack
}

//create new stack is O(n)
func newStack(in []interface{}) *stack {
	newS := stack{}
	for _, d := range in {
		newS.push(d)
	}
	return &newS
}

//length function
func (s *stack) len() int {
	return len(s.data)
}

//pop is O(1)
func (s *stack) pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	out := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return out
}

//push is O(1)
func (s *stack) push(in interface{}) {
	s.data = append(s.data, in)
}

//peek is O(1)
func (s *stack) peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

//create new minstack
func newMinStack(in []interface{}, c Comparer) *minstack {
	newS := minstack{compare: c}
	for _, d := range in {
		newS.push(d)
	}
	return &newS
}

//minstack length
func (s *minstack) len() int {
	return len(s.data)
}

//minstack pop
func (s *minstack) pop() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	out := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	if s.min.peek() != nil && s.compare(out, s.min.peek()) <= 0 {
		s.min.pop()
	}
	return out
}

//minstack push
func (s *minstack) push(in interface{}) {
	if s.min.peek() == nil || s.compare(s.min.peek(), in) > 0 {
		s.min.push(in)
	}
	s.data = append(s.data, in)
}

//minstack peek
func (s *minstack) peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[len(s.data)-1]
}

//min
func (s *minstack) findmin() interface{} {
	return s.min.peek()
}
