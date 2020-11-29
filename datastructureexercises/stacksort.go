package datastructureexercises

type comparer func(a, b interface{}) int

//Write a program to sort a stack such that the smallest items are on the top.
//stackSorter is O(n^2)
func stackSorter(s *stack, srt comparer) (final stack) {
	for s.len() > 0 {
		vplace := s.pop()
		for srt(vplace, final.peek()) > 0 {
			s.push(final.pop())
		}
		final.push(vplace)
	}
	return
}
