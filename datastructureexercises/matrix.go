package datastructureexercises

import "fmt"

//Given an NxN matrix, where each "pixel" is 4 bytes, write a method to rotate the image by 90 degrees (in place).

type matrix [][]string

// m:[[aa] [bb]], [[ba] [ba]]
// l: 1
// i: 0
// a: [aa]
// j: 0

func (m matrix) rotate() {
	fmt.Println(m)
	l := (len((m)) - 1) / 2
	for i := 0; i <= l; i++ {
		for j := 0; j < len(m); j++ {
			// x := m[i][j]
			// y := m[l-i][l-j]
			// fmt.Printf("%v %v \n", x, y)
			m[i][j], m[l-i][l-j] = m[l-i][l-j], m[i][j]
			// m[l-i][l-j] = x
			// m[i][j] = y
		}
	}
	fmt.Println(m)
}
