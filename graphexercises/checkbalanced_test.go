package graphexercises

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var tests = []struct {
		d        string
		depth    int
		expected bool
	}{
		{`{"Data":4,"Left":{"Data":2,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null},"Right":{"Data":3,"Left":null,"Right":null}},"Right":{"Data":6,"Left":{"Data":5,"Left":null,"Right":null},"Right":{"Data":7,"Left":null,"Right":null}}}`, 4, true},
		{`{"Data":3,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}},"Right":{"Data":5,"Left":{"Data":4,"Left":null,"Right":null},"Right":{"Data":6,"Left":null,"Right":null}}}`, 3, true},
		{`{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}}`, 2, true},
		{`{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null}`, 2, true},
		{`{"Data":0,"Left":null,"Right":null}`, 1, true},
		{"null", 1, true},
		{`{"Data":4,"Left":{"Data":2,"Left":{"Data":1,"Left":{"Data":0,"Left":{"Data":3,"Left":{"Data":3,"Left":null,"Right":null},"Right":null},"Right":null},"Right":null},"Right":{"Data":3,"Left":null,"Right":null}},"Right":{"Data":3,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}},"Right":{"Data":5,"Left":{"Data":4,"Left":null,"Right":null},"Right":{"Data":6,"Left":null,"Right":null}}}}`, 6, false},
		{`{"Data":1,"Left":{"Data":0,"Left":{"Data":0,"Left":{"Data":0,"Left":null,"Right":null},"Right":null},"Right":{"Data":0,"Left":null,"Right":null}},"Right":{"Data":0,"Left":null,"Right":null}}`, 4, false},
	}
	for i, datest := range tests {
		var actual bst
		//actual:= bst{}
		err := json.Unmarshal([]byte(datest.d), &actual)
		if err != nil {
			fmt.Println(err)
		}
		b, dActual := actual.isBalanced()
		if datest.expected != b {
			t.Errorf("%v: actual %v, expected %v", i, b, datest.expected)
		}
		if datest.expected && datest.depth != dActual {
			t.Errorf("%v: actual %v, expected %v", i, dActual, datest.depth)
		}
	}
}
