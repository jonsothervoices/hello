package recursion

import "fmt"

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
