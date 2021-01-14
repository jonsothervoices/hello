package recursion

import (
	"testing"
)

func TestTowerHan(t *testing.T) {
	var tests = []struct {
		a []int
	}{
		{[]int{1}},
		{[]int{2, 1}},
		{[]int{3, 2, 1}},
		{[]int{4, 3, 2, 1}},
		{[]int{8, 7, 6, 5, 4, 3, 2, 1}},
	}
	for i, datest := range tests {
		s := stack{}
		for _, v := range datest.a {
			s.push(v)
		}
		var actual stack
		towerHan(&s, &stack{}, &actual, s.len())
		if actual.len() != len(datest.a) {
			t.Errorf("%v: actual length %v, expected %v", i, actual.len(), len(datest.a))
			t.FailNow()
		}
		for j := len(datest.a) - 1; j >= 0; j-- {
			v := actual.pop()
			if v != datest.a[j] {
				t.Errorf("%v: actual %v, expected %v", i, v, datest.a[j])
			}
		}
	}
}
