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

//Merge Sort
func mergeSort(s sortableSlice) sortableSlice {
	if len(s) < 2 {
		return s
	}
	mid := len(s) / 2
	return sortSlices(mergeSort(s[:mid]), mergeSort(s[mid:]))
}

func sortSlices(a, b sortableSlice) sortableSlice {
	s := sortableSlice{}
	for len(a) > 0 || len(b) > 0 {
		if len(b) == 0 {
			s = append(s, a[0])
			a = removeFirst(a)
			continue
		}
		if len(a) == 0 {
			s = append(s, b[0])
			b = removeFirst(b)
			continue
		}
		if a[0] < b[0] {
			s = append(s, a[0])
			a = removeFirst(a)
			continue
		}
		s = append(s, b[0])
		b = removeFirst(b)
	}
	return s
}

func removeFirst(s sortableSlice) sortableSlice {
	if len(s) <= 1 {
		return sortableSlice{}
	}
	return s[1:]
}

//10.1: sorted Merge: Given 2 sorted arrays A and B, where A has a large enough buffer at the end to hold B, merge B into A.