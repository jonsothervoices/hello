package recursion

import (
	"fmt"
	"sort"
)

//8.2 Robot in a Gric: A robot starts in the upper left corner of a grid with r rows and c columns. The robot can move right and down, but certain cells are "off limits". Write a function to find a path from top left to bottom right.

//go right until you cant, then go down; if cant find path, try again
//
func robotGrid(m [][]bool) ([]string, error) {
	r := len(m)
	c := len(m[0])
	if r == 1 && c == 1 {
		return []string{}, nil
	}
	//do i have a right move?
	if c > 1 && m[0][1] {
		mSub := make([][]bool, r)
		for i := range m {
			mSub[i] = m[i][1:]
		}
		ret, err := robotGrid(mSub)
		//if return no error, propend/append string with the move
		if err == nil {
			return append([]string{"r"}, ret...), nil
		}
	}
	//do i have a down move?
	if r > 1 && m[1][0] {
		ret, err := robotGrid(m[1:][:])
		//if return no error, propend/append string with the move
		if err == nil {
			return append([]string{"d"}, ret...), nil
		}
	}
	return nil, fmt.Errorf("the cake is a lie")
}

//8.3 Magic Index: Given a sorted array A[0...n-1], write a function to find a magic index such that A[i]=i, if it exists.

func magicIndex(s []int) int {
	return recMagicIndex(s, 0, len(s)-1)
}

func recMagicIndex(s []int, start, end int) int {
	mid := (start + end) / 2
	if mid == s[mid] {
		return mid
	}
	if start >= end {
		return -1
	}
	if mid < s[mid] {
		return recMagicIndex(s, start, mid)
	}
	return recMagicIndex(s, mid+1, end)
}

//8.4: Power Set: Write a function to find the power set of a given set.
//find all subsets of size 1,2...n
//~~3:56

func mapDedupe(s []int) []int {
	newS := []int{}
	m := make(map[int]bool)
	for _, v := range s {
		if ok := m[v]; !ok {
			m[v] = true
			newS = append(newS, v)
		}
	}
	return newS
}

func sortDedupe(s []int) []int {
	newS := []int{}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	for i, v := range s {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		newS = append(newS, v)
	}
	return newS
}

//if s is empty, return
//remove item from s
//run powerset on new s
//loop through result, copying each slice and adding item to each
//return result
//add s to return

func powerSet(s []int) [][]int {
	return buildPowerSet(sortDedupe(s))
}

func buildPowerSet(s []int) [][]int {
	if len(s) == 0 {
		return [][]int{[]int{}}
	}
	ret := powerSet(s[1:])
	newS := [][]int{}
	for _, v := range ret {
		newV := append(v, s[0])
		sort.Slice(newV, func(i, j int) bool {
			return newV[i] < newV[j]
		})
		newS = append(newS, newV)
	}
	// ret = append(ret, s)
	ret = append(ret, newS...)
	return ret
}
