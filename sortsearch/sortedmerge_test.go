package sortsearch

import (
	"testing"
)

func TestSortedMerge(t *testing.T) {
	var tests = []struct {
		a        sortableSlice
		b        sortableSlice
		expected sortableSlice
	}{
		{sortableSlice{1, 3}, sortableSlice{2, 4, 6}, sortableSlice{1, 2, 3, 4, 6}},
		{sortableSlice{1, 3, 5, 7, 9, 11}, sortableSlice{2, 4, 6, 8, 10, 12}, sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}},
		{sortableSlice{1, 3, 3, 7, 9, 11}, sortableSlice{2, 4, 6, 8, 8, 12}, sortableSlice{1, 2, 3, 3, 4, 6, 7, 8, 8, 9, 11, 12}},
		{sortableSlice{1, 5}, sortableSlice{2, 4, 6, 8, 10, 12}, sortableSlice{1, 2, 4, 5, 6, 8, 10, 12}},
		{sortableSlice{1, 3, 5, 7, 9, 11}, sortableSlice{2, 4, 6, 8, 9, 12}, sortableSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 9, 11, 12}},
		{sortableSlice{}, sortableSlice{2, 4, 6, 8, 10, 12}, sortableSlice{2, 4, 6, 8, 10, 12}},
		{sortableSlice{1, 3, 5, 9, 11, 70}, sortableSlice{2, 4, 6, 8, 12, 1000}, sortableSlice{1, 2, 3, 4, 5, 6, 8, 9, 11, 12, 70, 1000}},
	}
	for i, datest := range tests {
		actual := sortedMerge(datest.a, datest.b)
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
