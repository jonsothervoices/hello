package leetcode

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		theBoard sBoard
		expected bool
	}{
		{newSBoard(invBoard1), false},
		{newSBoard(valBoard1), true},
	}
	for i, datest := range tests {
		actual := datest.theBoard.isValid()
		if actual != datest.expected {
			t.Errorf("Case %v: Expected %v, Actual %v", i+1, datest.expected, actual)
		}
	}
}

func TestIsSolved(t *testing.T) {
	var tests = []struct {
		theBoard sBoard
		expected bool
	}{
		{newSBoard(solvedBoard1), true},
		{newSBoard(almostBoard1), false},
	}
	for i, datest := range tests {
		actual := datest.theBoard.isSolved()
		if actual != datest.expected {
			t.Errorf("Case %v: Expected %v, Actual %v", i+1, datest.expected, actual)
		}
	}
}

func TestSolveEasySquares(t *testing.T) {
	var tests = []struct {
		theBoard sBoard
		expected sBoard
	}{
		{newSBoard(almostBoard1), newSBoard(solvedBoard1)},
	}
	for i, datest := range tests {
		datest.theBoard.solveEasySquares()
		for rInd, row := range datest.theBoard.Data {
			for jInd := range row {
				if datest.expected.Data[rInd][jInd] != datest.theBoard.Data[rInd][jInd] {
					t.Errorf("Case %v: at coordinates %v, %v: Expected %v, Actual %v", i+1, rInd, jInd, datest.expected.Data[rInd][jInd], datest.theBoard.Data[rInd][jInd])
				}
			}
		}

	}
}

// func TestFindAndSolve(t *testing.T) {
// 	var tests = []struct {
// 		theBoard   [][]byte
// 		infoBoard truthBoard
// 		expected [][]byte
// 	}{
// 		{[][]byte{{'0','3','4','6','7','8','9','1','2'},{'6','7','2','1','9','5','3','4','8'},{'1','9','8','3','4','2','5','6','7'},{'8','5','9','7','6','1','4','2','3'},{'4','2','6','8','5','3','7','9','1'},{'7','1','3','9','2','4','8','5','6'},{'9','6','1','5','3','7','2','8','4'},{'2','8','7','4','1','9','6','3','5'},{'3','4','5','2','8','6','1','7','9'}},[][]byte{{'5','3','4','6','7','8','9','1','2'},{'6','7','2','1','9','5','3','4','8'},{'1','9','8','3','4','2','5','6','7'},{'8','5','9','7','6','1','4','2','3'},{'4','2','6','8','5','3','7','9','1'},{'7','1','3','9','2','4','8','5','6'},{'9','6','1','5','3','7','2','8','4'},{'2','8','7','4','1','9','6','3','5'},{'3','4','5','2','8','6','1','7','9'}}},
//
// 	}
// 	for i, datest := range tests {
// 		actual := onlyOnePossible(datest.theMap)
// 		if actual != datest.expected {
// 			t.Errorf("Case %v: expected %v, actual %v", i+1, datest.expected, actual)
// 		}
// 	}
// }

// func TestOnlyOnePossible(t *testing.T) {
// 	var tests = []struct {
// 		theMap   truthMap
// 		expected int
// 	}{
// 		{truthMap{3: true}, 3},
// 		{truthMap{}, -1},
// 		{truthMap{3: true, 5: true}, -1},
// 		{truthMap{1: true, 2: true, 3: false}, -1},
// 		{truthMap{7: true}, 7},
// 		{truthMap{1: true, 9: true}, -1},
// 	}
// 	for i, datest := range tests {
// 		actual := onlyOnePossible(datest.theMap)
// 		if actual != datest.expected {
// 			t.Errorf("Case %v: expected %v, actual %v", i+1, datest.expected, actual)
// 		}
// 	}
// }
//
// func TestEliminateRow(t *testing.T) {
// 	var tests = []struct {
// 		row      int
// 		n        int
// 		theBoard truthBoard
// 		expected truthBoard
// 	}{
// 		{0, 5, truthBoard{{{5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}}}, truthBoard{{{5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}}}},
// 		//{1, 2, truthBoard{}, truthBoard{}},
// 		{0, 5, truthBoard{{{5: true, 6: true, 7: false}, {3: true, 5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}}, {{5: true, 6: true, 7: false}, {5: true}}}, truthBoard{{{5: false, 6: true, 7: false}, {3: true, 5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}, {5: false}}, {{5: true, 6: true, 7: false}, {5: true}}}},
// 	}
//
// 	for i, datest := range tests {
// 		datest.theBoard.eliminateRow(datest.row, datest.n)
// 		for ii, v := range datest.theBoard {
// 			for j, u := range v {
// 				for k := 1; k < 10; k++ {
// 					if u[k] != datest.expected[ii][j][k] {
// 						t.Errorf("Case %v: location %v, %v, key %v: actual %v, expected %v", i+1, ii, j, k, u, datest.expected[ii][j][k])
// 					}
// 				}
// 			}
// 		}
// 	}
// }
//
// func TestEliminateCol(t *testing.T) {
// 	var tests = []struct {
// 		col      int
// 		n        int
// 		theBoard truthBoard
// 		expected truthBoard
// 	}{
// 		{0, 5, truthBoard{{{5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}}, {{5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}}, {{5: true}}, {{5: true}}, {{5: true}}, {{5: true}}, {{5: true}}, {{5: true}}, {{5: true}}}, truthBoard{{{5: false}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}}, {{5: false}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}, {5: true}}, {{5: false}}, {{5: false}}, {{5: false}}, {{5: false}}, {{5: false}}, {{5: false}}, {{5: false}}}},
// 		{0, 5, truthBoard{{{5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}}, {{5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}}, {{5: true, 6: true}}, {{5: true, 6: true}}, {{5: true, 6: true}}, {{5: true, 6: true}}, {{5: true, 6: true}}, {{5: true, 6: true}}, {{5: true, 6: true}}}, truthBoard{{{5: false, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}}, {{5: false, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}, {5: true, 6: true}}, {{5: false, 6: true}}, {{5: false, 6: true}}, {{5: false, 6: true}}, {{5: false, 6: true}}, {{5: false, 6: true}}, {{5: false, 6: true}}, {{5: false, 6: true}}}},
// 	}
//
// 	for i, datest := range tests {
// 		datest.theBoard.eliminateCol(datest.col, datest.n)
// 		for ii, v := range datest.theBoard {
// 			for j, u := range v {
// 				for k := 1; k < 10; k++ {
// 					if u[k] != datest.expected[ii][j][k] {
// 						t.Errorf("Case %v: location %v, %v, key %v: actual %v, expected %v", i+1, ii, j, k, u, datest.expected[ii][j][k])
// 					}
// 				}
// 			}
// 		}
// 	}
// }

var blankBoard = [][]byte{
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
}

// invBoard1 has invalid first square
var invBoard1 = [][]byte{
	{'.', '4', '5', '6', '.', '.', '.', '.', '.'},
	{'1', '.', '9', '.', '.', '.', '.', '.', '.'},
	{'2', '8', '7', '.', '.', '.', '.', '.', '.'},
	{'3', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
}

// valBoard1 is valid, ie has no contradictions
var valBoard1 = [][]byte{
	{'3', '.', '9', '7', '4', '8', '6', '.', '2'},
	{'7', '.', '.', '6', '.', '2', '.', '.', '9'},
	{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
	{'.', '.', '7', '9', '8', '6', '2', '4', '1'},
	{'2', '6', '4', '3', '1', '7', '5', '9', '8'},
	{'1', '9', '8', '5', '2', '4', '3', '6', '7'},
	{'9', '.', '.', '8', '6', '3', '.', '2', '.'},
	{'.', '.', '2', '4', '9', '1', '.', '.', '6'},
	{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
}

// solvedBoard1 is fully solved
var solvedBoard1 = [][]byte{
	{'5', '3', '4', '6', '7', '8', '9', '1', '2'},
	{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
	{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
	{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
	{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
	{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
	{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
	{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
	{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
}

// almostBoard1 is almost solved
var almostBoard1 = [][]byte{
	{'5', '.', '4', '6', '7', '8', '9', '1', '2'},
	{'6', '7', '2', '1', '9', '5', '3', '4', '8'},
	{'1', '9', '8', '3', '4', '2', '5', '6', '7'},
	{'8', '5', '9', '7', '6', '1', '4', '2', '3'},
	{'4', '2', '6', '8', '5', '3', '7', '9', '1'},
	{'7', '1', '3', '9', '2', '4', '8', '5', '6'},
	{'9', '6', '1', '5', '3', '7', '2', '8', '4'},
	{'2', '8', '7', '4', '1', '9', '6', '3', '5'},
	{'3', '4', '5', '2', '8', '6', '1', '7', '9'},
}
