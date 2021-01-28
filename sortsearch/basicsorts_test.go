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
		actual := datest.a.bubbleSort()
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, w := range actual {
			if w != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, w, datest.expected[j])
			}
		}
	}
}

func TestSelectionSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		actual := datest.a.selectionSort()
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

func TestInsertionSort(t *testing.T) {
	var tests = setup()
	for i, datest := range tests {
		actual := datest.a.insertionSort()
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

func setup() []sortTest {
	return []struct {
		a        sortableSlice
		expected sortableSlice
	}{
		{sortableSlice{3, 5, 2, 4, 1, 10, 7, 6, 9, 8}, sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{sortableSlice{2, 1}, sortableSlice{1, 2}},
		{sortableSlice{}, sortableSlice{}},
		{sortableSlice{0}, sortableSlice{0}},
		{sortableSlice{3, 2, 5, 3, 8, 1}, sortableSlice{1, 2, 3, 3, 5, 8}},
		{sortableSlice{3, 2, 5, 3, 8, 1, 12}, sortableSlice{1, 2, 3, 3, 5, 8, 12}},
	}
}
