package mathexercises

import (
	"fmt"
	"testing"
	"time"
)

func TestUncacheFib(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{5, 3},
		{10, 34},
		{12, 89},
		{1, 0},
		{2, 1},
		{30, 514229},
		// {51, 12586269025},
	}
	fmt.Println("running uncached")
	for i, datest := range tests {
		tm := time.Now()
		actual := uncacheNthFib(datest.n)
		fmt.Println("took", int64(time.Since(tm)))
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestCacheFib(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{5, 3},
		{10, 34},
		{12, 89},
		{1, 0},
		{2, 1},
		//{30, 514229},
		{51, 12586269025},
		{91, 2880067194370816120},
	}
	fmt.Println("running cached")
	for i, datest := range tests {
		tm := time.Now()
		actual := cacheNthFib(datest.n, nil)
		fmt.Println("took", int64(time.Since(tm)))
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}

func TestItFib(t *testing.T) {
	var tests = []struct {
		n        int
		expected int
	}{
		{5, 3},
		{10, 34},
		{12, 89},
		{1, 0},
		{2, 1},
		{30, 514229},
		{51, 12586269025},
		{91, 2880067194370816120},
	}
	fmt.Println("running iterated")
	for i, datest := range tests {
		tm := time.Now()
		actual := itNthFib(datest.n)
		fmt.Println("took", int64(time.Since(tm)))
		if actual != datest.expected {
			t.Errorf("%v: actual %v, expected %v", i, actual, datest.expected)
		}
	}
}
