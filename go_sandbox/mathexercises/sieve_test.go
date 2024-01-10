package mathexercises

import "testing"

func TestSieveOfEratosthenes(t *testing.T) {
	var tests = []struct {
		max      int
		expected []int
	}{
		{30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{10, []int{2, 3, 5, 7}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for i, datest := range tests {
		actual := SieveOfEratosthenes(datest.max)
		if len(actual) != len(datest.expected) {
			t.Errorf("%v: actual length %v, expected length %v", i, len(actual), len(datest.expected))
			t.FailNow()
		}
		for j, v := range actual {
			if v != datest.expected[j] {
				t.Errorf("%v: actual %v, expected %v", i, v, datest.expected[j])
			}
		}
	}
}
