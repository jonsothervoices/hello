package bits

import (
	"fmt"
	"strconv"
	"strings"
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

func flipToWin(n int) int {
	ans := 0
	nInt := int64(n)
	nBin := strconv.FormatInt(nInt, 2)
	if !strings.Contains(nBin, "0") {
		return len(nBin)
	}
	for i, v := range nBin {
		if v == '0' {
			if zeroNextToOne(nBin, i) {
				sum := lSum(nBin, i) + rSum(nBin, i)
				fmt.Printf("lsum %v, rsum %v\n", lSum(nBin, i), rSum(nBin, i))
				if sum > ans {
					ans = sum
				}
			}
		}
	}
	return ans + 1
}

func zeroNextToOne(s string, i int) bool {
	if i == 0 {
		if len(s) > 1 {
			return s[i+1] == '1'
		}
		return true
	}
	if i == len(s)-1 {
		if len(s) > 1 {
			return s[i-1] == '1'
		}
		return true
	}
	return s[i+1] == '1' || s[i-1] == '1'
}

func lSum(s string, n int) int {
	ret := 0
	for i := n - 1; i >= 0 && s[i] != '0'; i-- {
		ret++
	}
	return ret
}

func rSum(s string, n int) int {
	ret := 0
	for i := n + 1; i < len(s) && s[i] != '0'; i++ {
		ret++
	}
	return ret
}
