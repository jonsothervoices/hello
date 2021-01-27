package sortsearch

type sortableSlice []int

//Bubble Sort
//12:40~~1:00
func (s sortableSlice) bubbleSort() sortableSlice {
	swapped := true
	n := len(s)
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
				swapped = true
			}
		}
		n--
	}
	return s
}

//Selection Sort
//1:14~~1:43 (totally got it at 2:01)
func (s sortableSlice) selectionSort() sortableSlice {
	for start := range s {
		min := start
		for i := start; i < len(s); i++ {
			if s[i] < s[min] {
				min = i
			}
		}
		s[start], s[min] = s[min], s[start]
	}
	return s
}

//Insertion Sort
func (s sortableSlice) insertionSort() sortableSlice {
	for start := range s {
		for i := start; i > 0; i-- {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
			} else {
				break
			}
		}
	}
	return s
}

//10.1: sorted Merge: Given 2 sorted arrays A and B, where A has a large enough buffer at the end to hold B, merge B into A.
