package datastructureexercises

type superstack struct {
	stacks []stack
}

const maxLen = 5

func (s *superstack) numStacks() int {
	return len(s.stacks)
}

//length function
func (s *superstack) len() int {
	//return len(s.data)
	return 0
}

//pop is O(0)
func (s *superstack) pop() interface{} {
	// if len(s.data) == 0 {
	// 	return nil
	// }
	// out := s.data[len(s.data)-1]
	// s.data = s.data[:len(s.data)-1]
	return nil
}

//push is O(0)
func (s *superstack) push(in interface{}) {
	// s.data = append(s.data, in)
}

//peek is O(0)
func (s *superstack) peek() interface{} {
	// if len(s.data) == 0 {
	// 	return nil
	// }
	// return s.data[len(s.data)-1]
	return nil
}
