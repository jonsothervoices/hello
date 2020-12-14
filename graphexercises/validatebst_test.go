package graphexercises

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		d        string
		expected bool
	}{
		{`{"Data":4,"Left":{"Data":2,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null},"Right":{"Data":3,"Left":null,"Right":null}},"Right":{"Data":6,"Left":{"Data":5,"Left":null,"Right":null},"Right":{"Data":7,"Left":null,"Right":null}}}`, true},
		{`{"Data":3,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}},"Right":{"Data":5,"Left":{"Data":4,"Left":null,"Right":null},"Right":{"Data":6,"Left":null,"Right":null}}}`, true},
		{`{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":{"Data":2,"Left":null,"Right":null}}`, true},
		{`{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null}`, true},
		{`{"Data":0,"Left":null,"Right":null}`, true},
		{`{"Data":5}`, true},
		{`{"Data":4,"Right":{"Data":6,"Right":{"Data":7,"Right":null}}}`, true},
		{`{"Data":4,"Right":{"Data":7,"Right":{"Data":6,"Right":null}}}`, false},
		{`{"Data":4,"Left":{"Data":2,"Left":{"Data":1,"Left":{"Data":0,"Left":null,"Right":null},"Right":null},"Right":{"Data":3,"Left":null,"Right":null}},"Right":{"Data":5,"Left":{"Data":6,"Left":null,"Right":null},"Right":{"Data":7,"Left":null,"Right":null}}}`, false},
	}
	for i, datest := range tests {
		var actual bst
		err := json.Unmarshal([]byte(datest.d), &actual)
		if err != nil {
			fmt.Println(err)
		}
		b := actual.isValid()
		if datest.expected != b {
			t.Errorf("%v: actual %v, expected %v", i, b, datest.expected)
		}
	}
}
