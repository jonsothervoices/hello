package datastructureexercises

//Given an NxN matrix, where each "pixel" is 4 bytes, write a method to rotate the image by 90 degrees (in place).

type matrix [][]string

//rotate has O(n^2)
func (m matrix) rotate() {
	max := len(m) - 1
	for y := 0; y <= max/2; y++ {
		for x := y; x < max-y; x++ {
			m[x][max-y], m[max-y][max-x], m[max-x][y], m[y][x] =
				m[y][x], m[x][max-y], m[max-y][max-x], m[max-x][y]
		}
	}
}
