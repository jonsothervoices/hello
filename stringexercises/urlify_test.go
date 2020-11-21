package stringexercises

import "testing"

func TestURLify(t *testing.T) {
	var tests = []struct {
		a        string
		expected string
	}{
		{"No lemons, no melon", "No%20lemons,%20no%20melon"},
		{"racecar", "racecar"},
		{"it is a truth universally accepted that a single man in possession of a fortune must be in want of a wife", "it%20is%20a%20truth%20universally%20accepted%20that%20a%20single%20man%20in%20possession%20of%20a%20fortune%20must%20be%20in%20want%20of%20a%20wife"},
		{"s", "s"},
		{"", ""},
		{"  ", "%20%20"},
		{" a ", "%20a%20"},
		{"%20", "%20"},
	}
	for i, datest := range tests {
		actual := urlify(datest.a)
		if datest.expected != actual {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
