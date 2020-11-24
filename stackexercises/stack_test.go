package stackexercises

import "testing"

func TestPop(t *testing.T) {
	var tests = []struct {
		a        stringStack
		lenExp   int
		expected string
	}{
		{stringStack{[]string{"a", "b", "c", "d", "e"}}, 4, "e"},
		{stringStack{[]string{"a"}}, 0, "a"},
		{stringStack{[]string{""}}, 0, ""},
		{stringStack{[]string{}}, 0, ""},
	}
	for i, datest := range tests {
		actual := datest.a.pop()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if len(datest.a.data) != datest.lenExp {
			t.Errorf("%v: length %v, expected %v", i, len(datest.a.data), datest.lenExp)
		}
	}
}

func TestPush(t *testing.T) {
	var tests = []struct {
		a        stringStack
		b        string
		expected string
		lenExp   int
	}{
		{stringStack{[]string{"a", "b", "c", "d"}}, "e", "e", 5},
		{stringStack{[]string{"a"}}, "b", "b", 2},
		{stringStack{[]string{""}}, "", "", 2},
		{stringStack{[]string{}}, "", "", 1},
	}
	for i, datest := range tests {
		datest.a.push(datest.b)
		actual := datest.a.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestPeek(t *testing.T) {
	var tests = []struct {
		a        stringStack
		expected string
	}{
		{stringStack{[]string{"a", "b", "c", "d"}}, "d"},
		{stringStack{[]string{"a"}}, "a"},
		{stringStack{[]string{""}}, ""},
	}
	for i, datest := range tests {
		actual := datest.a.peek()
		actual2 := datest.a.peek()
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
		if actual != actual2 {
			t.Errorf("%v: stack altered from %v to %v", i, actual, actual2)
		}
	}
}
