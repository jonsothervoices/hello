package sortsearch

type sortableSlice []int

//Bubble Sort
//12:40~~1:00
func (s sortableSlice) bubbleSort() {
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
	return
}

//Selection Sort
//1:14~~1:43 (totally got it at 2:01)
func (s sortableSlice) selectionSort() {
	for start := range s {
		min := start
		for i := start; i < len(s); i++ {
			if s[i] < s[min] {
				min = i
			}
		}
		s[start], s[min] = s[min], s[start]
	}
	return
}

//Insertion Sort
func (s sortableSlice) insertionSort() {
	for start := range s {
		for i := start; i > 0; i-- {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
			} else {
				break
			}
		}
	}
	return
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

//Quick Sort
//2:57~~3:40

func (s sortableSlice) quickSort() {
	if len(s) < 2 {
		return
	}
	hi := len(s) - 1
	pivot := s[hi]
	i := 0
	for j := range s {
		if s[j] < pivot {
			s[i], s[j] = s[j], s[i]
			i++
		}
	}
	s[i], s[hi] = s[hi], s[i]
	s[:i].quickSort()
	s[i+1:].quickSort()
	return
}

//Radix Sort (max 4-digit numbers)
func (s sortableSlice) radixSort() {
	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	digits := 0
	for max > 0 {
		digits++
		max /= 10
		buckets := [10]sortableSlice{}
		for _, v := range s {
			n := getDigit(v, digits, 1)
			buckets[n] = append(buckets[n], v)
		}
		for i := range buckets {
			if i == 0 {
				continue
			}
			buckets[i] = append(buckets[i-1], buckets[i]...)
		}
		for i := range s {
			s[i] = buckets[9][i]
		}
	}
	return
}

func getDigit(a, digit, current int) int {
	if current == digit {
		return a % 10
	}
	return getDigit(a/10, digit, current+1)
}

//10.1: sorted Merge: Given 2 sorted arrays A and B, where A has a large enough buffer at the end to hold B, merge B into A.
