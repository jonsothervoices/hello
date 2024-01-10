package sudoku

//
// import (
// 	"encoding/json"
// 	"fmt"
// 	"path/filepath"
// 	"runtime"
// 	"strings"
// )
//
// //Log logs
// func Log(m string, ms ...interface{}) {
// 	pn := "sudoku"
// 	// runtime.Caller(1): Gimme the guy right under me on the call Stack
// 	// AKA the one that called me
// 	_, path, line, _ := runtime.Caller(1)
// 	pkg, file := filepath.Split(path)
// 	if dirs := strings.Split(strings.TrimPrefix(strings.TrimSuffix(pkg, `\`), `\`), `\`); len(dirs) > 0 {
// 		pkg = dirs[len(dirs)-1]
// 	}
// 	fmt.Printf("%v - %s/%s:%d %s", pn, pkg, file, line, fmt.Sprintf(m, ms...))
// }
//
// type truthMap map[int]bool
//
// func newTruthMap() truthMap {
// 	return truthMap{
// 		0: false,
// 		1: true,
// 		2: true,
// 		3: true,
// 		4: true,
// 		5: true,
// 		6: true,
// 		7: true,
// 		8: true,
// 		9: true,
// 	}
// }
//
// //truthboard type: a 9x9 grid of truthmaps
// type truthBoard [9][9]truthMap
//
// func newBoard() truthBoard {
// 	return truthBoard{
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 		{newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap(), newTruthMap()},
// 	}
// }
//
// func solveSudoku(board [][]byte) {
// 	board = doSolveSudoku(board)
// }
//
// func boardCopier(board [][]byte) (ret [][]byte) {
// 	jsonBoard, _ := json.Marshal(board)
// 	json.Unmarshal(jsonBoard, &ret)
// 	return
// }
//
// func doSolveSudoku(board [][]byte) [][]byte {
// 	infoBoard := newBoard()
// 	updateAll(&board, &infoBoard)
// 	if foundContradiction(&infoBoard) {
// 		return nil
// 	}
// 	done := isSolved(&board)
// 	for !done {
// 		data := findAndSolve(&board, &infoBoard)
// 		// findAContra here?
// 		// updateAll(&board, &infoBoard)
// 		if data[0] != -1 {
// 			infoBoard.eliminateRow(data[0], data[2])
// 			infoBoard.eliminateCol(data[1], data[2])
// 			infoBoard.eliminateBlock(data[0], data[1], data[2])
// 		} else {
// 			// GUESS AND START BACKTRACKING
//
// 			// guesser:
// 			// Loops through all 2-possible squares
// 			// Ends when a CONTRADICTION is found
// 			// find a 2-possible square
// 			breakTime := false
// 			// weNeedThrees := false
// 			for i, v := range infoBoard {
// 				for j, u := range v {
// 					if theGuesses := twoPossible(u); theGuesses != nil {
// 						// Recurse!
// 						// if recursion returns nil, try the second guess
// 						// If neither works, reset the square and go to the second 2-possible
// 						// APPARENTLY this is enough, youll never have to go to the 3-possible squares
// 						newBoard := boardCopier(board)
// 						newBoard[i][j] = byte('0' + theGuesses[0])
// 						if try := doSolveSudoku(newBoard); try != nil {
// 							fmt.Println("successful guess")
// 							fmt.Println(newBoard)
// 							fmt.Println(try)
// 							return try
// 						}
// 						board[i][j] = byte('0' + theGuesses[1])
// 						infoBoard.eliminateRow(i, theGuesses[1])
// 						infoBoard.eliminateCol(j, theGuesses[1])
// 						infoBoard.eliminateBlock(i, j, theGuesses[1])
// 						breakTime = true
// 						break
// 						// close if 2 possible
// 					}
// 					// close inner loop
// 				}
// 				if breakTime == true {
// 					break
// 				}
// 				// close outer loop
// 			}
// 			if !breakTime {
// 				for i, v := range infoBoard {
// 					for j, u := range v {
// 						if theGuesses := threePossible(u); theGuesses != nil {
// 							// Recurse!
// 							// if recursion returns nil, try the second guess
// 							// If neither works, reset the square and go to the second 2-possible
// 							// APPARENTLY this is enough, youll never have to go to the 3-possible squares
//
// 							board[i][j] = byte('0' + theGuesses[0])
// 							if try := doSolveSudoku(board); try != nil {
// 								return try
// 							}
// 							board[i][j] = byte('0' + theGuesses[1])
// 							if try := doSolveSudoku(board); try != nil {
// 								return try
// 							}
// 							board[i][j] = byte('0' + theGuesses[2])
// 							infoBoard.eliminateRow(i, theGuesses[2])
// 							infoBoard.eliminateCol(j, theGuesses[2])
// 							infoBoard.eliminateBlock(i, j, theGuesses[2])
// 							breakTime = true
// 							break
// 							// close if 2 possible
// 						}
// 						// close inner loop
// 					}
// 					if breakTime == true {
// 						break
// 					}
// 					// close outer loop
// 				}
// 				if !breakTime {
// 					fmt.Println("infinite loop")
// 				}
// 			}
// 			// close else
// 		}
// 		done = isSolved(&board)
// 	}
// 	return board
// }
//
// func foundContradiction(infoBoard *truthBoard) bool {
// 	for _, v := range *infoBoard {
// 		for _, u := range v {
// 			if nonePossible(u) {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
//
// func isSolved(board *[][]byte) bool {
// 	for _, v := range *board {
// 		for _, u := range v {
// 			if u == '.' {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }
//
// // FIND AND SOLVE:
// // Solves one square.
// // ASSUMPTION: At any point, there exists at least one SOLVABLE SQUARE
// // ie., we NEVER HAVE TO GUESS.
// //
// //
//
// func findAndSolve(board *[][]byte, infoBoard *truthBoard) []int {
// 	check := 0
// 	for i, v := range *board {
// 		for j, u := range v {
// 			if u != '.' {
// 				continue
// 			}
// 			for k := 1; k < 10; k++ {
// 				if infoBoard[i][j][k] {
// 					if isUniqueInBlock(i, j, k, infoBoard) || isUniqueInRow(i, j, k, infoBoard) || isUniqueInCol(i, j, k, infoBoard) {
// 						(*board)[i][j] = byte('0' + k)
// 						return []int{i, j, k}
// 					}
// 				}
// 			}
// 			check = onlyOnePossible(infoBoard[i][j])
// 			if check != -1 {
// 				(*board)[i][j] = byte('0' + check)
// 				return []int{i, j, check}
// 			}
// 		}
// 	}
// 	return []int{-1, 0, 0}
// }
//
// func onlyOnePossible(square truthMap) int {
// 	firstFound := false
// 	ret := -1
// 	for k, v := range square {
// 		if v && !firstFound {
// 			firstFound = true
// 			ret = k
// 			continue
// 		}
// 		if v {
// 			return -1
// 		}
// 	}
// 	return ret
// }
//
// func nonePossible(square truthMap) bool {
// 	if square[0] {
// 		return false
// 	}
//
// 	for _, v := range square {
// 		if v {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func twoPossible(square truthMap) []int {
// 	if square[0] {
// 		return nil
// 	}
// 	firstFound := false
// 	secondFound := false
// 	ret := []int{}
// 	for k, v := range square {
// 		if !v {
// 			continue
// 		}
// 		if firstFound && secondFound {
// 			return nil
// 		}
// 		if !firstFound {
// 			firstFound = true
// 			ret = append(ret, k)
// 			continue
// 		}
// 		if !secondFound {
// 			secondFound = true
// 			ret = append(ret, k)
// 			continue
// 		}
// 	}
// 	if !secondFound {
// 		return nil
// 	}
// 	return ret
// }
//
// func threePossible(square truthMap) []int {
// 	if square[0] {
// 		return nil
// 	}
// 	firstFound := false
// 	secondFound := false
// 	thirdFound := false
// 	ret := []int{}
// 	for k, v := range square {
// 		if !v {
// 			continue
// 		}
// 		if firstFound && secondFound && thirdFound {
// 			return nil
// 		}
// 		if !firstFound {
// 			firstFound = true
// 			ret = append(ret, k)
// 			continue
// 		}
// 		if !secondFound {
// 			secondFound = true
// 			ret = append(ret, k)
// 			continue
// 		}
// 		if !thirdFound {
// 			thirdFound = true
// 			ret = append(ret, k)
// 			continue
// 		}
// 	}
// 	if !thirdFound {
// 		return nil
// 	}
// 	return ret
// }
//

//isUniqueInBlock returns true if there is no possible other place in the block for the number n to be
//besides, importantly, board[row][col]
// func isUniqueInBlock(row, col, n int, infoBoard *truthBoard) bool {
// 	// first things first, which block are we in?
// 	var theRow int
// 	var theCol int
// 	if row < 3 {
// 		theRow = 0
// 	} else if row > 5 {
// 		theRow = 2
// 	} else {
// 		theRow = 1
// 	}
//
// 	if col < 3 {
// 		theCol = 0
// 	} else if col > 5 {
// 		theCol = 2
// 	} else {
// 		theCol = 1
// 	}
//
// 	// Now, build the loop bounds
// 	// NOTE: High bounds are exclusive!
// 	var rLow int
// 	var rHigh int
// 	var cLow int
// 	var cHigh int
// 	switch theRow {
// 	case 0:
// 		rLow = 0
// 		rHigh = 3
// 	case 1:
// 		rLow = 3
// 		rHigh = 6
// 	case 2:
// 		rLow = 6
// 		rHigh = 9
// 	}
// 	switch theCol {
// 	case 0:
// 		cLow = 0
// 		cHigh = 3
// 	case 1:
// 		cLow = 3
// 		cHigh = 6
// 	case 2:
// 		cLow = 6
// 		cHigh = 9
// 	}

// finally, the loop
// we MUST avoid the box in question,
// so we'll need a new bool every j loop
// 	notOurBox := false
// 	for i := rLow; i < rHigh; i++ {
// 		for j := cLow; j < cHigh; j++ {
// 			notOurBox = i != row || j != col
// 			if notOurBox && (*infoBoard)[i][j][n] == true {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

//
// // isUniqueInRow returns true if there is no possible other place in the row for the number n to be
// //besides, importantly, board[row][col]
// func isUniqueInRow(row, col, n int, infoBoard *truthBoard) bool {
// 	for i, v := range (*infoBoard)[row] {
// 		if i != col && v[n] {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// // isUniqueInRow returns true if there is no possible other place in the column for the number n to be
// //besides, importantly, board[row][col]
// func isUniqueInCol(row, col, n int, infoBoard *truthBoard) bool {
// 	for i, v := range *infoBoard {
// 		if i == row {
// 			continue
// 		}
// 		for j, u := range v {
// 			if j != col {
// 				continue
// 			}
// 			if u[n] {
// 				return false
// 			}
// 		}
//
// 	}
// 	return true
// }
//
// // UPDATE ALL
// //
// //
// //
// //
//
// // EliminateRow eliminates the row on the truthmap
// func (infoBoard *truthBoard) eliminateRow(row, n int) {
// 	for _, v := range (*infoBoard)[row] {
// 		v[n] = false
// 	}
// }
//
// // EliminateCol eliminates the col on the truthmap
// func (infoBoard *truthBoard) eliminateCol(col, n int) {
// 	for _, v := range *infoBoard {
// 		v[col][n] = false
// 	}
// }
//
// // EliminateBlock eliminates the block on the truthmap
// func (infoBoard *truthBoard) eliminateBlock(row, col, n int) {
//
// 	// first things first, which block are we in?
//
// 	var theRow int
// 	var theCol int
//
// 	if row < 3 {
// 		theRow = 0
// 	} else if row > 5 {
// 		theRow = 2
// 	} else {
// 		theRow = 1
// 	}
//
// 	if col < 3 {
// 		theCol = 0
// 	} else if col > 5 {
// 		theCol = 2
// 	} else {
// 		theCol = 1
// 	}
//
// 	// Now, build the loop bounds
// 	// NOTE: High bounds are exclusive!
// 	var rLow int
// 	var rHigh int
// 	var cLow int
// 	var cHigh int
// 	switch theRow {
// 	case 0:
// 		rLow = 0
// 		rHigh = 3
// 	case 1:
// 		rLow = 3
// 		rHigh = 6
// 	case 2:
// 		rLow = 6
// 		rHigh = 9
// 	}
// 	switch theCol {
// 	case 0:
// 		cLow = 0
// 		cHigh = 3
// 	case 1:
// 		cLow = 3
// 		cHigh = 6
// 	case 2:
// 		cLow = 6
// 		cHigh = 9
// 	}
//
// 	// finally, the loop
// 	for i := rLow; i < rHigh; i++ {
// 		for j := cLow; j < cHigh; j++ {
// 			(*infoBoard)[i][j][n] = false
// 		}
// 	}
// }
//
// // // updateAll sweeps the whole board for information and concats it all into the infoboard
// // // update will:
// // // range through the board looking for a filled square
// // // once it finds one it will:
// // // update that square's truthboard "0" keyval to true
// // // update the row, col, and block's possibilities
// // // NOTE: We're being very inefficient, time and memory-wise.
// // // If nXn sudoku ever catches on, that may be a problem
// func updateAll(board *[][]byte, infoBoard *truthBoard) {
// 	for i, v := range *board {
// 		for j, u := range v {
// 			if u != '.' {
// 				//first, update infoboard
// 				(*infoBoard)[i][j][0] = true
// 				//now, eliminate n from row on infoboard
// 				infoBoard.eliminateRow(i, int(u-48))
// 				infoBoard.eliminateCol(j, int(u-48))
// 				infoBoard.eliminateBlock(i, j, int(u-48))
// 				//place "true" at valid solution
// 				(*infoBoard)[i][j][int(u)] = true
// 			}
// 		}
// 	}
// }
//
// // ['.', '.', '9', '7', '4', '8', '.', '.', '.']
// // ['7', '.', '.', '.', '.', '.', '.', '.', '.']
// // ['.', '2', '.', '1', '.', '9', '.', '.', '.']
// // ['.', '.', '7', '.', '.', '.', '2', '4', '.']
// // ['.', '6', '4', '.', '1', '.', '5', '9', '.']
// // ['.', '9', '8', '.', '.', '.', '3', '.', '.']
// // ['.', '.', '.', '8', '.', '3', '.', '2', '.']
// // ['.', '.', '.', '.', '.', '.', '.', '.', '6']
// // ['.', '.', '.', '2', '7', '5', '9', '.', '.']
// //
// // ['.', '.', '9', '7', '4', '8', '1', '.', '2']
// // ['7', '.', '.', '6', '.', '2', '.', '.', '9']
// // ['.', '2', '.', '1', '.', '9', '.', '.', '.']
// // ['.', '.', '7', '9', '8', '6', '2', '4', '1']
// // ['2', '6', '4', '3', '1', '7', '5', '9', '8']
// // ['1', '9', '8', '5', '2', '4', '3', '6', '7']
// // ['9', '.', '.', '8', '6', '3', '.', '2', '.']
// // ['.', '.', '2', '4', '9', '1', '.', '.', '6']
// // ['.', '.', '.', '2', '7', '5', '9', '.', '.']
// //
// // ['.', '3', '9', '7', '4', '8', '1', '.', '2']
// // ['7', '.', '.', '6', '.', '2', '.', '.', '9']
// // ['.', '2', '.', '1', '.', '9', '.', '.', '.']
// // ['.', '.', '7', '9', '8', '6', '2', '4', '1']
// // ['2', '6', '4', '3', '1', '7', '5', '9', '8']
// // ['1', '9', '8', '5', '2', '4', '3', '6', '7']
// // ['9', '.', '.', '8', '6', '3', '.', '2', '.']
// // ['.', '.', '2', '4', '9', '1', '.', '.', '6']
// // ['.', '.', '.', '2', '7', '5', '9', '.', '.']
//
// // ['.', '.', '9', '7', '4', '8', '.', '.', '.']
// // ['7', '.', '.', '.', '.', '.', '.', '.', '.']
// // ['.', '2', '.', '1', '.', '9', '.', '.', '.']
// // ['.', '.', '7', '.', '.', '.', '2', '4', '.']
// // ['.', '6', '4', '.', '1', '.', '5', '9', '.']
// // ['.', '9', '8', '.', '.', '.', '3', '.', '.']
// // ['.', '.', '.', '8', '.', '3', '.', '2', '.']
// // ['.', '.', '.', '.', '.', '.', '.', '.', '6']
// // ['.', '.', '.', '2', '7', '5', '9', '.', '.']
// //
// // ['.', '.', '9', '7', '4', '8', '6', '.', '2']
// // ['7', '.', '.', '6', '.', '2', '.', '.', '9']
// // ['.', '2', '.', '1', '.', '9', '.', '.', '.']
// // ['.', '.', '7', '9', '8', '6', '2', '4', '1']
// // ['2', '6', '4', '3', '1', '7', '5', '9', '8']
// // ['1', '9', '8', '5', '2', '4', '3', '6', '7']
// // ['9', '.', '.', '8', '6', '3', '.', '2', '.']
// // ['.', '.', '2', '4', '9', '1', '.', '.', '6']
// // ['.', '.', '.', '2', '7', '5', '9', '.', '.']
// //
// // ['3', '.', '9', '7', '4', '8', '6', '.', '2']
// // ['7', '.', '.', '6', '.', '2', '.', '.', '9']
// // ['.', '2', '.', '1', '.', '9', '.', '.', '.']
// // ['.', '.', '7', '9', '8', '6', '2', '4', '1']
// // ['2', '6', '4', '3', '1', '7', '5', '9', '8']
// // ['1', '9', '8', '5', '2', '4', '3', '6', '7']
// // ['9', '.', '.', '8', '6', '3', '.', '2', '.']
// // ['.', '.', '2', '4', '9', '1', '.', '.', '6']
// // ['.', '.', '.', '2', '7', '5', '9', '.', '.']
