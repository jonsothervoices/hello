package sortsearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	var tests = []struct {
		s        sortableSlice
		a        int
		expected int
	}{
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 10, 9},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, 2},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 1, 0},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12, 11},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 7, 6},
		// {sortableSlice{1, 3, 3, 7, 9, 11}, 25, -1},
		{sortableSlice{1}, 1, 0},
		{sortableSlice{}, 3, -1},
		// {sortableSlice{2, 4, 6, 8, 10, 10, 12}, 10, 4},
		// {sortableSlice{2, 4, 6, 10, 10, 10, 12}, 10, 3},
		// {sortableSlice{10, 10, 10, 10, 10, 10, 10}, 10, 0},
		// {sortableSlice{10, 10, 10, 10, 10, 10}, 10, 0},
	}
	for i, datest := range tests {
		actual := binarySearch(datest.s, datest.a)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestBinarySearchWithDupes(t *testing.T) {
	var tests = []struct {
		s        sortableSlice
		a        int
		expected int
	}{
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 10, 9},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 3, 2},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 1, 0},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 12, 11},
		{sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 7, 6},
		{sortableSlice{1, 3, 3, 7, 9, 11}, 25, -1},
		{sortableSlice{1}, 1, 0},
		{sortableSlice{}, 3, -1},
		{sortableSlice{2, 4, 6, 8, 10, 10, 12}, 10, 4},
		{sortableSlice{2, 4, 6, 10, 10, 10, 12}, 10, 3},
		{sortableSlice{10, 10, 10, 10, 10, 10, 10}, 10, 0},
		{sortableSlice{10, 10, 10, 10, 10, 10}, 10, 0},
	}
	for i, datest := range tests {
		actual := binarySearchWithDupes(datest.s, datest.a)
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
