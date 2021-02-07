package sortsearch

import (
	"testing"
)

func TestFindDupes(t *testing.T) {
	var tests = []struct {
		s        sortableSlice
		expected sortableSlice
	}{
		{sortableSlice{5, 6, 11, 8, 9, 10, 8, 11, 8, 2, 3, 4}, sortableSlice{8, 11}},
		{sortableSlice{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 4}, sortableSlice{5}},
		{sortableSlice{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 5}, sortableSlice{5}},
		{sortableSlice{5, 6, 8, 3, 1, 4}, sortableSlice{}},
		{sortableSlice{5, 5, 5, 5, 5}, sortableSlice{5}},
		{sortableSlice{}, sortableSlice{}},
	}
	for i, datest := range tests {
		actual := findDupes(datest.s)
		actual.radixSort()
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected length %v", i, len(actual), len(datest.expected))
		}
		for j, v := range actual {
			if v != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, v, datest.expected[j])
			}
		}
	}
}
