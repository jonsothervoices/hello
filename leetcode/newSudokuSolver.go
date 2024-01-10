package leetcode

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

type sBoard struct {
	Data   [][]byte
	AuxMap [][]map[byte]bool
}

func solveSudoku(board [][]byte) {
	myBoard := newSBoard(board)
	fmt.Printf("%v", doSolveSudoku(myBoard))
}

// Log logs, just a fancy printer
func Log(m string, ms ...interface{}) {
	pn := "sudoku"
	// runtime.Caller(1): Gimme the guy right under me on the call Stack
	// AKA the one that called me
	_, path, line, _ := runtime.Caller(1)
	pkg, file := filepath.Split(path)
	if dirs := strings.Split(strings.TrimPrefix(strings.TrimSuffix(pkg, `\`), `\`), `\`); len(dirs) > 0 {
		pkg = dirs[len(dirs)-1]
	}
	fmt.Printf("%v - %s/%s:%d %s\n", pn, pkg, file, line, fmt.Sprintf(m, ms...))
}

// The big daddy, the main driver
func doSolveSudoku(board sBoard) bool {
	if !board.isValid() {
		return false
	}
	fmt.Println("New Stack Frame (guess) with New Board: ")
	fmt.Println(board.Data)
	//solve all solvable squares
	for !board.isSolved() {
		if !board.solveEasySquares() {
			fmt.Println("Solve failed, running deduce")
			fmt.Println(board.Data)
			if !board.deduceEasySquares() {
				fmt.Println("Deduce failed, running guess")

				// Guess returns error when there are no 2-solvable squares
				row, col, guess := board.guess()
				if guess == 0 {
					return false
				}
				boardCopy := copySBoard(&board)
				fmt.Printf("Guessing %v at %v, %v\n", guess, row, col)
				fmt.Println("AuxMap at location:")
				fmt.Println(boardCopy.AuxMap[row][col])
				// probably updates the auxMap
				boardCopy.set(row, col, guess)
				fmt.Println(boardCopy.Data)
				// recurse!
				if try := doSolveSudoku(boardCopy); try {
					board.copy(boardCopy)
					return true
				}
				fmt.Println()
				board.crossOut(row, col, guess)
				// board.eliminate(row, col, guess)
			}
		}
		// fmt.Println(board.Data)
	}
	return true
}

// Just eliminates one possibility, used for guessing
func (board *sBoard) crossOut(row, col int, v byte) {
	board.AuxMap[row][col][v] = false
}

// isValid() checks the auxMap for a 0-possible square
// if it finds one return false, else true
func (board *sBoard) isValid() bool {
	for _, row := range board.AuxMap {
		for _, square := range row {
			foundATrue := false
			for k, v := range square {
				if k == '0' {
					continue
				}
				if v {
					foundATrue = true
					break
				}
			}
			if !foundATrue {
				fmt.Println("invalid board")
				return false
			}
		}
	}
	fmt.Println("valid board")
	return true
}

// isSolved() checks the board for any unsolved squares
// if it finds one return false, else true
func (board *sBoard) isSolved() bool {
	for _, v := range board.Data {
		for _, u := range v {
			if u == '.' {
				return false
			}
		}
	}
	return true
}

// solveEasySquares() checks the auxMap for a 1-possible square
// solves every one it finds
func (board *sBoard) solveEasySquares() (ret bool) {
	ret = false
	for i, row := range board.AuxMap {
		for j, square := range row {
			if board.Data[i][j] != '.' {
				continue
			}
			foundOne := false
			var theVal byte
			for k, v := range square {
				if k == '0' {
					continue
				}
				if !v {
					continue
				}
				if foundOne {
					foundOne = false
					break
				}
				foundOne = true
				theVal = k
			}
			if foundOne {
				board.set(i, j, theVal)
				ret = true
			}
		}
	}
	return
}

// deduceEasySquares() checks the auxMap for any squares that are uniquely possible
// these squares are the only in their row, col, or block that have some value still possible
// if we find one, set it
func (board *sBoard) deduceEasySquares() (ret bool) {
	ret = false
	for i, row := range board.AuxMap {
		for j, square := range row {
			if board.Data[i][j] != '.' {
				continue
			}
			for val := range square {
				if val == '0' {
					continue
				}
				if board.deduce(i, j, val) {
					board.set(i, j, val)
					ret = true
					break
				}
			}
		}
	}
	return
}

func (board *sBoard) deduce(row, col int, val byte) bool {
	return board.isUniqueInRow(row, col, val) || board.isUniqueInCol(row, col, val) || board.isUniqueInBlock(row, col, val)
}

// isUniqueInRow returns true if there is no possible other place in the row for the number val to be
//besides, importantly, board.Data[row][col]
func (board *sBoard) isUniqueInRow(row, col int, val byte) bool {
	for i, square := range board.AuxMap[row] {
		if i != col && square[val] {
			return false
		}
	}
	return true
}

// isUniqueInCol returns true if there is no possible other place in the column for the number n to be
//besides, importantly, board[row][col]
func (board *sBoard) isUniqueInCol(row, col int, val byte) bool {
	for i, r := range board.AuxMap {
		if i == row {
			continue
		}
		for j, square := range r {
			if j != col {
				continue
			}
			if square[val] {
				return false
			}
		}
	}
	return true
}

//isUniqueInBlock returns true if there is no possible other place in the block for the number n to be
//besides, importantly, board[row][col]
func (board *sBoard) isUniqueInBlock(row, col int, val byte) bool {
	// first things first, which block are we in?
	var theRow int
	var theCol int
	if row < 3 {
		theRow = 0
	} else if row > 5 {
		theRow = 2
	} else {
		theRow = 1
	}

	if col < 3 {
		theCol = 0
	} else if col > 5 {
		theCol = 2
	} else {
		theCol = 1
	}

	// Now, build the loop bounds
	// NOTE: High bounds are exclusive!
	var rLow int
	var rHigh int
	var cLow int
	var cHigh int
	switch theRow {
	case 0:
		rLow = 0
		rHigh = 3
	case 1:
		rLow = 3
		rHigh = 6
	case 2:
		rLow = 6
		rHigh = 9
	}
	switch theCol {
	case 0:
		cLow = 0
		cHigh = 3
	case 1:
		cLow = 3
		cHigh = 6
	case 2:
		cLow = 6
		cHigh = 9
	}

	// finally, the loop
	// we MUST avoid the box in question,
	// so we'll need a new bool every j loop
	notOurBox := false
	for i := rLow; i < rHigh; i++ {
		for j := cLow; j < cHigh; j++ {
			notOurBox = i != row || j != col
			if notOurBox && board.AuxMap[i][j][val] {
				return false
			}
		}
	}
	return true
}

// guess() searches the auxMap for the most highly eliminated square, keeping the first one it finds
// it then picks the first possibility for this guess
// returns the coordinates and one possibility
func (board *sBoard) guess() (row, col int, guess byte) {
	bestPoss := 9
	for i, r := range board.AuxMap {
		for j, square := range r {
			if board.Data[i][j] != '.' {
				continue
			}
			possCount := 0
			for k, v := range square {
				if k != '0' && v {
					guess = k
					possCount++
				}
			}
			if possCount == 2 {
				return i, j, guess
			}
			if possCount < bestPoss {
				row = i
				col = j
				bestPoss = possCount
			}
		}
	}
	return
}

// overloaded assignment
// ie. copy constructor that returns a new board
func copySBoard(a *sBoard) (b sBoard) {
	aJSON, _ := json.Marshal((*a).Data)
	json.Unmarshal(aJSON, &b.Data)
	aMapJSON, _ := json.Marshal((*a).AuxMap)
	json.Unmarshal(aMapJSON, &b.AuxMap)
	return
}

// Board constructor
//Builds and updates auxMap
func newSBoard(data [][]byte) (b sBoard) {
	b.AuxMap = make([][]map[byte]bool, 9)
	for i := 0; i < 9; i++ {
		b.AuxMap[i] = make([]map[byte]bool, 9)
		for j := 0; j < 9; j++ {
			b.AuxMap[i][j] = map[byte]bool{
				'0': false, '1': true, '2': true, '3': true, '4': true, '5': true, '6': true, '7': true, '8': true, '9': true}
		}
	}
	b.Data = data
	b.update()
	return
}

// copy constructor
func (board *sBoard) copy(b sBoard) {
	aJSON, _ := json.Marshal(b.Data)
	json.Unmarshal(aJSON, &board.Data)
	aMapJSON, _ := json.Marshal(b.AuxMap)
	json.Unmarshal(aMapJSON, &board.AuxMap)
}

// eliminate() updates the auxMap
// given a newly solved square, it updates the associated possibilities on the relevant row, column, and block of the auxMap
func (board *sBoard) eliminate(row, col int, v byte) {

	// eliminate the row
	for _, square := range board.AuxMap[row] {
		square[v] = false
	}

	// eliminate col
	for _, r := range board.AuxMap {
		r[col][v] = false
	}

	// eliminate block
	// first things first, which block are we in?
	var theRow int
	var theCol int

	if row < 3 {
		theRow = 0
	} else if row > 5 {
		theRow = 2
	} else {
		theRow = 1
	}

	if col < 3 {
		theCol = 0
	} else if col > 5 {
		theCol = 2
	} else {
		theCol = 1
	}

	// Now, build the loop bounds
	// NOTE: High bounds are exclusive!
	var rLow int
	var rHigh int
	var cLow int
	var cHigh int
	switch theRow {
	case 0:
		rLow = 0
		rHigh = 3
	case 1:
		rLow = 3
		rHigh = 6
	case 2:
		rLow = 6
		rHigh = 9
	}
	switch theCol {
	case 0:
		cLow = 0
		cHigh = 3
	case 1:
		cLow = 3
		cHigh = 6
	case 2:
		cLow = 6
		cHigh = 9
	}

	// finally, the loop
	for i := rLow; i < rHigh; i++ {
		for j := cLow; j < cHigh; j++ {
			board.AuxMap[i][j][v] = false
		}
	}
	// finally, set our square to true at our val
	board.AuxMap[row][col][v] = true
}

// Set() sets
func (board *sBoard) set(row, col int, v byte) {
	board.Data[row][col] = v
	board.eliminate(row, col, v)
}

func (board *sBoard) update() {
	for i, v := range board.Data {
		for j, u := range v {
			if u != '.' {
				//first, update auxMap
				board.AuxMap[i][j][0] = true
				//now, eliminate u from auxMap
				board.eliminate(i, j, u)
			}
		}
	}
}
