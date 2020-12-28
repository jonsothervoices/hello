package bits

import (
	"strconv"
)

func getBit(n, b int) bool { return (1<<b)&n != 0 }

func setBit(n, b int) int { return (1 << b) | n }

func clearBit(n, b int) int { return n & (^(1 << b)) }

func updateBit(n, b int, v bool) int {
	if v {
		return setBit(n, b)
	}
	return clearBit(n, b)
}

//5.1 Insertion: Given two 32-bit numbers M and N, and bit position i, write a function to insert M into N such that M starts at bit i.
//4:17-5:07~~50 minutes
func insertBits(n, m, i int) int {
	//create bool slice
	mBin := strconv.FormatInt(int64(m), 2)
	for j, v := range mBin {
		if v == '1' {
			n = setBit(n, i-j)
		} else {
			n = clearBit(n, i-j)
		}
	}
	return n
}

//5.2 Binary to String: Given a real number from 0 to 1, print the binary representation. If the number cannot be represented accurately in binary with at most 32 characters, print ERROR.

//5.3 Flip bit to Win: Given an integer, find the length of the longest sequece of 1s you could create by flipping 1 bit.

// func flipToWin(n int) int{
// 	//create slice from n
// 	nBin := strconv.FormatInt(int64(n), 2)
// 	l:= len(nBin)-1
// 	counter:=0
// 	seq1:=0
// 	seq2:=2
// 	//find the 2 longest sequences of 1 that are seperated by 1 zero
// 	for i:=l;i>=0; i--{
// 		if nBin[l]==1{
// 			counter++
// 		}
// 		else {
// 			seq1=counter
// 		}
// 	}
// 	//return the length of those 2 sequences plus 1
// }
