package datastructureexercises

const simpleQueueType = "simpleQueue"
const stackQueueType = "stackQueue"

type queue interface {
	len() int
	remove() interface{}
	add(interface{})
	peek() interface{}
}
type simpleQueue struct {
	data []interface{}
}
type stackQueue struct {
	stackdata    stack
	stackmapping stack
}

func queueFactory(t string, data []interface{}) queue {
	var newQueue queue
	switch t {
	case stackQueueType:
		newQueue = &stackQueue{}
	default:
		newQueue = &simpleQueue{}
	}
	for _, d := range data {
		newQueue.add(d)
	}
	return newQueue
}

//simple length function
func (s *simpleQueue) len() int {
	return len(s.data)
}

//remove is O(0)
func (s *simpleQueue) remove() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	out := s.data[0]
	s.data = s.data[1:]
	return out
}

//add is O(0)
func (s *simpleQueue) add(in interface{}) {
	s.data = append(s.data, in)
}

//peek is O(0)
func (s *simpleQueue) peek() interface{} {
	if len(s.data) == 0 {
		return nil
	}
	return s.data[0]
}

//stack length function
func (s *stackQueue) len() int {
	return s.stackdata.len()
}

//stack remove
func (s *stackQueue) remove() interface{} {
	if s.len() == 0 {
		return nil
	}
	slen := s.stackdata.len()
	for i := 0; i < slen; i++ {
		s.stackmapping.push(s.stackdata.pop())
	}
	v := s.stackmapping.pop()
	for i := 0; i < slen-1; i++ {
		s.stackdata.push(s.stackmapping.pop())
	}
	return v
}

//stack add
func (s *stackQueue) add(in interface{}) {
	slen := s.len()
	s.stackmapping.push(in)
	for i := 0; i < slen; i++ {
		s.stackmapping.push(s.stackdata.pop())
	}
	for i := 0; i < slen+1; i++ {
		s.stackdata.push(s.stackmapping.pop())
	}
}

//stack peek
func (s *stackQueue) peek() interface{} {
	if s.stackdata.len() == 0 {
		return nil
	}
	slen := s.stackdata.len()
	for i := 0; i < slen; i++ {
		s.stackmapping.push(s.stackdata.pop())
	}
	v := s.stackmapping.peek()
	for i := 0; i < slen-1; i++ {
		s.stackdata.push(s.stackmapping.pop())
	}
	return v
}
