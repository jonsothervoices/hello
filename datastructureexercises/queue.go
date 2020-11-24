package datastructureexercises

type queue struct {
	data []interface{}
}

//length functtion
func (s *queue) len() int {
	return len(s.data)
}

//remove is O(0)\
func (s *queue) remove() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	out := s.data[0]
	s.data = s.data[1:]
	return out
}

//add is O(0)
func (s *queue) add(in interface{}) {
	s.data = append(s.data, in)
}

//peek is O(0)
func (s *queue) peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}
