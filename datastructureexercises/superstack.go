package datastructureexercises

type superstack struct {
	stacks []stack
}

const maxLen = 5

func (s *superstack) numStacks() int {
	return len(s.stacks)
}

//NewSuperStack Create Superstack from raw data
func newSuperStack(in []interface{}) *superstack {
	newS := superstack{}
	for _, d := range in {
		newS.push(d)
	}
	return &newS
}

//length function
func (s *superstack) len() int {
	totalLength := 0
	for _, n := range s.stacks {
		totalLength += n.len()
	}
	return totalLength
}

//pop is O(0)
func (s *superstack) pop() interface{} {
	if s.len() == 0 {
		return nil
	}
	v := s.stacks[len(s.stacks)-1].pop()
	if s.stacks[len(s.stacks)-1].len() == 0 {
		s.stacks = s.stacks[:len(s.stacks)-1]
	}
	return v
}

//push is O(0)
func (s *superstack) push(in interface{}) {
	if s.numStacks() == 0 {
		s.stacks = append(s.stacks, stack{})
	}
	if s.stacks[len(s.stacks)-1].len() == 5 {
		s.stacks = append(s.stacks, stack{})
	}
	s.stacks[len(s.stacks)-1].push(in)
}

//peek is O(0)
func (s *superstack) peek() interface{} {
	if s.numStacks() == 0 {
		return nil
	}
	return s.stacks[len(s.stacks)-1].peek()
}
