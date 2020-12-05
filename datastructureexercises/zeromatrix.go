package datastructureexercises

type intmatrix [][]int

//Write a program such that if an element of an MxN matrix is zero, it's entire row and column are set to 0.

//ZeroMatrix is O(mn)
func (m intmatrix) ZeroMatrix() {
	rowsWithZeros := make(map[int]bool)
	columnswithZeros := make(map[int]bool)
	for y, r := range m {
		for x, c := range r {
			if c == 0 {
				columnswithZeros[x] = true
				rowsWithZeros[y] = true
			}
		}
	}
	for y, r := range m {
		if rowsWithZeros[y] {
			m[y] = make([]int, len(r))
			continue
		}
		for x := range r {
			if columnswithZeros[x] {
				m[y][x] = 0
			}
		}
	}
}
