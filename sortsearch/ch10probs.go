package sortsearch

import "math"

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

//Radix Sort (base 10)
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

//Standard Binary Search
func binarySearch(s sortableSlice, v int) int {
	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		if s[0] == v {
			return 0
		}
		return -1
	}
	//find midpoint
	mid := len(s) / 2
	//if midpoint is greater than v, return recurse on first half
	if s[mid] == v {
		return mid
	}
	if s[mid] > v {
		return binarySearch(s[:mid], v)
	}
	//if less, recurse on second half
	ret := binarySearch(s[mid:], v)
	//if rec is -1, return -1
	if ret >= 0 {
		return ret + mid
	}
	return ret
}

//Binary search (return first))
//3:41~~5:07
func binarySearchWithDupes(s sortableSlice, v int) int {
	if len(s) == 0 {
		return -1
	}
	if len(s) == 1 {
		if s[0] == v {
			return 0
		}
		return -1
	}
	//find midpoint
	mid := len(s) / 2
	midAlt := int(math.Min(float64(mid+1), float64(len(s)-1)))
	//if midpoint is greater than v, return recurse on first half
	if s[mid] >= v {
		if len(s) == 2 {
			switch v {
			case s[0]:
				{
					return 0
				}
			case s[1]:
				{
					return 1
				}
			default:
				{
					return -1
				}
			}
		}
		return binarySearchWithDupes(s[:midAlt], v)
	}
	//if less, recurse on second half
	ret := binarySearchWithDupes(s[midAlt:], v)
	//if rec is -1, return -1
	if ret >= 0 {
		return ret + midAlt
	}
	return ret
}

//10.1: sorted Merge: Given 2 sorted arrays A and B, where A has a large enough buffer at the end to hold B, merge B into A.
//3:05~~3:49
func sortedMerge(a, b sortableSlice) sortableSlice {
	ret := sortableSlice{}
	indexA := 0
	indexB := 0
	for indexA < len(a) || indexB < len(b) {
		if indexA == len(a) {
			return append(ret, b[indexB:]...)
		}
		if indexB == len(b) {
			return append(ret, a[indexA:]...)
		}
		if a[indexA] < b[indexB] {

			ret = append(ret, a[indexA])
			indexA++
			continue
		}
		if a[indexA] > b[indexB] {
			ret = append(ret, b[indexB])
			indexB++
			continue
		}
		ret = append(ret, a[indexA], b[indexB])
		indexB++
		indexA++
	}
	return ret
}

//10.8 find duplicates: Given a array of integers from 1 to N, write an algorithm to generate all duplicates.
//2:58~~3:11
func findDupes(s sortableSlice) (ret sortableSlice) {
	m := make(map[int]int)
	for _, v := range s {
		if m[v] == 1 {
			ret = append(ret, v)
		}
		m[v]++
	}
	return ret
}
