package recursion

import (
	"fmt"
	"math"
	"sort"
)

type stack struct {
	data []interface{}
}

func (s *stack) len() int {
	return len(s.data)
}

func (s *stack) push(a interface{}) *stack {
	s.data = append(s.data, a)
	return s
}

func (s *stack) peek() interface{} {
	if s.len() == 0 {
		return nil
	}
	return s.data[s.len()-1]
}

func (s *stack) pop() interface{} {
	ret := s.peek()
	if ret == nil {
		return nil
	}
	s.data = s.data[:s.len()-1]
	return ret
}

//8.2 Robot in a Grid: A robot starts in the upper left corner of a grid with r rows and c columns. The robot can move right and down, but certain cells are "off limits". Write a function to find a path from top left to bottom right.

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

func strDedupe(s []string) []string {
	newS := []string{}
	m := make(map[string]bool)
	for _, v := range s {
		if ok := m[v]; !ok {
			m[v] = true
			newS = append(newS, v)
		}
	}
	return newS
}

func runeDedupe(s string) string {
	var newS string
	m := make(map[rune]bool)
	for _, v := range s {
		if ok := m[v]; !ok {
			m[v] = true
			newS = newS + string(v)
		}
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

//8.5 Recursive muliply: write a recursive function to multiply 2 positive integers.
//~~3:01-3:04
//optimization~~3:04-3:11

func recMultiply(a, b uint) uint {
	if b <= a {
		return doRecMultiply(a, b)
	}
	return doRecMultiply(b, a)
}

func doRecMultiply(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	if b == 1 {
		return a
	}
	return a + recMultiply(a, b-1)
}

//8.6: Towers of Hanoi
//~~
func towerHan(s, buff, dest *stack, n int) {
	if n == 0 {
		return
	}
	towerHan(s, dest, buff, n-1)
	dest.push(s.pop())
	towerHan(buff, s, dest, n-1)
}

//8.7: Premutaions without dups: Write a function to find all the permutations of a string with unique characters

func uniquePerm(s string) []string {
	if len(s) == 1 {
		return []string{s}
	}
	if len(s) == 0 {
		return []string{""}
	}
	source := uniquePerm(s[1:])
	dest := []string{}
	for _, v := range source {
		for j := range v {
			dest = append(dest, fmt.Sprintf("%v%v%v", v[:j], string(s[0]), v[j:]))
		}
		dest = append(dest, fmt.Sprintf("%v%v", v, string(s[0])))
	}
	return dest
}

//8.8: Permutations with dups:
//1:22~~
func nonUniquePerm(s string) []string {
	// m:= make (map[rune]int)
	// for _,v:= range s{
	// 		m[v]=m[v]+1
	// }
	return strDedupe(uniquePerm(s))
}

//COME BACK?
// func runNonUniquePerm(m map[rune]int) []string {
// 	if len(m) == 1 {
// 		for k,v:= range m{
// 			return []string{strings.Repeat(string(k), v)}
// 		}
// 	}
// 	if len(m) == 0 {
// 		return []string{""}
// 	}
// 	var rKey rune
// 	var rVal int
// 	for k,v := range m {
// 		rKey,rVal=k,v
// 		delete(m,k)
// 		break
// 	}
// 	for i:=0;i<=rKey;i++{
//
// 	}
// 	source := runNonUniquePerm(m)
// 	dest := []string{}
// 	for k, v := range source {
// 		for j := range v {
// 			dest = append(dest, fmt.Sprintf("%v%v%v", v[:j], string(rKey), v[j:]))
// 		}
// 		dest = append(dest, fmt.Sprintf("%v%v", v, string(rKey)))
// 	}
// 	return dest
// }

//8.9ish: Parens: write a function to check that a given string has valid parentheses

func parens(s string) bool {
	var stack int
	for _, v := range s {
		if v == '(' {
			stack++
		}
		if v == ')' {
			stack--
		}
		if stack < 0 {
			return false
		}
	}
	return stack == 0
}

//8.10: Paint fill: given a 2D array of colors, Implement the "paint fill" function.
//3:12~~4:00
func paintFill(image [][]string, c string, y, x int) [][]string {
	//range check
	if y > len(image)-1 {
		return image
	}
	if x > len(image[y])-1 {
		return image
	}
	//dupe color check
	if image[y][x] == c {
		return image
	}
	//identify origin color
	orig := image[y][x]
	//color origin
	mImage := image
	mImage[y][x] = c
	//check for valid left
	if x > 0 && mImage[y][x-1] == orig {
		mImage = paintFill(mImage, c, y, x-1)
	}
	//check for valid right
	if x < len(mImage[y])-1 && mImage[y][x+1] == orig {
		mImage = paintFill(mImage, c, y, x+1)
	}
	//check for valid up
	if y > 0 && mImage[y-1][x] == orig {
		mImage = paintFill(mImage, c, y-1, x)
	}
	//check for valid down
	if y < len(mImage)-1 && mImage[y+1][x] == orig {
		mImage = paintFill(mImage, c, y+1, x)
	}
	return mImage
}

//8.11: Coins: given infinite coins (quarters, dimes, nickels, and pennies) calculate the number of ways of representing n cents.
//3:01~~5:15
//TODO: COME BACK???
func coins(a int) int {
	if a == 0 {
		return 0
	}
	return doCoins(a, 25)
}

func doCoins(a, d int) int {
	if d == 1 {
		return 1
	}
	ret := 0
	for i := 0; i*d <= a; i++ {
		rem := a - (i * d)
		ret += doCoins(rem, nextCoin(d))
	}
	return ret
}

func nextCoin(d int) int {
	if d == 25 {
		return 10
	}
	if d == 10 {
		return 5
	}
	return 1
}

//8.12: Eight Queens: Write a function to print all the ways of arranging 8 queens on an 8x8 chess board such that none share the same row, column, or diagonal.
type board [8]uint8

func (b board) isAllowed(r, c int) bool {
	//check row
	if b[r] > 0 {
		return false
	}
	//creat column mask
	cMask := uint8(128) >> c
	//column check
	for _, v := range b {
		if cMask&v != 0 {
			return false
		}
	}
	//check diags
	for i, v := range b {
		shift := int(math.Abs(float64(r - i)))
		if cMask<<shift&v > 0 || cMask>>shift&v > 0 {
			return false
		}
	}
	return true
}

func (b board) sum() uint8 {
	var s uint8
	for _, v := range b {
		s += v
	}
	return s
}

func eightQueens(b board, row int) (ret []board) {
	if b.sum() == 255 {
		return []board{b}
	}
	m := make(map[string]bool)
	for i, u := range b {
		if u != 0 {
			continue
		}
		bScratch := b
		if ok := bScratch.isAllowed(i, 7-row); !ok {
			continue
		}
		bScratch[i] = (1 << row) | bScratch[i]
		current := eightQueens(bScratch, row+1)
		for _, v := range current {
			if _, ok := m[fmt.Sprintf("%v", v)]; !ok {
				m[fmt.Sprintf("%v", v)] = true
				ret = append(ret, v)
			}
		}
	}
	return
}
