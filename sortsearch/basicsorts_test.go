package sortsearch

import (
	"testing"
)

type sortTest = struct {
	a        sortableSlice
	expected sortableSlice
}

func TestBubbleSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		datest.a.bubbleSort()
		if len(datest.a) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(datest.a), len(datest.expected))
			t.FailNow()
		}
		for j, w := range datest.a {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		datest.a.selectionSort()
		if len(datest.a) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(datest.a), len(datest.expected))
		}
		for j, w := range datest.a {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestInsertionSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		datest.a.insertionSort()
		if len(datest.a) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(datest.a), len(datest.expected))
		}
		for j, w := range datest.a {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestMergeSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		actual := mergeSort(datest.a)
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(actual), len(datest.expected))
		}
		for j, w := range actual {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestQuickSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		datest.a.quickSort()
		if len(datest.a) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(datest.a), len(datest.expected))
		}
		for j, w := range datest.a {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestRadixSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		datest.a.radixSort()
		if len(datest.a) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(datest.a), len(datest.expected))
		}
		for j, w := range datest.a {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestGetDigit(t *testing.T) {
	var tests = []struct {
		a        int
		digit    int
		expected int
	}{
		{3452, 1, 2},
		{3452, 2, 5},
		{3452, 3, 4},
		{3452, 4, 3},
	}
	for i, datest := range tests {
		actual := getDigit(datest.a, datest.digit, 1)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func setup() []sortTest {
	return []struct {
		a        sortableSlice
		expected sortableSlice
	}{
		{sortableSlice{3452, 5234, 5462, 262, 4, 1, 1000, 75, 6417, 999, 8}, sortableSlice{1, 4, 8, 75, 262, 999, 1000, 3452, 5234, 5462, 6417}},
		{sortableSlice{732, 23, 1, 55, 7130, 321, 223, 5}, sortableSlice{1, 5, 23, 55, 223, 321, 732, 7130}},
		{sortableSlice{3, 5, 2, 4, 1, 10, 7, 6, 9, 8}, sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{sortableSlice{2, 1}, sortableSlice{1, 2}},
		{sortableSlice{}, sortableSlice{}},
		{sortableSlice{0}, sortableSlice{0}},
		{sortableSlice{3, 2, 5, 3, 8, 1}, sortableSlice{1, 2, 3, 3, 5, 8}},
		{sortableSlice{3, 2, 5, 3, 8, 1, 12}, sortableSlice{1, 2, 3, 3, 5, 8, 12}},
	}
}
