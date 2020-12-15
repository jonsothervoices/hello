package graphexercises

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestNewBST(t *testing.T) {
	var tests = []struct {
		d        []int
		expected string
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7}, `{"Data":4,"Left":{"Data":2,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null},"Right":{"Data":3,"Left":null,"Right":null}},"Right":{"Data":6,"Left":{"Data":5,"Left":null,"Right":null},"Right":{"Data":7,"Left":null,"Right":null}}}`},
		{[]int{0, 1, 2, 3, 4, 5, 6}, `{"Data":3,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}},"Right":{"Data":5,"Left":{"Data":4,"Left":null,"Right":null},"Right":{"Data":6,"Left":null,"Right":null}}}`},
		{[]int{0, 1, 2}, `{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}}`},
		{[]int{0, 1}, `{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null}`},
		{[]int{0}, `{"Data":0,"Left":null,"Right":null}`},
		{[]int{}, "null"},
	}
	for i, datest := range tests {
		g := newBST(datest.d, nil)
		actual, err := json.Marshal(g)
		if err != nil && datest.expected != "null" {
			fmt.Println(err)
		}
		if datest.expected != string(actual) {
			t.Errorf("%v: actual %v, expected %v", i, string(actual), datest.expected)
		}
	}
}
